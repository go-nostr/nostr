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
	Validate() error
}

// BaseMessage TBD
type BaseMessage []interface{}

// Marshal TBD
func (m *BaseMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal TBD
func (m *BaseMessage) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Validate TBD
func (m *BaseMessage) Validate() error {
	return nil
}

// AuthMessage TBD
type AuthMessage struct {
	BaseMessage
}

// NewAuthMessage TBD
func NewAuthMessage(challenge string, event Event) Message {
	if challenge != "" {
		return &AuthMessage{BaseMessage{string(MessageTypeAuth), challenge}}
	}
	return &AuthMessage{BaseMessage{string(MessageTypeAuth), event}}
}

// Marshal TBD
func (m *AuthMessage) Marshal() ([]byte, error) {
	return m.BaseMessage.Marshal()
}

// Unmarshal TBD
func (m *AuthMessage) Unmarshal(data []byte) error {
	return m.BaseMessage.Unmarshal(data)
}

// Validate TBD
func (m *AuthMessage) Validate() error {
	return nil
}

// CloseMessage TBD
type CloseMessage struct {
	BaseMessage
}

// NewCloseMessage TBD
func NewCloseMessage(subscriptionID string) Message {
	return &CloseMessage{BaseMessage{string(MessageTypeClose), subscriptionID}}
}

// Marshal TBD
func (m *CloseMessage) Marshal() ([]byte, error) {
	return m.BaseMessage.Marshal()
}

// Unmarshal TBD
func (m *CloseMessage) Unmarshal(data []byte) error {
	return m.BaseMessage.Unmarshal(data)
}

// Validate TBD
func (m *CloseMessage) Validate() error {
	return nil
}

// CountMessage TBD
type CountMessage struct {
	BaseMessage
}

// NewCountMessage TBD
func NewCountMessage(subscriptionID string, count *Count, filters ...Filter) Message {
	return &CountMessage{BaseMessage{string(MessageTypeCount), subscriptionID, count}}
}

// Marshal TBD
func (m *CountMessage) Marshal() ([]byte, error) {
	return m.BaseMessage.Marshal()
}

// Unmarshal TBD
func (m *CountMessage) Unmarshal(data []byte) error {
	return m.BaseMessage.Unmarshal(data)
}

// Validate TBD
func (m *CountMessage) Validate() error {
	return nil
}

// EOSEMessage TBD
type EOSEMessage struct {
	BaseMessage
}

// NewEOSEMessage TBD
func NewEOSEMessage(subscriptionID string) Message {
	return &EOSEMessage{BaseMessage{string(MessageTypeEOSE), subscriptionID}}
}

// Marshal TBD
func (m *EOSEMessage) Marshal() ([]byte, error) {
	return m.BaseMessage.Marshal()
}

// Unmarshal TBD
func (m *EOSEMessage) Unmarshal(data []byte) error {
	return m.BaseMessage.Unmarshal(data)
}

// Validate TBD
func (m *EOSEMessage) Validate() error {
	return nil
}

// EventMessage TBD
type EventMessage struct {
	BaseMessage
}

// NewEventMessage TBD
func NewEventMessage(subscriptionID string, event *Event) Message {
	return &EventMessage{BaseMessage{string(MessageTypeEvent), subscriptionID, event}}
}

// Marshal marshals the baseMessage into a JSON byte array.
func (m *EventMessage) Marshal() ([]byte, error) {
	return m.BaseMessage.Marshal()
}

// Unmarshal unmarshals a JSON byte array into an EventMessage.
func (m *EventMessage) Unmarshal(data []byte) error {
	return m.BaseMessage.Unmarshal(data)
}

// Validate TBD
func (m *EventMessage) Validate() error {
	return nil
}

// NoticeMessage TBD
type NoticeMessage struct {
	BaseMessage
}

// NewNoticeMessage TBD
func NewNoticeMessage(notice string) Message {
	return &NoticeMessage{BaseMessage{string(MessageTypeNotice), notice}}
}

// Marshal TBD
func (m *NoticeMessage) Marshal() ([]byte, error) {
	return m.BaseMessage.Marshal()
}

// Unmarshal TBD
func (m *NoticeMessage) Unmarshal(data []byte) error {
	return m.BaseMessage.Unmarshal(data)
}

// Validate TBD
func (m *NoticeMessage) Validate() error {
	return nil
}

// OkMessage TBD
type OkMessage struct {
	BaseMessage
}

// NewOkMessage TBD
func NewOkMessage(eventID string, isSaved bool, message string) Message {
	return &OkMessage{BaseMessage{string(MessageTypeOk), eventID, isSaved, message}}
}

// Marshal TBD
func (m *OkMessage) Marshal() ([]byte, error) {
	return m.BaseMessage.Marshal()
}

// Unmarshal TBD
func (m *OkMessage) Unmarshal(data []byte) error {
	return m.BaseMessage.Unmarshal(data)
}

// Validate TBD
func (m *OkMessage) Validate() error {
	return nil
}

// RequestMessage TBD
type RequestMessage struct {
	BaseMessage
}

// NewRequestMessage TBD
func NewRequestMessage(subscriptionID string, filter *Filter) Message {
	return &RequestMessage{BaseMessage{string(MessageTypeRequest), subscriptionID, filter}}
}

// Marshal TBD
func (m *RequestMessage) Marshal() ([]byte, error) {
	return m.BaseMessage.Marshal()
}

// Unmarshal TBD
func (m *RequestMessage) Unmarshal(data []byte) error {
	return m.BaseMessage.Unmarshal(data)
}

// Validate TBD
func (m *RequestMessage) Validate() error {
	return nil
}
