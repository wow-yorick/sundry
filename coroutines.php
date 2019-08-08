<?php

function xrange($start, $end, $step = 1) {
    for($i = $start;$i <= $end; $i+= $step) {
        yield $i;
    }
}

// foreach(xrange(1,10000) as $num ) {
//     echo $num, "\n";
// }
// $range = xrange(1, 10000);
// var_dump($range);
// var_dump($range instanceof Iterator);

// function logger($fileName) {
//     $fileHandle = fopen($fileName, 'a');
//     while(true) {
//         fwrite($fileHandle, yield . "\n");
//     }
// }
// echo __DIR__;
// $logger = logger(__DIR__.'/log');
// $logger->send('Foo');
// $logger->send('Bar');
// echo 'ok';
// function gen() {
//     $ret = (yield 'yield1');
//     var_dump($ret);
//     $ret = (yield 'yield2');
//     var_dump($ret);
// }

// $gen = gen();
// var_dump($gen->current());
// var_dump($gen->send('ret1'));
// var_dump($gen->send('ret2'));

class Task {
    protected $taskId;
    protected $coroutine;
    protected $sendValue = null;
    protected $beforeFirstYield = true;
    protected $exception = null;

    public function __construct($taskId, Generator $coroutine) {
        $this->taskId = $taskId;
        //$this->coroutine = $coroutine;
        $this->coroutine = stackedCoroutine($coroutine);
    }

    public function setException($exception) {
        $this->exception = $exception;
    }

    public function getTaskId() {
        return $this->taskId;
    }

    public function setSendValue($sendValue) {
        $this->sendValue = $sendValue;
    }

    public function run() {
        if($this->beforeFirstYield) {
            $this->beforeFirstYield = false;
            return $this->coroutine->current();
        } elseif ($this->Exception) {
            $retval = $this->coroutine->throw($this->exception);
            $this->exception = null;
            return $retval;
        } else {
            $retval = $this->coroutine->send($this->sendValue);
            $this->sendValue = null;
            return $retval;
        }
    }

    public function isFinished() {
        return !$this->coroutine->valid();
    }
}

// class SqlQueue {
//     public $queueMap = array();
//     public function __construct() {
        
//     }
//     public function enqueue($data, $key = -1) {
//         array_push($this->queueMap, $data);
//     }

//     public function dequeue() {
//         $queueKeys = array_keys($this->queueMap);
//         $retval = $this->queueMap[$queueKeys[0]];
//         unset($this->queueMap[$queueKeys[0]]);
//         return $retval;
//         //return array_pop($this->queueMap);
//     }

//     public function isEmpty() {
//         return empty($this->queueMap);
//     }
// }

class Scheduler {
    protected $maxTaskId = 0;
    protected $taskMap = []; //taskId => task
    protected $taskQueue;
    protected $waitingForRead = [];
    protected $waitingForWrite = [];

    public function __construct() {
        $this->taskQueue = new SplQueue();
    }

    public function newTask(Generator $coroutine) {
        $tid = ++$this->maxTaskId;
        $task = new Task($tid, $coroutine);
        $this->taskMap[$tid] = $task;
        $this->schedule($task);
        return $tid;
    }

    public function killTask($tid) {
        if (!isset($this->taskMap[$tid])) {
            return false;
        }
        unset($this->taskMap[$tid]);

        foreach($this->taskQueue as $i => $task) {
            if($task->getTaskId() == $tid) {
                unset($this->taskQueue[$i]);
                break;
            }
        }

        return true;
    }

    public function waitForRead($socket, Task $task) {
        if(isset($this->waitingForRead[(int) $socket])) {
            $this->waitingForRead[(int) $socket][1][] = $task;
        } else {
            $this->waitingForRead[(int) $socket] = [$socket,[$task]];
        }
    }

    public function waitForWrite($socket, Task $task) {
        if(isset($this->waitingForWrite[(int) $socket])) {
            $this->waitingForWrite[(int) $socket][1][] = $task;
        } else {
            $this->waitingForWrite[(int) $socket] = [$socket,[$task]];
        }
    }

    protected function ioPoll($timeout) {
        $rSocks = [];
        foreach($this->waitingForRead as list($socket)) {
            $rSocks[] = $socket;
        }

        $wSocks = [];
        foreach($this->waitingForWrite as list($socket)) {
            $wSocks[] = $socket;
        }

        $eSocks = [];

        if(!stream_select($rSocks, $wSocks, $eSocks, $timeout)) {
            return;
        }

        foreach($rSocks as $socket) {
            list(, $tasks) = $this->waitingForRead[(int) $socket];
            unset($this->waitingForRead[(int) $socket]);

            foreach($tasks as $task) {
                $this->schedule($task);
            }
        }

        foreach($wSocks as $socket) {
            list(, $tasks) = $this->waitingForWrite[(int) $socket];
            unset($this->waitingForWrite[(int) $socket]);

            foreach($tasks as $task) {
                $this->schedule($task);
            }
        }
    }

    protected function ioPollTask() {
        while(true) {
            if($this->taskQueue->isEmpty()) {
                $this->ioPoll(null);
            } else {
                $this->ioPoll(0);
            }
            yield;
        }
    }

    public function schedule(Task $task) {
        $this->taskQueue->enqueue($task);
    }

    public function run() {
        while(!$this->taskQueue->isEmpty()) {
            $task = $this->taskQueue->dequeue();
            $retval = $task->run();

            if($retval instanceof SystemCall) {
                try {
                    $retval($task, $this);
                } catch(Exception $e) {
                    $task->setException($e);
                    $this->schedule($task);
                }
                continue;
            }

            if($task->isFinished()) {
                unset($this->taskMap[$task->getTaskId()]);
            } else {
                $this->schedule($task);
                //print_r($this->taskQueue->queueMap);exit;
            }
        }

    }
}

class SystemCall {
    protected $callback;

    public function __construct(callable $callback) {
        $this->callback = $callback;
    }

    public function __invoke(Task $task, Scheduler $scheduler) {
        $callback = $this->callback;
        return $callback($task, $scheduler);
    }
}

function getTaskId() {
    return new SystemCall(function(Task $task, Scheduler $scheduler) {
        $task->setSendValue($task->getTaskId());
        $scheduler->schedule($task);
    });
}

function newTask(Generator $coroutine) {
    return new SystemCall(function(Task $task, Scheduler $scheduler) use ($coroutine) {
        $task->setSendValue($scheduler->newTask($coroutine));
        $scheduler->schedule($task);
    });
}

function killTask($tid) {
    return new SystemCall(
        function(Task $task, Scheduler $scheduler) use ($tid) {
            if($retval = $scheduler->killTask($tid)) {
                //$task->setSendValue($scheduler->killTask($tid));
                //$task->setSendValue($retval);
                $scheduler->schedule($task);
            } else {
                throw new InvalidArgumentException('invalid task ID!');
            }
        }
    );
}

function waitForRead($socket) {
    return new SystemCall(
        function(Task $task, Scheduler $scheduler) use ($socket) {
            $scheduler->waitForRead($socket, $task);
        }
    );
}

function waitForWrite($socket) {
    return new SystemCall(
        function(Task $task, Scheduler $scheduler) use ($socket) {
            $scheduler->waitForWrite($socket, $task);
        }
    );
}

function server($port) {
    echo "Starting server at port $port... \n";

    $socket = @stream_socket_server("tcp://localhost:$port", $errNo, $errStr);
    if(!$socket) throw new Exception($errStr, $errNo);

    stream_set_blocking($socket, 0);

    $socket = new CoSocket($socket);
    while(true) {
        // yield waitForRead($socket);
        // $clientSocket = stream_socket_accept($socket, 0);
        // yield newTask(handleClient($clientSocket));
        yield newTask(
            handleClient(yield $socket->accept())
        );

    }
}

function handleClient($socket) {
    // yield waitForRead($socket);
    // $data = fread($socket, 8192);
    $data = (yield $socket->read(8192));

    $msg = "Received following request:\n\n$data";
    $msgLength = strlen($msg);
        $response = <<<RES
HTTP/1.1 200 OK\r
Content-Type: text/plain\r
Content-Length: $msgLength\r
Connection: close\r
\r
$msg
RES;
        // yield waitForWrite($socket);
        // fwrite($socket, $response);
        // fclose($socket);
        yield $socket->write($response);
        yield $socket->close();
}

// $scheduler = new Scheduler;
// $scheduler->newTask(server(8000));
// $scheduler->run();

// function task($max) {
//     $tid = (yield getTaskId());
//     for($i = 1;$i <=$max; ++$i) {
//         echo "this task $tid iteration $i.\n";
//         yield;
//     }
// }

// function childTask() {
//     $tid = (yield getTaskId());
//     while(true) {
//         echo "Child task $tid still alive!\n";
//         yield;
//     }
// }

// function task() {
//     $tid = (yield getTaskId());
//     $childTid = (yield newTask(childTask()));

//     for($i = 1; $i <= 6; ++$i) {
//         echo "Parent task $tid iteration $i.\n";
//         yield;

//         if($i == 3) yield killTask($childTid);
//     }
// }

function echoTimes($msg, $max) {
    for($i; $i <= $max; ++$i) {
        echo "$msg iteration $i \n";
        yield;
    }
}

function task() {
    echoTimes('foo', 10);
    echo "---\n";
    echoTimes('bar', 5);
    yield;
}

class CoroutineReturnValue {
    protected $value;

    public function __construct($value) {
        $this->value = $value;
    }

    public function getValue() {
        return $this->value;
    }
}

function retval($value) {
    return new CoroutineReturnValue($value);
}

function stackedCoroutine(Generator $gen) {
    $stack = new SplStack;
    $exception = null;
    for(;;) {
        try {
            if($exception) {
                $gen->throw($exception);
                $exception = null;
                continue;
            }
        } catch(Exception $e) {
            if($stack->isEmpty()) {
                throw $e;
            }
            $gen = $stack->pop();
            $exception = $e;
        }
        $value = $gen->current();

        if($value instanceof Generator) {
            $stack->push($gen);
            $gen = $value;
            continue;
        }

        $isReturnValue = $value instanceof CoroutineReturnValue;
        if(!$gen->valid() || $isReturnValue) {
            if($stack->isEmpty()) {
                return;
            }

            $gen = $stack->pop();
            $gen->send($isReturnValue ? $value->getValue() : NULL);
            continue;
        }
        try {
            $sendValue = yield $gen->key() => $value;
        } catch(Exception $e) {
            $gen->throw($e);
            continue;
        } 
        $gen->send($sendValue);
    }
}

class CoSocket {
    protected $socket;

    public function __construct($socket) {
        $this->socket = $socket;
    }

    public function accept() {
        yield waitForRead($this->socket);
        yield retval(new CoSocket(stream_socket_accept($this->socket, 0)));
    }

    public function read($size) {
        yield waitForRead($this->socket);
        yield retval(fread($this->socket), $size);
    }

    public function write($string) {
        yield waitForWrite($this->socket);
        fwrite($this->socket, $string);
    }

    public function close() {
        @fclose($this->socket);
    }
}

function task() {
    try{
        yield killTask(500);
    } catch(Exception $e) {
        echo "Tried to kill task 500 but failed:".$e->getMessage();
    }
}

//(yield retval("I'm a return value!"));
//$scheduler = new Scheduler();
//$scheduler->newTask(task());
// $scheduler->newTask(task(10));
// $scheduler->newTask(task(5));
//$scheduler->run();

// function gen() {
//     echo "Foo\n";
//     try {
//         yield 'ab';
//     } catch(Exception $e) {
//         echo "Exception: {$e->getMessage()}\n";
//     }
//     echo "Bar\n";
// }

// $gen = gen();
// var_dump($gen->rewind());
// $gen->throw(new Exception('Test'));
// echo 'ok';