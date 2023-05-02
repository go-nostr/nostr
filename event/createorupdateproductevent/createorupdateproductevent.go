package createorupdateproductevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for creating or updating a product
const Kind = 30018

// New creates a new CreateOrUpdateProductEvent.
func New() *CreateOrUpdateProductEvent {
	event := &CreateOrUpdateProductEvent{&event.Event{}}
	event.Set("kind", Kind)
	return event
}

// CreateOrUpdateProduct represents a create or update product event.
type CreateOrUpdateProductEvent struct {
	*event.Event
}
