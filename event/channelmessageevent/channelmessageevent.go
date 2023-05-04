package channelmessagevent

import "github.com/go-nostr/nostr/event"

// Event for posting messages in a channel
const Kind = 42

// New creates a new channel message event.
func New() *event.Event {
	return event.New(Kind, "")
}
