package eventmessage

import (
	"fmt"

	"github.com/go-nostr/nostr/event"
	"github.com/go-nostr/nostr/message"
)

const Type = "EVENT"

// New creates a new EventMessage.
func New(subscriptionID string, evt *event.Event) message.Message {
	if len(subscriptionID) > 64 {
		panic(fmt.Errorf("invalid subscription id"))
	}
	return message.New(Type, subscriptionID, evt)
}
