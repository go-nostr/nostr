package createorupdateproductevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for creating or updating a product
const Kind = 30018

// New creates a new CreateOrUpdateProductEvent.
func New() *event.Event {
	return event.New(Kind, "")
}
