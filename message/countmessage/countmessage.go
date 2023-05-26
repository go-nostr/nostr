package countmessage

import (
	"github.com/go-nostr/nostr/message"
)

const Type = "COUNT"

// New creates a new CountMessage.
func New(subscriptionID string, count *Count) message.Message {
	if count != nil {
		return message.New(Type, count)
	}
	return message.New(Type, subscriptionID)
}
