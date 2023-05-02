package badgedefinitionevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for defining badges
const Kind = 30009

type Options struct{}

// New creates a new BadgeDefinitionEvent.
func New(opt *Options) *BadgeDefinitionEvent {
	return &BadgeDefinitionEvent{
		Event: event.New(&event.Options{
			Kind: Kind,
		}),
	}
}

// BadgeDefinition represents a badge definition event.
type BadgeDefinitionEvent struct {
	*event.Event
}
