package clientauthenticationevent

import (
	"github.com/go-nostr/nostr/event"
)

// Event for client authentication process
const Kind = 22242

// New creates a new ClientAuthenticationEvent.
func New() *ClientAuthenticationEvent {
	event := &ClientAuthenticationEvent{&event.Event{}}
	event.Set("kind", Kind)
	return event
}

// ClientAuthentication represents a client authentication event.
type ClientAuthenticationEvent struct {
	*event.Event
}
