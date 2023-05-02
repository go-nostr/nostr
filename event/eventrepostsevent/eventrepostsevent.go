package eventrepostsevent

import "github.com/go-nostr/nostr/event"

// Event for signaling to followers that another event is worth reading
const Kind = 6

// NewEventRepostsEvent creates a new event reposts event.
func NewEventRepostsEvent() *EventRepostsEvent {
	event := &EventRepostsEvent{&event.Event{}}
	event.Set("kind", Kind)
	return event
}

// EventRepostsEvent represents an event reposts event.
type EventRepostsEvent struct {
	*event.Event
}
