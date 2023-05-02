package channelmetadataevent

import "github.com/go-nostr/nostr/event"

// Kind for setting channel metadata
const Kind = 41

// New creates a new channel metadata event.
func New() *ChannelMetadataEvent {
	event := &ChannelMetadataEvent{}
	event.Set("kind", Kind)
	return event
}

// ChannelMetadataEvent represents a channel metadata event.
type ChannelMetadataEvent struct {
	*event.Event
}
