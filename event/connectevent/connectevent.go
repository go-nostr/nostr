package connectevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for Nostr client connection
const Kind = 24133

// New creates a new NostrConnectEvent.
func New() *NostrConnectEvent {
	event := &NostrConnectEvent{&event.Event{}}
	event.Set("kind", Kind)
	return event
}

// NostrConnect represents a nostr connect event.
type NostrConnectEvent struct {
	*event.Event
}
