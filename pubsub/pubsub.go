package pubsub

type Pubsub struct {
}

type Event interface{}

type Publisher interface {
	Publish(eventType string, event Event) error
	Register(eventType string, subscriber Subscriber) error
}

type Subscriber interface {
	Notify(event Event) error
}
