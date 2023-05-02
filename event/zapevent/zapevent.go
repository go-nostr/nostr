package zapevent

import "github.com/go-nostr/nostr/event"

// Kind for performing a Zap action
const Kind = 9735

// New creates a new ZapEvent.
func New() *ZapEvent {
	event := &ZapEvent{&event.Event{}}
	event.Set("kind", Kind)
	return event
}

// Zap represents a zap event.
type ZapEvent struct {
	*event.Event
}
