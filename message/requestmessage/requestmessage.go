package requestmessage

import (
	"encoding/json"

	"github.com/go-nostr/nostr"
	"github.com/go-nostr/nostr/message"
)

const Type = "REQ"

// New creates a new RequestMessage with the provided subscription ID and filters.
func New(sid string, filters ...*nostr.Filter) *RequestMessage {
	mess := &RequestMessage{&message.Message{}}
	mess.Push(Type)
	mess.Push(sid)
	for _, f := range filters {
		mess.Push(f)
	}
	return mess
}

// RequestMessage is a specialized message type for making requests.
type RequestMessage struct {
	*message.Message
}

// SubscriptionID returns the subscription ID contained in the RequestMessage.
func (m *RequestMessage) SubscriptionID() []byte {
	if len(*m.Message) < 2 {
		return []byte{}
	}
	var sid string
	if err := json.Unmarshal((*m.Message)[1], &sid); err != nil {
		return []byte{}
	}
	return []byte(sid)
}
