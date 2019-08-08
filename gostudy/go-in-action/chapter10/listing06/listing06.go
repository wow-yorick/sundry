package main

import "sundry/gostudy/go-in-action/chapter10/listing06/pubsub"

type publisher interface {
	Publish(key string, v interface{}) error
	Subscribe(key string) error
}

type mock struct{}

func (m *mock) Publish(key string, v interface{}) error {
	return nil
}

func (m *mock) Subscribe(key string) error {
	return nil
}

func main() {
	pubs := []publisher{
		pubsub.New("localhost"),
		&mock{},
	}

	for _, p := range pubs {
		p.Publish("key", "value")
		p.Subscribe("key")
	}
}
