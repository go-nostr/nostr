package nostr

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
)

// EventKind constants represent the different kinds of events that can be
// handled by the nostr system.
const (
	EventKindMetadata                = 0     // Event for setting metadata
	EventKindShortTextNote           = 1     // Event for storing a text note
	EventKindRecommendRelay          = 2     // Event for recommending a server
	EventKindContacts                = 3     // Event for contact List and petnames
	EventKindEncryptedDirectMessages = 4     // Event for encrypted direct messages
	EventKindEventDeletion           = 5     // Event for deleting events
	EventKindReposts                 = 6     // Event for signaling to followers that another event is worth reading
	EventKindReaction                = 7     // Event for reacting to other notes
	EventKindBadgeAward              = 8     // Event for awarding badges to users
	EventKindChannelCreation         = 40    // Event for creating new channels
	EventKindChannelMetadata         = 41    // Event for setting channel metadata
	EventKindChannelMessage          = 42    // Event for posting messages in a channel
	EventKindChannelHideMessage      = 43    // Event for hiding messages in a channel
	EventKindChannelMuteUser         = 44    // Event for muting a user in a channel
	EventKindReporting               = 1984  // Event for reporting content or users
	EventKindZapRequest              = 9734  // Event for requesting a Zap action
	EventKindZap                     = 9735  // Event for performing a Zap action
	EventKindMuteList                = 10000 // Event for managing a mute list
	EventKindPinList                 = 10001 // Event for managing a pin list
	EventKindRelayListMetadata       = 10002 // Event for managing relay list metadata
	EventKindClientAuthentication    = 22242 // Event for client authentication process
	EventKindNostrConnect            = 24133 // Event for Nostr client connection
	EventKindCategorizedPeopleList   = 30000 // Event for managing categorized people list
	EventKindCategorizedBookmarkList = 30001 // Event for managing categorized bookmark list
	EventKindProfileBadges           = 30008 // Event for profile badges management
	EventKindBadgeDefinition         = 30009 // Event for defining badges
	EventKindCreateOrUpdateStall     = 30017 // Event for creating or updating a stall
	EventKindCreateOrUpdateProduct   = 30018 // Event for creating or updating a product
	EventKindLongFormContent         = 30023 // Event for posting long-form content
	EventKindApplicationSpecificData = 30078 // Event for managing application-specific data
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
	Tags() []Tag
	Unmarshal(data []byte) error
	Validate() error
	Values() []any
}

// NewRawEvent creates a new empty raw event.
func NewRawEvent() Event {
	return &RawEvent{}
}

// RawEvent is a map-based representation of an event with keys and values
// encoded as json.RawMessage.
type RawEvent map[string]json.RawMessage

// Content retrieves the content of the RawEvent as a byte slice.
func (e *RawEvent) Content() []byte {
	var content string
	if err := json.Unmarshal((*e)["content"], &content); err != nil {
		return []byte{}
	}
	return []byte(content)
}

// CreatedAt TBD
func (e *RawEvent) CreatedAt() int {
	var createdAt int
	if err := json.Unmarshal((*e)["created_at"], &createdAt); err != nil {
		return -1
	}
	return createdAt
}

// Get retrieves the value associated with the given key from the RawEvent.
func (e *RawEvent) Get(key string) any {
	var val any
	if err := json.Unmarshal((*e)[key], &val); err != nil {
		return nil
	}
	return val
}

// ID retrieves the ID of the RawEvent as a byte slice.
func (e *RawEvent) ID() []byte {
	h := sha256.Sum256(e.Serialize())
	return []byte(hex.EncodeToString(h[:]))
}

// Keys returns a list of keys in the RawEvent.
func (e *RawEvent) Keys() []string {
	var keys []string
	for key := range *e {
		keys = append(keys, key)
	}
	return keys
}

// Kind retrieves the kind of the RawEvent as an integer.
func (e *RawEvent) Kind() int {
	var kind int
	if err := json.Unmarshal((*e)["kind"], &kind); err != nil {
		return -1
	}
	return kind
}

// Marshal marshals the RawEvent to a byte slice.
func (e *RawEvent) Marshal() ([]byte, error) {
	return e.MarshalJSON()
}

// MarshalJSON marshals the RawEvent to a byte slice.
func (e *RawEvent) MarshalJSON() ([]byte, error) {
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

// PublicKey retrieves the public key of the RawEvent as a byte slice.
func (e *RawEvent) PublicKey() []byte {
	var pubKey string
	if err := json.Unmarshal((*e)["pubkey"], &pubKey); err != nil {
		return []byte{}
	}
	return []byte(pubKey)
}

// Serialize TBD
func (e *RawEvent) Serialize() []byte {
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

// Set associates a value with a given key in the RawEvent.
func (e *RawEvent) Set(key string, val any) error {
	data, err := json.Marshal(val)
	if err != nil {
		return err
	}
	(*e)[key] = data
	return nil
}

// Signature retrieves the signature of the RawEvent as a byte slice.
func (e *RawEvent) Signature() []byte {
	var sig string
	if err := json.Unmarshal((*e)["sig"], &sig); err != nil {
		return []byte{}
	}
	return []byte(sig)
}

// Sign signs the RawEvent.
func (e *RawEvent) Sign(privateKey string) error {
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
		e.Set("tags", []RawTag{})
	}
	e.Set("id", hex.EncodeToString(h[:]))
	e.Set("sig", hex.EncodeToString(sig.Serialize()))
	return nil
}

// Tags retrieves the list of tags associated with the RawEvent.
func (e *RawEvent) Tags() []Tag {
	tags := make([]Tag, 0)
	var args []json.RawMessage
	if err := json.Unmarshal((*e)["tags"], &args); err != nil {
		return tags
	}
	for _, arg := range args {
		var tag RawTag
		if err := tag.Unmarshal(arg); err != nil {
			return tags
		}
		tags = append(tags, &tag)
	}
	return tags
}

// Unmarshal unmarshals the RawEvent from a byte slice.
func (e *RawEvent) Unmarshal(data []byte) error {
	return e.UnmarshalJSON(data)
}

// UnmarshalJSON TBD
func (e *RawEvent) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, e); err != nil {
		return err
	}
	return nil
}

// Validate validates the RawEvent.
func (e *RawEvent) Validate() error {
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

// Values returns a list of values in the RawEvent.
func (e *RawEvent) Values() []any {
	var vals []any
	for _, v := range *e {
		vals = append(vals, v)
	}
	return vals
}

// NewMetadataEvent creates a new metadata event.
func NewMetadataEvent() Event {
	event := &MetadataEvent{&RawEvent{}}
	event.Set("kind", EventKindMetadata)
	return event
}

// MetadataEvent represents a metadata event.
type MetadataEvent struct {
	*RawEvent
}

// NewShortTextNoteEvent creates a new short text note event.
func NewShortTextNoteEvent(content string) Event {
	event := &ShortTextNoteEvent{&RawEvent{}}
	event.Set("kind", EventKindShortTextNote)
	event.Set("content", content)
	return event
}

// ShortTextNoteEvent represents a short text note event.
type ShortTextNoteEvent struct {
	*RawEvent
}

// NewRecommendRelayEvent creates a new recommend relay event.
func NewRecommendRelayEvent() Event {
	event := &RecommendRelayEvent{}
	event.Set("kind", EventKindRecommendRelay)
	return event
}

// RecommendRelayEvent represents a recommend relay event.
type RecommendRelayEvent struct {
	*RawEvent
}

// NewContactsEvent creates a new contacts event.
func NewContactsEvent() Event {
	event := &ContactsEvent{}
	event.Set("kind", EventKindContacts)
	return event
}

// ContactsEvent represents a contacts event.
type ContactsEvent struct {
	*RawEvent
}

// NewEncryptedDirectMessagesEvent creates a new encrypted direct messages event.
func NewEncryptedDirectMessagesEvent() Event {
	event := &EncryptedDirectMessagesEvent{}
	event.Set("kind", EventKindEncryptedDirectMessages)
	return event
}

// EncryptedDirectMessagesEvent represents an encrypted direct messages event.
type EncryptedDirectMessagesEvent struct {
	*RawEvent
}

// NewEventDeletionEvent creates a new event deletion event.
func NewEventDeletionEvent() Event {
	event := &EventDeletionEvent{}
	event.Set("kind", EventKindEventDeletion)
	return event
}

// EventDeletionEvent represents an event deletion event.
type EventDeletionEvent struct {
	*RawEvent
}

// NewEventRepostsEvent creates a new event reposts event.
func NewEventRepostsEvent() Event {
	event := &EventRepostsEvent{}
	event.Set("kind", EventKindReposts)
	return event
}

// EventRepostsEvent represents an event reposts event.
type EventRepostsEvent struct {
	*RawEvent
}

// NewReactionEvent creates a new reaction event.
func NewReactionEvent() Event {
	event := &ReactionEvent{}
	event.Set("kind", EventKindReaction)
	return event
}

// ReactionEvent represents a reaction event.
type ReactionEvent struct {
	*RawEvent
}

// NewBadgeAwardEvent creates a new badge award event.
func NewBadgeAwardEvent() Event {
	event := &BadgeAwardEvent{}
	event.Set("kind", EventKindBadgeAward)
	return event
}

// BadgeAwardEvent represents a badge award event.
type BadgeAwardEvent struct {
	*RawEvent
}

// NewChannelCreationEvent creates a new channel creation event.
func NewChannelCreationEvent() Event {
	event := &ChannelCreationEvent{}
	event.Set("kind", EventKindChannelCreation)
	return event
}

// ChannelCreationEvent represents a channel creation event.
type ChannelCreationEvent struct {
	*RawEvent
}

// NewChannelMetadataEvent creates a new channel metadata event.
func NewChannelMetadataEvent() Event {
	event := &ChannelMetadataEvent{}
	event.Set("kind", EventKindChannelMetadata)
	return event
}

// ChannelMetadata represents a channel metadata event.
type ChannelMetadataEvent struct {
	*RawEvent
}

// NewChannelMessageEvent creates a new channel message event.
func NewChannelMessageEvent() Event {
	event := &ChannelMessageEvent{}
	event.Set("kind", EventKindChannelMessage)
	return event
}

// ChannelMessage TODO
type ChannelMessageEvent struct {
	*RawEvent
}

// NewChannelHideMessageEvent TODO
func NewChannelHideMessageEvent() Event {
	event := &ChannelHideMessageEvent{}
	event.Set("kind", EventKindChannelHideMessage)
	return event
}

// ChannelHideMessage represents a message hiding event in a channel.
type ChannelHideMessageEvent struct {
	*RawEvent
}

// NewChannelMuteUserEvent creates a new ChannelMuteUserEvent.
func NewChannelMuteUserEvent() Event {
	event := &ChannelMuteUserEvent{}
	event.Set("kind", EventKindChannelMuteUser)
	return event
}

// ChannelMuteUser represents a user mute event in a channel.
type ChannelMuteUserEvent struct {
	*RawEvent
}

// NewReportingEvent creates a new ReportingEvent.
func NewReportingEvent() Event {
	event := &ReportingEvent{}
	event.Set("kind", EventKindReporting)
	return event
}

// Reporting represents a reporting event.
type ReportingEvent struct {
	*RawEvent
}

// NewZapRequestEvent creates a new ZapRequestEvent.
func NewZapRequestEvent() Event {
	event := &ZapRequestEvent{}
	event.Set("kind", EventKindZapRequest)
	return event
}

// ZapRequest represents a zap request event.
type ZapRequestEvent struct {
	*RawEvent
}

// NewZapEvent creates a new ZapEvent.
func NewZapEvent() Event {
	event := &ZapEvent{}
	event.Set("kind", EventKindZap)
	return event
}

// Zap represents a zap event.
type ZapEvent struct {
	*RawEvent
}

// NewMuteListEvent creates a new MuteListEvent.
func NewMuteListEvent() Event {
	event := &MuteListEvent{}
	event.Set("kind", EventKindMuteList)
	return event
}

// MuteList represents a mute list event.
type MuteListEvent struct {
	*RawEvent
}

// NewPinListEvent creates a new PinListEvent.
func NewPinListEvent() Event {
	event := &PinListEvent{}
	event.Set("kind", EventKindPinList)
	return event
}

// PinList represents a pin list event.
type PinListEvent struct {
	*RawEvent
}

// NewRelayListMetadataEvent creates a new RelayListMetadataEvent.
func NewRelayListMetadataEvent() Event {
	event := &RelayListMetadataEvent{}
	event.Set("kind", EventKindRelayListMetadata)
	return event
}

// RelayListMetadata represents a relay list metadata event.
type RelayListMetadataEvent struct {
	*RawEvent
}

// NewClientAuthenticationEvent creates a new ClientAuthenticationEvent.
func NewClientAuthenticationEvent() Event {
	event := &ClientAuthenticationEvent{}
	event.Set("kind", EventKindClientAuthentication)
	return event
}

// ClientAuthentication represents a client authentication event.
type ClientAuthenticationEvent struct {
	*RawEvent
}

// NewNostrConnectEvent creates a new NostrConnectEvent.
func NewNostrConnectEvent() Event {
	event := &NostrConnectEvent{}
	event.Set("kind", EventKindNostrConnect)
	return event
}

// NostrConnect represents a nostr connect event.
type NostrConnectEvent struct {
	*RawEvent
}

// NewCategorizedPeopleListEvent creates a new CategorizedPeopleListEvent.
func NewCategorizedPeopleListEvent() Event {
	event := &CategorizedPeopleListEvent{}
	event.Set("kind", EventKindCategorizedPeopleList)
	return event
}

// CategorizedPeopleList represents a categorized people list event.
type CategorizedPeopleListEvent struct {
	*RawEvent
}

// NewCategorizedBookmarkListEvent creates a new CategorizedBookmarkListEvent.
func NewCategorizedBookmarkListEvent() Event {
	event := &CategorizedBookmarkListEvent{}
	event.Set("kind", EventKindCategorizedBookmarkList)
	return event
}

// CategorizedBookmarkList represents a categorized bookmark list event.
type CategorizedBookmarkListEvent struct {
	*RawEvent
}

// NewProfileBadgesEvent creates a new ProfileBadgesEvent.
func NewProfileBadgesEvent() Event {
	event := &ProfileBadgesEvent{}
	event.Set("kind", EventKindProfileBadges)
	return event
}

// ProfileBadges represents a profile badges event.
type ProfileBadgesEvent struct {
	*RawEvent
}

// NewBadgeDefinitionEvent creates a new BadgeDefinitionEvent.
func NewBadgeDefinitionEvent() Event {
	event := &BadgeDefinitionEvent{}
	event.Set("kind", EventKindBadgeDefinition)
	return event
}

// BadgeDefinition represents a badge definition event.
type BadgeDefinitionEvent struct {
	*RawEvent
}

// NewCreateOrUpdateStallEvent creates a new CreateOrUpdateStallEvent.
func NewCreateOrUpdateStallEvent() Event {
	event := &CreateOrUpdateStallEvent{}
	event.Set("kind", EventKindCreateOrUpdateStall)
	return event
}

// CreateOrUpdateStall represents a create or update stall event.
type CreateOrUpdateStallEvent struct {
	*RawEvent
}

// NewCreateOrUpdateProductEvent creates a new CreateOrUpdateProductEvent.
func NewCreateOrUpdateProductEvent() Event {
	event := &CreateOrUpdateProductEvent{}
	event.Set("kind", EventKindCreateOrUpdateProduct)
	return event
}

// CreateOrUpdateProduct represents a create or update product event.
type CreateOrUpdateProductEvent struct {
	*RawEvent
}

// NewLongFormContentEvent creates a new LongFormContentEvent.
func NewLongFormContentEvent() Event {
	event := &LongFormContentEvent{}
	event.Set("kind", EventKindLongFormContent)
	return event
}

// LongFormContent represents a long form content event.
type LongFormContentEvent struct {
	*RawEvent
}

// NewApplicationSpecificDataEvent creates a new ApplicationSpecificDataEvent.
func NewApplicationSpecificDataEvent() Event {
	event := &ApplicationSpecificDataEvent{}
	event.Set("kind", EventKindApplicationSpecificData)
	return event
}

// ApplicationSpecificData represents an application-specific data event.
type ApplicationSpecificDataEvent struct {
	*RawEvent
}
