package channelmuteuserevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for muting a user in a channel
const Kind = 44

// New creates a new ChannelMuteUserEvent.
func New() *ChannelMuteUserEvent {
	event := &ChannelMuteUserEvent{&event.Event{}}
	event.Set("kind", Kind)
	return event
}

// ChannelMuteUser represents a user mute event in a channel.
type ChannelMuteUserEvent struct {
	*event.Event
}
