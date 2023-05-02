package eventmessage

import (
	"encoding/json"

	"github.com/go-nostr/nostr"
	"github.com/go-nostr/nostr/event"
	"github.com/go-nostr/nostr/message"
)

const Type = "EVENT"

// New creates a new EventMessage.
func New(sid string, evnt nostr.Event) *EventMessage {
	mess := &EventMessage{&message.Message{}}
	mess.Push(Type)
	if sid != "" {
		mess.Push(sid)
	}
	if evnt != nil {
		mess.Push(evnt)
	}
	return mess
}

// EventMessage is a specialized message type for handling events.
type EventMessage struct {
	*message.Message
}

// Event returns the event contained in the EventMessage.
func (m *EventMessage) Event() nostr.Event {
	if len(*m.Message) < 3 {
		return nil
	}
	var evnt event.Event
	if err := json.Unmarshal((*m.Message)[2], &evnt); err != nil {
		return nil
	}
	return &evnt
}

// SubscriptionID returns the subscription ID contained in the EventMessage.
func (m *EventMessage) SubscriptionID() string {
	if len(*m.Message) < 3 {
		return ""
	}
	var sid string
	if err := json.Unmarshal((*m.Message)[1], &sid); err != nil {
		return ""
	}
	return sid
}
