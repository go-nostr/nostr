package mutelistevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for managing a mute list
const Kind = 10000

// New creates a new MuteListEvent.
func New() *MuteListEvent {
	event := &MuteListEvent{&event.Event{}}
	event.Set("kind", Kind)
	return event
}

// MuteList represents a mute list event.
type MuteListEvent struct {
	*event.Event
}
