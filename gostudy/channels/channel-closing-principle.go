package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

//一个发送者 多个接收者
func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	const MaxRandomNumber = 100000
	const NumReceivers = 100

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)

	dataCh := make(chan int, 100)

	//sender
	go func() {
		for {
			if value := rand.Intn(MaxRandomNumber); value == 0 {
				close(dataCh)
				return
			} else {
				dataCh <- value
			}
		}
	}()

	//receivers
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()

			for value := range dataCh {
				log.Println(value)
			}
		}()
	}

	wgReceivers.Wait()

}
