package nostr

import (
	"encoding/json"
)

type MessageType string

const (
	// MessageTypeAuth represents an authentication message type.
	MessageTypeAuth MessageType = "AUTH"
	// MessageTypeClose represents a Close message type.
	MessageTypeClose MessageType = "CLOSE"
	// MessageTypeCount represents a count message type, usually for counting events.
	MessageTypeCount MessageType = "COUNT"
	// MessageTypeEndOfStoredEvents represents an End of Stored Events message type.
	MessageTypeEOSE MessageType = "EOSE"
	// MessageTypeEvent represents an Event message type.
	MessageTypeEvent MessageType = "EVENT"
	// MessageTypeNotice represents a Notice message type, usually for notifications.
	MessageTypeNotice MessageType = "NOTICE"
	// MessageTypeOk represents a success confirmation message type.
	MessageTypeOk MessageType = "OK"
	// MessageTypeRequest represents a Request message type.
	MessageTypeRequest MessageType = "REQ"
)

// MessageHandlerFunc ...
type MessageHandlerFunc func(mess Message)

// Handle ...
func (f MessageHandlerFunc) Handle(mess Message) {
	f(mess)
}

type MessageHandler interface {
	Handle(mess Message)
}

// Message is an interface for encoding and marshaling messages.
type Message interface {
	Marshal() ([]byte, error)
	Unmarshal(data []byte) error
}

// AuthMessage TBD
type AuthMessage struct {
	Type      MessageType `json:"type,omitempty"`
	Challenge string      `json:"subscription_id,omitempty"`
	Event     *Event      `json:"event,omitempty"`
}

// NewAuthMessage TBD
func NewAuthMessage(challenge string, event *Event) Message {
	return &AuthMessage{
		Type:      MessageTypeAuth,
		Challenge: challenge,
		Event:     event,
	}
}

// Marshal TBD
func (m *AuthMessage) Marshal() ([]byte, error) {
	args := []interface{}{
		MessageTypeAuth,
	}
	if m.Challenge != "" {
		args = append(args, m.Challenge)
	} else if m.Event != nil {
		args = append(args, m.Event)
	}
	return json.Marshal(args)
}

// Unmarshal TBD
func (m *AuthMessage) Unmarshal(data []byte) error {
	var args []json.RawMessage
	if err := json.Unmarshal(data, &args); err != nil {
		return err
	}
	if err := json.Unmarshal(args[0], &m.Type); err != nil {
		return err
	}
	if len(args[1]) > 0 && args[1][0] == '{' {
		if err := json.Unmarshal(args[1], &m.Event); err != nil {
			return err
		}
	} else {
		if err := json.Unmarshal(args[1], &m.Challenge); err != nil {
			return err
		}
	}
	return nil
}

// CloseMessage TBD
type CloseMessage struct {
	Type           MessageType `json:"type,omitempty"`
	SubscriptionID string      `json:"subscription_id,omitempty"`
}

// NewCloseMessage TBD
func NewCloseMessage(subscriptionID string) *CloseMessage {
	return &CloseMessage{
		Type:           MessageTypeEvent,
		SubscriptionID: subscriptionID,
	}
}

// Marshal TBD
func (m *CloseMessage) Marshal() ([]byte, error) {
	return json.Marshal([]interface{}{
		MessageTypeClose,
		m.SubscriptionID,
	})
}

// Unmarshal TBD
func (m *CloseMessage) Unmarshal(data []byte) error {
	var args []json.RawMessage
	if err := json.Unmarshal(data, &args); err != nil {
		return err
	}
	if err := json.Unmarshal(args[0], &m.Type); err != nil {
		return err
	}
	if err := json.Unmarshal(args[1], &m.SubscriptionID); err != nil {
		return err
	}
	return nil
}

// CountMessage TBD
type CountMessage struct {
	Type           MessageType `json:"type,omitempty"`
	SubscriptionID string      `json:"subscription_id,omitempty"`
	Count          *Count      `json:"count,omitempty"`
	Filters        []Filter    `json:"filter,omitempty"`
}

// NewCountMessage TBD
func NewCountMessage(subscriptionID string, count *Count, filters ...Filter) Message {
	return &CountMessage{
		Type:           MessageTypeEvent,
		SubscriptionID: subscriptionID,
		Count:          count,
		Filters:        filters,
	}
}

// Marshal TBD
func (m *CountMessage) Marshal() ([]byte, error) {
	args := []interface{}{
		MessageTypeCount,
		m.SubscriptionID,
	}
	if m.Count != nil {
		args = append(args, m.Count)
	} else {
		for _, f := range m.Filters {
			args = append(args, f)
		}
	}
	return json.Marshal(args)
}

// Unmarshal TBD
func (m *CountMessage) Unmarshal(data []byte) error {
	var args []json.RawMessage
	if err := json.Unmarshal(data, &args); err != nil {
		return err
	}
	if err := json.Unmarshal(args[0], &m.Type); err != nil {
		return err
	}
	if err := json.Unmarshal(args[1], &m.SubscriptionID); err != nil {
		return err
	}
	if err := json.Unmarshal(args[2], &m.Count); err == nil {
		return nil
	}
	for i, v := range args[2:] {
		if err := json.Unmarshal(v, &m.Filters[i]); err != nil {
			return err
		}
	}
	return nil
}

// EOSEMessage TBD
type EOSEMessage struct {
	Type           MessageType `json:"type,omitempty"`
	SubscriptionID string      `json:"subscription_id,omitempty"`
}

// NewEOSEMessage TBD
func NewEOSEMessage(subscriptionID string) Message {
	return &EOSEMessage{
		Type:           MessageTypeEvent,
		SubscriptionID: subscriptionID,
	}
}

// Marshal TBD
func (m *EOSEMessage) Marshal() ([]byte, error) {
	args := []interface{}{
		MessageTypeEOSE,
		m.SubscriptionID,
	}
	return json.Marshal(args)
}

// Unmarshal TBD
func (m *EOSEMessage) Unmarshal(data []byte) error {
	var args []json.RawMessage
	if err := json.Unmarshal(data, &args); err != nil {
		return err
	}
	if err := json.Unmarshal(args[0], &m.Type); err != nil {
		return err
	}
	if err := json.Unmarshal(args[1], &m.SubscriptionID); err != nil {
		return err
	}
	return nil
}

// EventMessage TBD
type EventMessage struct {
	Type           MessageType `json:"type,omitempty"`
	SubscriptionID string      `json:"subscription_id,omitempty"`
	Event          *Event      `json:"event,omitempty"`
}

// NewEventMessage TBD
func NewEventMessage(subscriptionID string, event *Event) Message {
	return &EventMessage{
		Type:           MessageTypeEvent,
		SubscriptionID: subscriptionID,
		Event:          event,
	}
}

// Marshal marshals the baseMessage into a JSON byte array.
func (m *EventMessage) Marshal() ([]byte, error) {
	args := []interface{}{
		m.Type,
	}
	if m.SubscriptionID != "" {
		args = append(args, m.SubscriptionID)
	}
	if m.Event != nil {
		args = append(args, m.Event)
	}
	return json.Marshal(args)
}

// Unmarshal unmarshals a JSON byte array into an EventMessage.
func (m *EventMessage) Unmarshal(data []byte) error {
	var args []json.RawMessage
	if err := json.Unmarshal(data, &args); err != nil {
		return err
	}
	if err := json.Unmarshal(args[0], &m.Type); err != nil {
		return err
	}
	if len(args) > 2 {
		if err := json.Unmarshal(args[1], &m.SubscriptionID); err != nil {
			return err
		}
		if err := json.Unmarshal(args[2], &m.Event); err != nil {
			return err
		}
	} else {
		if err := json.Unmarshal(args[1], &m.Event); err != nil {
			return err
		}
	}
	return nil
}

// NoticeMessage TBD
type NoticeMessage struct {
	Type   MessageType `json:"type,omitempty"`
	Notice string      `json:"notice,omitempty"`
}

// NewNoticeMessage TBD
func NewNoticeMessage(notice string) Message {
	return &NoticeMessage{
		Type:   MessageTypeNotice,
		Notice: notice,
	}
}

// Marshal TBD
func (m *NoticeMessage) Marshal() ([]byte, error) {
	return json.Marshal([]interface{}{
		m.Type,
		m.Notice,
	})
}

// Unmarshal TBD
func (m *NoticeMessage) Unmarshal(data []byte) error {
	var args []json.RawMessage
	if err := json.Unmarshal(data, &args); err != nil {
		return err
	}
	if err := json.Unmarshal(args[0], &m.Type); err != nil {
		return err
	}
	if err := json.Unmarshal(args[1], &m.Notice); err != nil {
		return err
	}
	return nil
}

// OkMessage TBD
type OkMessage struct {
	Type    MessageType `json:"type,omitempty"`
	EventID string      `json:"event_id,omitempty"`
	IsSaved bool        `json:"is_saved,omitempty"`
	Message string      `json:"message,omitempty"`
}

// NewOkMessage TBD
func NewOkMessage(eventID string, isSaved bool, message string) Message {
	return &OkMessage{
		Type:    MessageTypeOk,
		EventID: eventID,
		IsSaved: isSaved,
		Message: message,
	}
}

// Marshal TBD
func (m *OkMessage) Marshal() ([]byte, error) {
	return json.Marshal([]interface{}{
		m.Type,
		m.EventID,
		m.IsSaved,
		m.Message,
	})
}

// Unmarshal TBD
func (m *OkMessage) Unmarshal(data []byte) error {
	var args []json.RawMessage
	if err := json.Unmarshal(data, &args); err != nil {
		return err
	}
	if err := json.Unmarshal(args[0], &m.Type); err != nil {
		return err
	}
	if err := json.Unmarshal(args[1], &m.EventID); err != nil {
		return err
	}
	if err := json.Unmarshal(args[2], &m.IsSaved); err != nil {
		return err
	}
	if err := json.Unmarshal(args[3], &m.Message); err != nil {
		return err
	}
	return nil
}

// RequestMessage TBD
type RequestMessage struct {
	Type           MessageType `json:"type,omitempty"`
	SubscriptionID string      `json:"subscription_id,omitempty"`
	Filter         *Filter     `json:"filter,omitempty"`
}

// NewRequestMessage TBD
func NewRequestMessage(subscriptionID string, filter *Filter) Message {
	return &RequestMessage{
		Type:           MessageTypeRequest,
		SubscriptionID: subscriptionID,
		Filter:         filter,
	}
}

// Marshal TBD
func (m *RequestMessage) Marshal() ([]byte, error) {
	return json.Marshal([]interface{}{
		m.Type,
		m.SubscriptionID,
		m.Filter,
	})
}

// Unmarshal TBD
func (m *RequestMessage) Unmarshal(data []byte) error {
	var args []json.RawMessage
	if err := json.Unmarshal(data, &args); err != nil {
		return err
	}
	if err := json.Unmarshal(args[0], &m.Type); err != nil {
		return err
	}
	if err := json.Unmarshal(args[1], &m.SubscriptionID); err != nil {
		return err
	}
	if err := json.Unmarshal(args[2], &m.Filter); err != nil {
		return err
	}
	return nil
}
