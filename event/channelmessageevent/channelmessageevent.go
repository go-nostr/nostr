package channelmessagevent

import "github.com/go-nostr/nostr/event"

// Event for posting messages in a channel
const Kind = 42

// New creates a new channel message event.
func New() *ChannelMessageEvent {
	event := &ChannelMessageEvent{}
	event.Set("kind", Kind)
	return event
}

// ChannelMessage TODO
type ChannelMessageEvent struct {
	*event.Event
}
