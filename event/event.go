package event

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/go-nostr/nostr/tag"
	"github.com/go-nostr/nostr/tag/rawtag"
)

type Options struct {
	Kind int
}

// New creates a new empty raw event.
func New(opt *Options) *Event {
	evnt := &Event{}
	if opt != nil {
		evnt.Set("kind", opt.Kind)
	}
	return evnt
}

// Event is a map-based representation of an event with keys and values
// encoded as json.RawMessage.
type Event map[string]json.RawMessage

// Content retrieves the content of the Event as a byte slice.
func (e *Event) Content() []byte {
	var content string
	if err := json.Unmarshal((*e)["content"], &content); err != nil {
		return []byte{}
	}
	return []byte(content)
}

// CreatedAt TBD
func (e *Event) CreatedAt() int {
	var createdAt int
	if err := json.Unmarshal((*e)["created_at"], &createdAt); err != nil {
		return -1
	}
	return createdAt
}

// Get retrieves the value associated with the given key from the Event.
func (e *Event) Get(key string) any {
	var val any
	if err := json.Unmarshal((*e)[key], &val); err != nil {
		return nil
	}
	return val
}

// ID retrieves the ID of the Event as a byte slice.
func (e *Event) ID() []byte {
	h := sha256.Sum256(e.Serialize())
	return []byte(hex.EncodeToString(h[:]))
}

// Keys returns a list of keys in the Event.
func (e *Event) Keys() []string {
	var keys []string
	for key := range *e {
		keys = append(keys, key)
	}
	return keys
}

// Kind retrieves the kind of the Event as an integer.
func (e *Event) Kind() int {
	var kind int
	if err := json.Unmarshal((*e)["kind"], &kind); err != nil {
		return -1
	}
	return kind
}

// Marshal marshals the Event to a byte slice.
func (e *Event) Marshal() ([]byte, error) {
	return e.MarshalJSON()
}

// MarshalJSON marshals the Event to a byte slice.
func (e *Event) MarshalJSON() ([]byte, error) {
	byt := make([]byte, 0)
	byt = append(byt, '{')
	len := len(e.Keys())
	for i, k := range e.Keys() {
		data, err := json.Marshal((*e)[k])
		if err != nil {
			return nil, err
		}
		byt = append(byt, []byte(fmt.Sprintf("\"%s\":%s", k, data))...)
		if i < len-1 {
			byt = append(byt, ',')
		}
	}
	byt = append(byt, '}')
	return byt, nil
}

// PublicKey retrieves the public key of the Event as a byte slice.
func (e *Event) PublicKey() []byte {
	var pubKey string
	if err := json.Unmarshal((*e)["pubkey"], &pubKey); err != nil {
		return []byte{}
	}
	return []byte(pubKey)
}

// Serialize TBD
func (e *Event) Serialize() []byte {
	byt := make([]byte, 0)
	byt = append(byt, []byte("[0")...)
	byt = append(byt, ',')
	byt = append(byt, []byte(fmt.Sprintf("\"%s\"", e.PublicKey()))...)
	byt = append(byt, ',')
	byt = append(byt, []byte(fmt.Sprintf("%d", e.CreatedAt()))...)
	byt = append(byt, ',')
	byt = append(byt, []byte(fmt.Sprintf("%d", e.Kind()))...)
	byt = append(byt, ',')
	tags, _ := json.Marshal(e.Tags())
	byt = append(byt, tags...)
	byt = append(byt, ',')
	byt = append(byt, []byte(fmt.Sprintf("\"%s\"", e.Content()))...)
	byt = append(byt, ']')
	return byt
}

// Set associates a value with a given key in the Event.
func (e *Event) Set(key string, val any) error {
	data, err := json.Marshal(val)
	if err != nil {
		return err
	}
	(*e)[key] = data
	return nil
}

// Signature retrieves the signature of the Event as a byte slice.
func (e *Event) Signature() []byte {
	var sig string
	if err := json.Unmarshal((*e)["sig"], &sig); err != nil {
		return []byte{}
	}
	return []byte(sig)
}

// Sign signs the Event.
func (e *Event) Sign(privateKey string) error {
	s, err := hex.DecodeString(privateKey)
	if err != nil {
		return fmt.Errorf("Sign called with invalid private key '%s': %w", privateKey, err)
	}
	e.Set("created_at", time.Now().Unix())
	prvKey, pubKey := btcec.PrivKeyFromBytes(s)
	e.Set("pubkey", hex.EncodeToString(pubKey.SerializeCompressed()[1:]))
	h := sha256.Sum256(e.Serialize())
	sig, err := schnorr.Sign(prvKey, h[:])
	if err != nil {
		return err
	}
	if e.Get("tags") == nil {
		e.Set("tags", []rawtag.RawTag{})
	}
	e.Set("id", hex.EncodeToString(h[:]))
	e.Set("sig", hex.EncodeToString(sig.Serialize()))
	return nil
}

// Tags retrieves the list of tags associated with the Event.
func (e *Event) Tags() []tag.Tag {
	tags := make([]tag.Tag, 0)
	var args []json.RawMessage
	if err := json.Unmarshal((*e)["tags"], &args); err != nil {
		return tags
	}
	for _, arg := range args {
		var tag rawtag.RawTag
		if err := tag.Unmarshal(arg); err != nil {
			return tags
		}
		tags = append(tags, &tag)
	}
	return tags
}

// Unmarshal unmarshals the Event from a byte slice.
func (e *Event) Unmarshal(data []byte) error {
	return e.UnmarshalJSON(data)
}

// UnmarshalJSON TBD
func (e *Event) UnmarshalJSON(data []byte) error {
	evnt := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &evnt); err != nil {
		return err
	}
	*e = evnt
	return nil
}

// Validate validates the Event.
func (e *Event) Validate() error {
	pubKeyHex := make([]byte, hex.DecodedLen(len(e.PublicKey())))
	if _, err := hex.Decode(pubKeyHex, e.PublicKey()); err != nil {
		return nil
	}
	pubKey, err := schnorr.ParsePubKey(pubKeyHex)
	if err != nil {
		return err
	}
	sigHex := make([]byte, hex.DecodedLen((len(e.Signature()))))
	if _, err := hex.Decode(sigHex, e.Signature()); err != nil {
		return err
	}
	if err != nil {
		return err
	}
	sig, err := schnorr.ParseSignature(sigHex)
	if err != nil {
		return err
	}
	hash := sha256.Sum256(e.Serialize())
	if !sig.Verify(hash[:], pubKey) {
		return fmt.Errorf("can't verify")
	}
	return nil
}

// Values returns a list of values in the Event.
func (e *Event) Values() []any {
	var vals []any
	for _, v := range *e {
		vals = append(vals, v)
	}
	return vals
}
