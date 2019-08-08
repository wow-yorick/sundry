package main

import (
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type (
	semaphore chan struct{}
)

type (
	readerWriter struct {
		name          string
		write         sync.WaitGroup
		readerControl semaphore

		shutdown       chan struct{}
		reportShutdown sync.WaitGroup
		maxReads       int
		maxReaders     int

		currentReads int32
	}
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	log.Println("Starting Process")

	first := start("First", 3, 6)

	second := start("Second", 2, 2)

	time.Sleep(2 * time.Second)

	shutdown(first, second)

	log.Println("Process Ended")
	return
}

func start(name string, maxReads int, maxReaders int) *readerWriter {
	rw := readerWriter{
		name:          name,
		shutdown:      make(chan struct{}),
		maxReads:      maxReads,
		maxReaders:    maxReaders,
		readerControl: make(semaphore, maxReads),
	}

	rw.reportShutdown.Add(maxReaders)
	for goroutine := 0; goroutine < maxReaders; goroutine++ {
		go rw.reader(goroutine)
	}

	rw.reportShutdown.Add(1)
	go rw.writer()

	return &rw
}

func shutdown(readerWriters ...*readerWriter) {
	var waitShutdown sync.WaitGroup
	waitShutdown.Add(len(readerWriters))

	for _, readerWriter := range readerWriters {
		go readerWriter.stop(&waitShutdown)
	}

	waitShutdown.Wait()
}

func (rw *readerWriter) stop(waitShutdown *sync.WaitGroup) {
	defer waitShutdown.Done()

	log.Printf("%s\t:####> Stop", rw.name)

	close(rw.shutdown)

	rw.reportShutdown.Wait()

	log.Printf("%s\t:####> Stopped", rw.name)
}

func (rw *readerWriter) reader(reader int) {
	defer rw.reportShutdown.Done()

	for {
		select {
		case <-rw.shutdown:
			log.Printf("%s\t: #> Reader Shutdown", rw.name)
			return

		default:
			rw.performRead(reader)
		}
	}
}

func (rw *readerWriter) performRead(reader int) {
	rw.ReadLock(reader)

	count := atomic.AddInt32(&rw.currentReads, 1)

	log.Printf("%s\t: [%d] Start\t- [%d] Reads\n", rw.name, reader, count)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)

	count = atomic.AddInt32(&rw.currentReads, -1)
	log.Printf("%s\t: [%d] Finish\t- [%d] Reads\n", rw.name, reader, count)

	rw.ReadUnlock(reader)
}

func (rw *readerWriter) writer() {
	defer rw.reportShutdown.Done()

	for {
		select {
		case <-rw.shutdown:
			log.Printf("%s\t: #> Writer Shutdown", rw.name)
			return
		default:
			rw.performWrite()
		}
	}
}

func (rw *readerWriter) performWrite() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)

	log.Printf("%s\t: ****> Writing Pending\n", rw.name)

	rw.WriteLock()

	log.Printf("%s\t: ****> Writing Start", rw.name)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("%s\t: ****> Writing Finish", rw.name)

	rw.WriteUnlock()
}

func (rw *readerWriter) ReadLock(reader int) {
	rw.write.Wait()

	rw.readerControl.Acquire(1)
}

func (rw *readerWriter) ReadUnlock(reader int) {
	rw.readerControl.Release(1)
}

func (rw *readerWriter) WriteLock() {
	rw.write.Add(1)

	rw.readerControl.Acquire(rw.maxReads)
}

func (rw *readerWriter) WriteUnlock() {
	rw.readerControl.Release(rw.maxReads)

	rw.write.Done()
}

func (s semaphore) Acquire(buffers int) {
	var e struct{}

	for buffer := 0; buffer < buffers; buffer++ {
		s <- e
	}
}

func (s semaphore) Release(buffers int) {
	for buffer := 0; buffer < buffers; buffer++ {
		<-s
	}
}
