package encrpyteddirectmessagesevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for encrypted direct messages
const Kind = 4

// NewEncryptedDirectMessagesEvent creates a new encrypted direct messages event.
func New() *event.Event {
	event := &event.Event{}
	event.Set("kind", Kind)
	return event
}
