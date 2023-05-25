package authmessage

import (
	"github.com/go-nostr/nostr/event"
	"github.com/go-nostr/nostr/message"
)

// Type for message "AUTH"
const Type = "AUTH"

// New creates a new "AUTH" message
func New(challenge string, evt *event.Event) message.Message {
	if evt != nil {
		return message.New(Type, evt)
	}
	return message.New(Type, challenge)
}
