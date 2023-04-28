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

type Tag interface {
	Marshal() ([]byte, error)
	Type() *string
	Unmarshal(data []byte) error
}

// RawTag TBD
type RawTag []any

// Marshal TBD
func (t RawTag) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

// Type TBD
func (t RawTag) Type() *string {
	if len(t) < 1 {
		return nil
	}
	if typ, ok := t[0].(string); ok {
		return &typ
	}
	return nil
}

// Unmarshal TBD
func (t RawTag) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &t)
}

func NewAmountTag(amount int) Tag {
	return AmountTag{TagTypeAmount, amount}
}

// AmountTag TBD
type AmountTag = RawTag

func (t AmountTag) Amount() *int {
	if ok := len(t) != 2; !ok {
		return nil
	}
	v, ok := t[1].(float64)
	if !ok {
		return nil
	}
	amount := int(v)
	return &amount
}

func NewBadgeDescTag() Tag {
	return BadgeNameTag{TagTypeBadgeDesc}
}

// BadgeDescTag TBD
type BadgeDescTag = RawTag

func NewBadgeNameTag() Tag {
	return BadgeNameTag{TagTypeBadgeName}
}

// BadgeNameTag TBD
type BadgeNameTag = RawTag

func NewBolt11Tag() Tag {
	return BadgeNameTag{TagTypeBolt11}
}

// Bolt11Tag TBD
type Bolt11Tag = RawTag

func NewChallengeTag() Tag {
	return BadgeNameTag{TagTypeChallenge}
}

// ChallengeTag TBD
type ChallengeTag = RawTag

func NewContentWarnTag() Tag {
	return BadgeNameTag{TagTypeContentWarn}
}

// ContentWarnTag TBD
type ContentWarnTag = RawTag

func NewDelegationTag() Tag {
	return BadgeNameTag{TagTypeDelegation}
}

// DelegationTag TBD
type DelegationTag = RawTag

func NewEventTag() Tag {
	return BadgeNameTag{TagTypeEvent}
}

// EventTag TBD
type EventTag = RawTag

func NewEventIDTag() Tag {
	return BadgeNameTag{TagTypeEventID}
}

// EventIDTag TBD
type EventIDTag = RawTag

func NewExpirationTag() Tag {
	return BadgeNameTag{TagTypeExpiration}
}

// ExpirationTag TBD
type ExpirationTag = RawTag

func NewGeohashTag() Tag {
	return BadgeNameTag{TagTypeGeohash}
}

// GeohashTag TBD
type GeohashTag = RawTag

func NewHashtagTag() Tag {
	return BadgeNameTag{TagTypeHashtag}
}

// HashtagTag TBD
type HashtagTag = RawTag

func NewIdentifierTag() Tag {
	return BadgeNameTag{TagTypeIdentifier}
}

// IdentifierTag TBD
type IdentifierTag = RawTag

func NewIdentityTag() Tag {
	return BadgeNameTag{TagTypeIdentity}
}

// IdentityTag TBD
type IdentityTag = RawTag

func NewImageTag() Tag {
	return BadgeNameTag{TagTypeImage}
}

// ImageTag TBD
type ImageTag = RawTag

func NewInvDescTag() Tag {
	return BadgeNameTag{TagTypeInvDesc}
}

// InvDescTag TBD
type InvDescTag = RawTag

func NewLNURLTag() Tag {
	return BadgeNameTag{TagTypeLNURL}
}

// LNURLTag TBD
type LNURLTag = RawTag

func NewNonceTag() Tag {
	return BadgeNameTag{TagTypeNonce}
}

// NonceTag TBD
type NonceTag = RawTag

func NewPreimageTag() Tag {
	return BadgeNameTag{TagTypePreimage}
}

// PreimageTag TBD
type PreimageTag = RawTag

func NewPubkeyTag() Tag {
	return BadgeNameTag{TagTypePubkey}
}

// PubkeyTag TBD
type PubkeyTag = RawTag

func NewPublishedAtTag() Tag {
	return BadgeNameTag{TagTypePublishedAt}
}

// PublishedAtTag TBD
type PublishedAtTag = RawTag

func NewReferenceTag() Tag {
	return BadgeNameTag{TagTypeReference}
}

// ReferenceTag TBD
type ReferenceTag = RawTag

func NewRelayTag() Tag {
	return BadgeNameTag{TagTypeRelay}
}

// RelayTag TBD
type RelayTag = RawTag

func NewRelaysTag() Tag {
	return BadgeNameTag{TagTypeRelays}
}

// RelaysTag TBD
type RelaysTag = RawTag

func NewSubjectTag() Tag {
	return BadgeNameTag{TagTypeSubject}
}

// SubjectTag TBD
type SubjectTag = RawTag

func NewSummaryTag() Tag {
	return BadgeNameTag{TagTypeSummary}
}

// SummaryTag TBD
type SummaryTag = RawTag

func NewThumbTag() Tag {
	return BadgeNameTag{TagTypeThumb}
}

// ThumbTag TBD
type ThumbTag = RawTag

func NewTitleTag() Tag {
	return BadgeNameTag{TagTypeTitle}
}

// TitleTag TBD
type TitleTag = RawTag

func NewZapTag() Tag {
	return BadgeNameTag{TagTypeZap}
}

// ZapTag TBD
type ZapTag = RawTag
