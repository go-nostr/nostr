package channelmetadataevent

import "github.com/go-nostr/nostr/event"

// Kind for setting channel metadata
const Kind = 41

// New creates a new channel metadata event.
func New() *event.Event {
	return event.New(Kind, "")
}
