package channelcreationevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for creating new channels
const Kind = 40

// New creates a new channel creation event.
func New() *ChannelCreationEvent {
	event := &ChannelCreationEvent{&event.Event{}}
	event.Set("kind", Kind)
	return event
}

// ChannelCreationEvent represents a channel creation event.
type ChannelCreationEvent struct {
	*event.Event
}
