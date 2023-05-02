package recommendrelayevent

import "github.com/go-nostr/nostr/event"

const Kind = 2

// New creates a new recommend relay event.
func New() *RecommendRelayEvent {
	event := &RecommendRelayEvent{&event.Event{}}
	event.Set("kind", Kind)
	return event
}

// RecommendRelayEvent represents a recommend relay event.
type RecommendRelayEvent struct {
	*event.Event
}
