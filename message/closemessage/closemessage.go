package closemessage

import "github.com/go-nostr/nostr/message"

// Type TBD
const Type = "CLOSE"

// New creates a new CloseMessage.
func New() *CloseMessage {
	mess := &CloseMessage{&message.Message{}}
	mess.Push(Type)
	return mess
}

// CloseMessage is a specialized message type for closing a connection.
type CloseMessage struct {
	*message.Message
}
