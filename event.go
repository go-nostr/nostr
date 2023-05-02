package nostr

import (
	"encoding/json"

	"github.com/go-nostr/nostr/tag"
)

// Event represents an abstract event interface that can be extended to
// handle various kinds of events.
type Event interface {
	json.Marshaler
	json.Unmarshaler
	Content() []byte
	CreatedAt() int
	Get(key string) any
	ID() []byte
	Keys() []string
	Kind() int
	Marshal() ([]byte, error)
	PublicKey() []byte
	Serialize() []byte
	Set(key string, val any) error
	Sign(privateKey string) error
	Signature() []byte
	Tags() []tag.Tag
	Unmarshal(data []byte) error
	Validate() error
	Values() []any
}
