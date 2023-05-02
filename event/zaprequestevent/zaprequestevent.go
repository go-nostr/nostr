package zaprequestevent

import "github.com/go-nostr/nostr/event"

// Kind for requesting a Zap action
const Kind = 9734

// New creates a new ZapRequestEvent.
func New() *ZapRequestEvent {
	event := &ZapRequestEvent{}
	event.Set("kind", Kind)
	return event
}

// ZapRequest represents a zap request event.
type ZapRequestEvent struct {
	*event.Event
}
