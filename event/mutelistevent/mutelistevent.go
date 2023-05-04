package mutelistevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for managing a mute list
const Kind = 10000

// New creates a new MuteListEvent.
func New() *event.Event {
	return event.New(Kind, "")
}
