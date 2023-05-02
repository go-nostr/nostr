package countmessage

import "github.com/go-nostr/nostr/message"

const Type = "COUNT"

// New creates a new CountMessage.
func New() *CountMessage {
	mess := &CountMessage{&message.Message{}}
	mess.Push(Type)
	return mess
}

// CountMessage is a specialized message type for counting events.
type CountMessage struct {
	*message.Message
}
