package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const MaxRandomNumber = 100000
	const NumberSenders = 1000

	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(1)

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{}, 1)

	//发送者
	for i := 0; i < NumberSenders; i++ {
		go func() {
			for {
				value := rand.Intn(MaxRandomNumber)
				select {
				case s := <-stopCh:
					fmt.Printf("the:%v\n", s)
					return
				case dataCh <- value:
				}
			}
		}()
	}

	// 接收者
	go func() {
		defer wgReceivers.Done()
		i := 0
		for value := range dataCh {
			//if value == MaxRandomNumber-1 {
			if i == 50 {
				// dataCh通道的接收方是
				//同时也是stopCh cahnnel的发送者。
				//关闭这里的停止通道是安全的。
				close(stopCh)
				return
			}
			log.Println(value)
			i++
		}
	}()

	wgReceivers.Wait()
	println("finish")
}
