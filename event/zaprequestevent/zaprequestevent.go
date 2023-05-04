package zaprequestevent

import "github.com/go-nostr/nostr/event"

// Kind for requesting a Zap action
const Kind = 9734

// New creates a new ZapRequestEvent.
func New() *event.Event {
	return event.New(Kind, "")
}
