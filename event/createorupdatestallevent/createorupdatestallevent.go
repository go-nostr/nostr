package createorupdatestallevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for creating or updating a stall
const Kind = 30017

// New creates a new CreateOrUpdateStallEvent.
func New() *CreateOrUpdateStallEvent {
	event := &CreateOrUpdateStallEvent{}
	event.Set("kind", Kind)
	return event
}

// CreateOrUpdateStall represents a create or update stall event.
type CreateOrUpdateStallEvent struct {
	*event.Event
}
