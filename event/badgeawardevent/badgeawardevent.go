package badgeawardevent

import "github.com/go-nostr/nostr/event"

// Kind for awarding badges to users
const Kind = 8

type Options struct {
}

// New creates a new badge award event.
func New(opt *Options) *BadgeAwardEvent {
	event := &BadgeAwardEvent{}
	event.Set("kind", Kind)
	return event
}

// BadgeAwardEvent represents a badge award event.
type BadgeAwardEvent struct {
	*event.Event
}
