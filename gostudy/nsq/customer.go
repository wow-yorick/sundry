package main

import (
	"log"
	"sync"

	"github.com/nsqio/go-nsq"
)

func main() {
	NSQDsAddrs := []string{"10.64.146.231:4150"}

	wgReceivers := &sync.WaitGroup{}
	wgReceivers.Add(3)

	go consumer1(NSQDsAddrs, wgReceivers)
	go consumer2(NSQDsAddrs, wgReceivers)
	go consumer3(NSQDsAddrs, wgReceivers)
	wgReceivers.Wait()
	println("finish")
}

func consumer1(NSQDsAddrs []string, wg *sync.WaitGroup) {
	defer wg.Done()
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("yorick-test", "sensor01", cfg)
	if err != nil {
		log.Fatal(err)
	}
	consumer.ChangeMaxInFlight(2)
	consumer.AddConcurrentHandlers()
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(
		func(message *nsq.Message) error {
			log.Println(string(message.Body) + " C1")
			return nil
		}), 2)
	if err := consumer.ConnectToNSQDs(NSQDsAddrs); err != nil {
		log.Fatal(err, " C1")
	}
	<-consumer.StopChan
}

func consumer2(NSQDsAddrs []string, wg *sync.WaitGroup) {
	defer wg.Done()
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("yorick-test", "sensor01", cfg)
	if err != nil {
		log.Fatal(err)
	}
	//设置每次处理的最大数据
	consumer.ChangeMaxInFlight(2)
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(
		func(message *nsq.Message) error {
			log.Println(string(message.Body) + " C2")
			return nil
		}), 2)
	if err := consumer.ConnectToNSQDs(NSQDsAddrs); err != nil {
		log.Fatal(err, " C2")
	}
	<-consumer.StopChan
}

func consumer3(NSQDsAddrs []string, wg *sync.WaitGroup) {
	defer wg.Done()
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("yorick-test", "sensor02", cfg)
	if err != nil {
		log.Fatal(err)
	}
	//设置每次处理的最大数据
	consumer.ChangeMaxInFlight(2)
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(
		func(message *nsq.Message) error {
			log.Println(string(message.Body) + " C3")
			return nil
		}), 2)
	if err := consumer.ConnectToNSQDs(NSQDsAddrs); err != nil {
		log.Fatal(err, " C3")
	}
	<-consumer.StopChan
}
