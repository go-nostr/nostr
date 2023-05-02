package eosemessage

import (
	"encoding/json"

	"github.com/go-nostr/nostr/message"
)

const Type = "EOSE"

// New creates a new EOSEMessage.
func New() *EOSEMessage {
	mess := &EOSEMessage{&message.Message{}}
	mess.Push(Type)
	return mess
}

// EOSEMessage is a specialized message type for indicating the end of stored events.
type EOSEMessage struct {
	*message.Message
}

// SubscriptionID returns the subscription ID contained in the RequestMessage.
func (m *EOSEMessage) SubscriptionID() []byte {
	if len(*m.Message) < 2 {
		return []byte{}
	}
	var sid string
	if err := json.Unmarshal((*m.Message)[1], &sid); err != nil {
		return []byte{}
	}
	return []byte(sid)
}
