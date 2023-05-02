package pinlistevent

import "github.com/go-nostr/nostr/event"

// Event for managing a pin list
const Kind = 10001

// New creates a new PinListEvent.
func New() *PinListEvent {
	event := &PinListEvent{&event.Event{}}
	event.Set("kind", Kind)
	return event
}

// PinList represents a pin list event.
type PinListEvent struct {
	*event.Event
}
