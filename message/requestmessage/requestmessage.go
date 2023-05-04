package requestmessage

import (
	"github.com/go-nostr/nostr/message"
)

const Type = "REQ"

// New creates a new RequestMessage with the provided subscription ID and filters.
func New(subscriptionID string, filters ...*Filter) message.Message {
	mess := message.New(Type, subscriptionID)
	for _, f := range filters {
		mess.Push(f)
	}
	return mess
}
