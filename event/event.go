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
)

// New creates a new empty raw event.
func New(kind int, content string, tag ...tag.Tag) *Event {
	return &Event{
		Kind:    kind,
		Content: content,
		Tags:    tag,
	}
}

// Event is a representation of an event
type Event struct {
	Content   string    `json:"content,omitempty"`
	CreatedAt int       `json:"created_at,omitempty"`
	ID        string    `json:"id,omitempty"`
	Kind      int       `json:"kind,omitempty"`
	PubKey    string    `json:"pubkey,omitempty"`
	Sig       string    `json:"sig,omitempty"`
	Tags      []tag.Tag `json:"tags,omitempty"`
}

// Marshal marshals the Event to a byte slice.
func (e *Event) Marshal() ([]byte, error) {
	return json.Marshal(*e)
}

// Serialize TBD
func (e *Event) Serialize() []byte {
	if e.Tags == nil {
		e.Tags = make([]tag.Tag, 0)
	}
	tagsData, _ := json.Marshal(e.Tags)
	return []byte(fmt.Sprintf("[0,%s,%d,%d,%s,%s]", e.PubKey, e.CreatedAt, e.Kind, tagsData, e.Content))
}

// Sign signs the event.
func (e *Event) Sign(prvKeyHex string) error {
	prvKeyStr, err := hex.DecodeString(prvKeyHex)
	if err != nil {
		return nil
	}
	prvKey, pubKey := btcec.PrivKeyFromBytes(prvKeyStr)
	e.PubKey = hex.EncodeToString(pubKey.SerializeCompressed()[1:])
	e.CreatedAt = int(time.Now().Unix())
	if e.Tags == nil {
		e.Tags = []tag.Tag{}
	}
	hash := sha256.Sum256(e.Serialize())
	sig, err := schnorr.Sign(prvKey, hash[:])
	if err != nil {
		return err
	}
	e.ID = hex.EncodeToString(hash[:])
	e.Sig = hex.EncodeToString(sig.Serialize())
	return nil
}

// Unmarshal unmarshals the Event from a byte slice.
func (e *Event) Unmarshal(data []byte) error {
	return json.Unmarshal(data, e)
}

// Verify verifies the event pubkey and signature.
func (e *Event) Verify() error {
	pubKeyStr, err := hex.DecodeString(e.PubKey)
	if err != nil {
		return nil
	}
	pubKey, err := schnorr.ParsePubKey(pubKeyStr)
	if err != nil {
		return err
	}
	sigStr, err := hex.DecodeString(e.Sig)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	sig, err := schnorr.ParseSignature(sigStr)
	if err != nil {
		return err
	}
	hsh := sha256.Sum256(e.Serialize())
	if !sig.Verify(hsh[:], pubKey) {
		return fmt.Errorf("unable to verify")
	}
	return nil
}
