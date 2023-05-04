package reactionevent

import "github.com/go-nostr/nostr/event"

// Kind for reacting to other notes
const Kind = 7

// New creates a new reaction event.
func New() *event.Event {
	return event.New(Kind, "")
}
