package okmessage

import "github.com/go-nostr/nostr/message"

const Type = "OK"

// New creates a new OkMessage.
func New() *OkMessage {
	mess := &OkMessage{&message.Message{}}
	mess.Push(Type)
	return mess
}

// OkMessage is a specialized message type for indicating success or confirmation.
type OkMessage struct {
	*message.Message
}
