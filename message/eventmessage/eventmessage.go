package eventmessage

import (
	"github.com/go-nostr/nostr/event"
	"github.com/go-nostr/nostr/message"
)

const Type = "EVENT"

// New creates a new EventMessage.
func New(subscriptionID string, evnt *event.Event) message.Message {
	return message.New(Type, subscriptionID, evnt)
}
