package reportingevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for reporting content or users
const Kind = 1984

// New creates a new ReportingEvent.
func New() *ReportingEvent {
	event := &ReportingEvent{}
	event.Set("kind", Kind)
	return event
}

// Reporting represents a reporting event.
type ReportingEvent struct {
	*event.Event
}
