package pubsub

type Pubsub struct {
}

type Event interface{}

type MessageQueue interface {
	Publish(eventType string, event Event) error
	Subscribe(eventType string, subscriber Subscriber) error
}

type Subscriber interface {
	Notify(event Event) error
}
