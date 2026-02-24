package pubsub

import (
	"sync"

	"github.com/johmara/openclaude/internal/claude"
)

// Broker is a generic pub/sub broker for claude events.
type Broker struct {
	mu          sync.RWMutex
	subscribers map[string]chan claude.Event
}

// NewBroker creates a new event broker.
func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[string]chan claude.Event),
	}
}

// Subscribe creates a new subscription channel with the given ID.
// Returns a buffered channel that receives events.
func (b *Broker) Subscribe(id string) <-chan claude.Event {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan claude.Event, 128)
	b.subscribers[id] = ch
	return ch
}

// Unsubscribe removes a subscription and closes its channel.
func (b *Broker) Unsubscribe(id string) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if ch, ok := b.subscribers[id]; ok {
		close(ch)
		delete(b.subscribers, id)
	}
}

// Publish sends an event to all subscribers.
func (b *Broker) Publish(evt claude.Event) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for _, ch := range b.subscribers {
		select {
		case ch <- evt:
		default:
			// Drop events if subscriber is slow
		}
	}
}

// PublishStream reads events from a source channel and publishes them.
func (b *Broker) PublishStream(source <-chan claude.Event) {
	for evt := range source {
		b.Publish(evt)
	}
}
