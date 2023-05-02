package eventdeletionevent

import "github.com/go-nostr/nostr/event"

// Kind for deleting events
const Kind = 5

// New creates a new event deletion event.
func New() *EventDeletionEvent {
	event := &EventDeletionEvent{&event.Event{}}
	event.Set("kind", Kind)
	return event
}

// EventDeletionEvent represents an event deletion event.
type EventDeletionEvent struct {
	*event.Event
}
