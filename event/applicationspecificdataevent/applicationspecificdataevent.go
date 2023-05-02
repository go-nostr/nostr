package applicationspecificdataevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for managing application-specific data
const Kind = 30078

type Options struct{}

// New creates a new Event.
func New(opt *Options) *ApplicationSpecificDataEvent {
	event := &ApplicationSpecificDataEvent{}
	event.Set("kind", Kind)
	return event
}

type ApplicationSpecificDataEvent struct {
	*event.Event
}
