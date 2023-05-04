package eventdeletionevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for deleting events
const Kind = 5

// New creates a new event deletion event.
func New(content string) *event.Event {
	return event.New(Kind, content)
}
