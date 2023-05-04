package zapevent

import "github.com/go-nostr/nostr/event"

// Kind for performing a Zap action
const Kind = 9735

// New creates a new ZapEvent.
func New() *event.Event {
	return event.New(Kind, "")
}
