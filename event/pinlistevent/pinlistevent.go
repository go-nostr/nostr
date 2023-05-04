package pinlistevent

import "github.com/go-nostr/nostr/event"

// Event for managing a pin list
const Kind = 10001

// New creates a new PinListEvent.
func New() *event.Event {
	return event.New(Kind, "")
}
