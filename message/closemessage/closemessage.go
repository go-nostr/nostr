package closemessage

import "github.com/go-nostr/nostr/message"

// Type TBD
const Type = "CLOSE"

// New creates a new CloseMessage.
func New() message.Message {
	return message.New(Type)
}
