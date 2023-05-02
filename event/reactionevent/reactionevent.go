package reactionevent

import "github.com/go-nostr/nostr/event"

// Kind for reacting to other notes
const Kind = 7

// New creates a new reaction event.
func New() *ReactionEvent {
	event := &ReactionEvent{}
	event.Set("kind", Kind)
	return event
}

// ReactionEvent represents a reaction event.
type ReactionEvent struct {
	*event.Event
}
