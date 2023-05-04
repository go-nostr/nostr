package channelmuteuserevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for muting a user in a channel
const Kind = 44

// New creates a new ChannelMuteUserEvent.
func New() *event.Event {
	return event.New(Kind, "")
}
