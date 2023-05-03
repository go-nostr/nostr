package authmessage

import "github.com/go-nostr/nostr/message"

// Type TBD
const Type = "CLOSE"

// New creates a new AuthMessage.
func New() *AuthMessage {
	mess := &AuthMessage{&message.Message{}}
	mess.Push(Type)
	return mess
}

// AuthMessage is a specialized message type for closing a connection.
type AuthMessage struct {
	*message.Message
}
