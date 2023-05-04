package clientauthenticationevent

import (
	"github.com/go-nostr/nostr/event"
)

// Event for client authentication process
const Kind = 22242

// New creates a new ClientAuthenticationEvent.
func New() *event.Event {
	return event.New(Kind, "")
}
