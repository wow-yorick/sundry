package main

//What is the overhead of defer?
//runtime.deferproc and runtime.deferreturn cause context copy and retrieval on stack memory

//What is the malloc threshold of Map object? How to modify it?
//The default limit is 128.
//It can be modified by changing the value of maxKeySize and maxValueSize in runtime.hashmap

// What does runtime.newobject() do?
// Does make() and `new` operation always invoke runtime.newobject()?

//runtime.newobject() is to allocate heap memory.
//make() and `new` operation will not invoke runtime.newobject() when it is inlined or optimized.

//Briefly introduce the benefits of SSA(Static Single Assignment) form.
// Remove unnecessary code
// Remove redundant varibles and logics
// Optimize register or memory allocation
// Transform varibles to constants when possible
// Transform "high cost" instructions to "low cost"

//Briefly describe the bootstrapping process of a go executable.

// Run the platform-specific assembly that is located under $GOROOT/src/runtime/
// runtime·args(): Parse terminal arguments
// runtime·osinit(): Initialize CPU cores
// runtime·schedinit(): Initialize goroutine scheduler, stack memory, terminal arguments, environment variables, debug parameter, gc, GOMAXPROCS, ...
// runtime·mstart(): Start gc monitor, enable gc, import dependencies and run init() functions, finally run main.main()

//What are the differences between unbuffered and buffered channels?

// For unbuffered channel, the sender will block on the channel until the receiver receives the data from the channel, whilst the receiver will also block on the channel until sender sends data into the channel.
// Compared with unbuffered counterpart, the sender of buffered channel will block when there is no empty slot of the channel, while the receiver will block on the channel when it is empty.

//What is the destructor in go?

//There is no destructor in go. But runtime.SetFinalizer() can set a callback function for a pointer.

//Briefly describe how gc works in go?
// MarkWorker goroutine recursively scan all the objects and colors them into white(inaccessible), gray(pending), black(accessible). But finally they will only be black and white objcts.
// In compile time, the compiler has already injected a snippet called write barrier to monitor all the modifications from any goroutines to heap memory.
// When "Stop the world" is performed, scheduler hibernates all threads and preempt all goroutines.
// Garbage collector will recycle all the inaccessible objects so heap or central can reuse.
// If the whole span of memory are unused, it can be freed to OS.
// Perform "Start the world" to wake cpu cores and threads, and resume the execution of goroutines.

//What is the difference between C.sleep() and time.Sleep()?

// C.sleep() invokes syscall sleep, which causes idle threads
// time.Sleep() is optimized for goroutine so syscall sleep is not involved.

//Briefly introduce the strategies that how go runtime allocates memory.

// For small objects(<=32KB), go runtime starts with cache firstly, then central, and finally heap.
// For big objects(>32KB), directly from heap.

//When go runtime allocates memory from heap, and when from stack?
// For small objects whose life cycle is only within the stack frame, stack memory is allocated.
// For small objects that will be passed across stack frames, heap memory.
// For big objects(>32KB), heap memory.
// For small objects that could escape to heap but actually inlined, stack memory.

//List the functions can stop or suspend the execution of current goroutine, and explain their differences.

// runtime.Gosched: give up CPU core and join the queue, thus will be executed automatically.
// runtime.gopark: blocked until the callback function unlockf in argument list return false.
// runtime.notesleep: hibernate the thread.
// runtime.Goexit: stop the execution of goroutine immediately and call defer, but it will not cause panic.
