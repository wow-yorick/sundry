package main

import (
	"log"
	"strconv"

	"github.com/nsqio/go-nsq"
)

func main() {
	cfg := nsq.NewConfig()
	nsqdAddr := "10.64.146.231:4150"
	producer, err := nsq.NewProducer(nsqdAddr, cfg)
	if err != nil {
		log.Fatal(err)
	}
	for i:=0;i< 5;i++ {
		is := strconv.Itoa(i)
		msg := "test-"+ is
		if err := producer.Publish("yorick-test", []byte(msg)); err != nil {
			log.Fatal("publish error: " + err.Error())
		}
		log.Printf("publish msg: %v \n", msg)
	}
	log.Println("finish")
}