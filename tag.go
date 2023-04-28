package nostr

import (
	"encoding/json"
)

// TagType TBD
type TagType string

const (
	TagTypeAmount      TagType = "amount"          // TagTypeAmount represents millisats.
	TagTypeBadgeDesc   TagType = "description"     // TagTypeBadgeDesc represents a badge description.
	TagTypeBadgeName   TagType = "name"            // TagTypeBadgeName represents a badge name.
	TagTypeBolt11      TagType = "bolt11"          // TagTypeBolt11 represents a Bolt11 invoice.
	TagTypeChallenge   TagType = "challenge"       // TagTypeChallenge represents a challenge string.
	TagTypeContentWarn TagType = "content-warning" // TagTypeContentWarn represents a content warning reason.
	TagTypeDelegation  TagType = "delegation"      // TagTypeDelegation represents a delegation with pubkey, conditions, and delegation token.
	TagTypeEvent       TagType = "a"               // TagTypeEvent represents coordinates to an event relay URL.
	TagTypeEventID     TagType = "e"               // TagTypeEventID represents an event ID (hex) used in relay URLs and markers.
	TagTypeExpiration  TagType = "expiration"      // TagTypeExpiration represents a Unix timestamp (string) for expiration.
	TagTypeGeohash     TagType = "g"               // TagTypeGeohash represents a geohash.
	TagTypeHashtag     TagType = "t"               // TagTypeHashtag represents a hashtag.
	TagTypeIdentifier  TagType = "d"               // TagTypeIdentifier represents an identifier.
	TagTypeIdentity    TagType = "i"               // TagTypeIdentity represents an identity used in proofs.
	TagTypeImage       TagType = "image"           // TagTypeImage represents an image URL with dimensions in pixels.
	TagTypeInvDesc     TagType = "description"     // TagTypeInvDesc represents an invoice description.
	TagTypeLNURL       TagType = "lnurl"           // TagTypeLNURL represents a Bech32 encoded LNURL.
	TagTypeNonce       TagType = "nonce"           // TagTypeNonce represents a random nonce.
	TagTypePreimage    TagType = "preimage"        // TagTypePreimage represents a hash of a Bolt11 invoice.
	TagTypePubkey      TagType = "p"               // TagTypePubkey represents a public key (hex) used in relay URLs.
	TagTypePublishedAt TagType = "published_at"    // TagTypePublishedAt represents a Unix timestamp (string) for when an item was published.
	TagTypeReference   TagType = "r"               // TagTypeReference represents a reference (URL, etc).
	TagTypeRelay       TagType = "relay"           // TagTypeRelay represents a relay URL.
	TagTypeRelays      TagType = "relays"          // TagTypeRelays represents a list of relays.
	TagTypeSubject     TagType = "subject"         // TagTypeSubject represents a subject.
	TagTypeSummary     TagType = "summary"         // TagTypeSummary represents an article summary.
	TagTypeThumb       TagType = "thumb"           // TagTypeThumb represents a badge thumbnail with dimensions in pixels.
	TagTypeTitle       TagType = "title"           // TagTypeTitle represents an article title.
	TagTypeZap         TagType = "zap"             // TagTypeZap represents a profile name with a type of value.
)

// Tag TBD
type Tag interface {
	Marshal() ([]byte, error)
	Type() TagType
	Unmarshal(data []byte) error
}

func NewBaseTag(typ TagType, args ...interface{}) Tag {
	tag := BaseTag([]interface{}{typ})
	for _, arg := range args {
		tag = append(tag, arg)
	}
	return tag
}

// BaseTag TBD
type BaseTag []interface{}

// Marshal TBD
func (t BaseTag) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

// Type TBD
func (t BaseTag) Type() TagType {
	if len(t) < 1 {
		panic("not ok")
	}
	return t[0].(TagType)
}

// Unmarshal TBD
func (t BaseTag) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &t)
}

// Validate TBD
func (t BaseTag) Validate() error {
	return nil
}

func NewAmountTag(amount float64) Tag {
	return &AmountTag{BaseTag{TagTypeAmount, amount}}
}

// AmountTag TBD
type AmountTag struct {
	BaseTag
}

// Amount TBD
func (t *AmountTag) Amount() float64 {
	amount, ok := t.BaseTag[1].(float64)
	if !ok {
		panic("not ok")
	}
	return amount
}

// Marshal TBD
func (t *AmountTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *AmountTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

func NewBadgeDescTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeBadgeDesc}}
}

// BadgeDescTag TBD
type BadgeDescTag struct {
	BaseTag
}

// Marshal TBD
func (t *BadgeDescTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *BadgeDescTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *BadgeDescTag) Validate() error {
	return nil
}

func NewBadgeNameTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeBadgeName}}
}

// BadgeNameTag TBD
type BadgeNameTag struct {
	BaseTag
}

// Marshal TBD
func (t *BadgeNameTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *BadgeNameTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *BadgeNameTag) Validate() error {
	return nil
}

func NewBolt11Tag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeBolt11}}
}

// Bolt11Tag TBD
type Bolt11Tag struct {
	BaseTag
}

// Marshal TBD
func (t *Bolt11Tag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *Bolt11Tag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *Bolt11Tag) Validate() error {
	return nil
}

func NewChallengeTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeChallenge}}
}

// ChallengeTag TBD
type ChallengeTag struct {
	BaseTag
}

// Marshal TBD
func (t *ChallengeTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *ChallengeTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *ChallengeTag) Validate() error {
	return nil
}

func NewContentWarnTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeContentWarn}}
}

// ContentWarnTag TBD
type ContentWarnTag struct {
	BaseTag
}

// Marshal TBD
func (t *ContentWarnTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *ContentWarnTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *ContentWarnTag) Validate() error {
	return nil
}

func NewDelegationTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeDelegation}}
}

// DelegationTag TBD
type DelegationTag struct {
	BaseTag
}

// Marshal TBD
func (t *DelegationTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *DelegationTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *DelegationTag) Validate() error {
	return nil
}

func NewEventTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeEvent}}
}

// EventTag TBD
type EventTag struct {
	BaseTag
}

// Marshal TBD
func (t *EventTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *EventTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *EventTag) Validate() error {
	return nil
}

func NewEventIDTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeEventID}}
}

// EventIDTag TBD
type EventIDTag struct {
	BaseTag
}

// Marshal TBD
func (t *EventIDTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *EventIDTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *EventIDTag) Validate() error {
	return nil
}

func NewExpirationTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeExpiration}}
}

// ExpirationTag TBD
type ExpirationTag struct {
	BaseTag
}

// Marshal TBD
func (t *ExpirationTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *ExpirationTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *ExpirationTag) Validate() error {
	return nil
}

func NewGeohashTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeGeohash}}
}

// GeohashTag TBD
type GeohashTag struct {
	BaseTag
}

// Marshal TBD
func (t *GeohashTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *GeohashTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *GeohashTag) Validate() error {
	return nil
}

func NewHashtagTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeHashtag}}
}

// HashtagTag TBD
type HashtagTag struct {
	BaseTag
}

// Marshal TBD
func (t *HashtagTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *HashtagTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *HashtagTag) Validate() error {
	return nil
}

func NewIdentifierTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeIdentifier}}
}

// IdentifierTag TBD
type IdentifierTag struct {
	BaseTag
}

// Marshal TBD
func (t *IdentifierTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *IdentifierTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *IdentifierTag) Validate() error {
	return nil
}

func NewIdentityTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeIdentity}}
}

// IdentityTag TBD
type IdentityTag struct {
	BaseTag
}

// Marshal TBD
func (t *IdentityTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *IdentityTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *IdentityTag) Validate() error {
	return nil
}

func NewImageTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeImage}}
}

// ImageTag TBD
type ImageTag struct {
	BaseTag
}

// Marshal TBD
func (t *ImageTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *ImageTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *ImageTag) Validate() error {
	return nil
}

func NewInvDescTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeInvDesc}}
}

// InvDescTag TBD
type InvDescTag struct {
	BaseTag
}

// Marshal TBD
func (t *InvDescTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *InvDescTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *InvDescTag) Validate() error {
	return nil
}

func NewLNURLTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeLNURL}}
}

// LNURLTag TBD
type LNURLTag struct {
	BaseTag
}

// Marshal TBD
func (t *LNURLTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *LNURLTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *LNURLTag) Validate() error {
	return nil
}

func NewNonceTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeNonce}}
}

// NonceTag TBD
type NonceTag struct {
	BaseTag
}

// Marshal TBD
func (t *NonceTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *NonceTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *NonceTag) Validate() error {
	return nil
}

func NewPreimageTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypePreimage}}
}

// PreimageTag TBD
type PreimageTag struct {
	BaseTag
}

// Marshal TBD
func (t *PreimageTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *PreimageTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *PreimageTag) Validate() error {
	return nil
}

func NewPubkeyTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypePubkey}}
}

// PubkeyTag TBD
type PubkeyTag struct {
	BaseTag
}

// Marshal TBD
func (t *PubkeyTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *PubkeyTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *PubkeyTag) Validate() error {
	return nil
}

func NewPublishedAtTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypePublishedAt}}
}

// PublishedAtTag TBD
type PublishedAtTag struct {
	BaseTag
}

// Marshal TBD
func (t *PublishedAtTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *PublishedAtTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *PublishedAtTag) Validate() error {
	return nil
}

func NewReferenceTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeReference}}
}

// ReferenceTag TBD
type ReferenceTag struct {
	BaseTag
}

// Marshal TBD
func (t *ReferenceTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *ReferenceTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *ReferenceTag) Validate() error {
	return nil
}

func NewRelayTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeRelay}}
}

// RelayTag TBD
type RelayTag struct {
	BaseTag
}

// Marshal TBD
func (t *RelayTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *RelayTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *RelayTag) Validate() error {
	return nil
}

func NewRelaysTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeRelays}}
}

// RelaysTag TBD
type RelaysTag struct {
	BaseTag
}

// Marshal TBD
func (t *RelaysTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *RelaysTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *RelaysTag) Validate() error {
	return nil
}

func NewSubjectTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeSubject}}
}

// SubjectTag TBD
type SubjectTag struct {
	BaseTag
}

// Marshal TBD
func (t *SubjectTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *SubjectTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *SubjectTag) Validate() error {
	return nil
}

func NewSummaryTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeSummary}}
}

// SummaryTag TBD
type SummaryTag struct {
	BaseTag
}

// Marshal TBD
func (t *SummaryTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *SummaryTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *SummaryTag) Validate() error {
	return nil
}

func NewThumbTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeThumb}}
}

// ThumbTag TBD
type ThumbTag struct {
	BaseTag
}

// Marshal TBD
func (t *ThumbTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *ThumbTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *ThumbTag) Validate() error {
	return nil
}

func NewTitleTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeTitle}}
}

// TitleTag TBD
type TitleTag struct {
	BaseTag
}

// Marshal TBD
func (t *TitleTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *TitleTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *TitleTag) Validate() error {
	return nil
}

func NewZapTag() Tag {
	return &BadgeNameTag{BaseTag{TagTypeZap}}
}

// ZapTag TBD
type ZapTag struct {
	BaseTag
}

// Marshal TBD
func (t *ZapTag) Marshal() ([]byte, error) {
	return t.BaseTag.Marshal()
}

// Unmarshal TBD
func (t *ZapTag) Unmarshal(data []byte) error {
	return t.BaseTag.Unmarshal(data)
}

// Validate TBD
func (t *ZapTag) Validate() error {
	return nil
}
