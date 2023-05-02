package channelhidemessageevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for hiding messages in a channel
const Kind = 43

// New TODO
func New() *ChannelHideMessageEvent {
	event := &ChannelHideMessageEvent{&event.Event{}}
	event.Set("kind", Kind)
	return event
}

// ChannelHideMessage represents a message hiding event in a channel.
type ChannelHideMessageEvent struct {
	*event.Event
}
