package pubsub

type PubSub struct {
	host string
	// PRETEND THERE ARE MORE FIELDS
}

func New(host string) *PubSub {
	ps := PubSub{
		host: host,
	}

	return &ps
}

func (ps *PubSub) Publish(key string, v interface{}) error {
	return nil
}

func (ps *PubSub) Subscribe(key string) error {
	return nil
}
