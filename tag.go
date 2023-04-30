package nostr

import (
	"encoding/json"
)

const (
	TagTypeAmount      = "amount"          // TagTypeAmount represents millisats.
	TagTypeBadgeDesc   = "description"     // TagTypeBadgeDesc represents a badge description.
	TagTypeBadgeName   = "name"            // TagTypeBadgeName represents a badge name.
	TagTypeBolt11      = "bolt11"          // TagTypeBolt11 represents a Bolt11 invoice.
	TagTypeChallenge   = "challenge"       // TagTypeChallenge represents a challenge string.
	TagTypeContentWarn = "content-warning" // TagTypeContentWarn represents a content warning reason.
	TagTypeDelegation  = "delegation"      // TagTypeDelegation represents a delegation with pubkey, conditions, and delegation token.
	TagTypeEvent       = "a"               // TagTypeEvent represents coordinates to an event relay URL.
	TagTypeEventID     = "e"               // TagTypeEventID represents an event ID (hex) used in relay URLs and markers.
	TagTypeExpiration  = "expiration"      // TagTypeExpiration represents a Unix timestamp (string) for expiration.
	TagTypeGeohash     = "g"               // TagTypeGeohash represents a geohash.
	TagTypeHashtag     = "t"               // TagTypeHashtag represents a hashtag.
	TagTypeIdentifier  = "d"               // TagTypeIdentifier represents an identifier.
	TagTypeIdentity    = "i"               // TagTypeIdentity represents an identity used in proofs.
	TagTypeImage       = "image"           // TagTypeImage represents an image URL with dimensions in pixels.
	TagTypeInvDesc     = "description"     // TagTypeInvDesc represents an invoice description.
	TagTypeLNURL       = "lnurl"           // TagTypeLNURL represents a Bech32 encoded LNURL.
	TagTypeNonce       = "nonce"           // TagTypeNonce represents a random nonce.
	TagTypePreimage    = "preimage"        // TagTypePreimage represents a hash of a Bolt11 invoice.
	TagTypePubkey      = "p"               // TagTypePubkey represents a public key (hex) used in relay URLs.
	TagTypePublishedAt = "published_at"    // TagTypePublishedAt represents a Unix timestamp (string) for when an item was published.
	TagTypeReference   = "r"               // TagTypeReference represents a reference (URL, etc).
	TagTypeRelay       = "relay"           // TagTypeRelay represents a relay URL.
	TagTypeRelays      = "relays"          // TagTypeRelays represents a list of relays.
	TagTypeSubject     = "subject"         // TagTypeSubject represents a subject.
	TagTypeSummary     = "summary"         // TagTypeSummary represents an article summary.
	TagTypeThumb       = "thumb"           // TagTypeThumb represents a badge thumbnail with dimensions in pixels.
	TagTypeTitle       = "title"           // TagTypeTitle represents an article title.
	TagTypeZap         = "zap"             // TagTypeZap represents a profile name with a type of value.
)

// Tag is an interface for encoding and marshaling messages.
type Tag interface {
	Marshal() ([]byte, error)
	Push(v any) error
	Type() string
	Unmarshal(data []byte) error
	Values() []any
}

// NewRawTag
func NewRawTag(typ string) Tag {
	tag := &RawTag{}
	tag.Push(typ)
	return tag
}

// RawTag TBD
type RawTag []json.RawMessage

// Marshal TBD
func (t *RawTag) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

// Push TBD
func (t *RawTag) Push(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	*t = append(*t, data)
	return nil
}

// Type TBD
func (t *RawTag) Type() string {
	if len(*t) < 1 {
		return ""
	}
	typ, ok := t.Values()[0].(string)
	if !ok {
		return ""
	}
	return typ
}

// Unmarshal TBD
func (t *RawTag) Unmarshal(data []byte) error {
	return json.Unmarshal(data, t)
}

// Values TBD
func (t *RawTag) Values() []any {
	var vals []any
	for _, arg := range *t {
		var val any
		json.Unmarshal(arg, &val)
		vals = append(vals, val)
	}
	return vals
}
