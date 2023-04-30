package nostr

import (
	"encoding/json"
)

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

type Event interface {
	ID() []byte
	Content() []byte
	Kind() *int
	Marshal() ([]byte, error)
	PublicKey() []byte
	Get(k string) any
	Set(k string, v any) error
	Sign() error
	Signature() []byte
	Tags() []Tag
	Unmarshal(data []byte) error
}

// RawEvent TBD
type RawEvent map[string]json.RawMessage

// Content TBD
func (e RawEvent) Content() []byte {
	var content string
	if err := json.Unmarshal(e["content"], &content); err != nil {
		return nil
	}
	return []byte(content)
}

// Get TBD
func (e RawEvent) Get(k string) any {
	var v any
	if err := json.Unmarshal(e[k], &v); err != nil {
		return nil
	}
	return v
}

// ID TBD
func (e RawEvent) ID() []byte {
	var id string
	if err := json.Unmarshal(e["id"], &id); err != nil {
		return nil
	}
	return []byte(id)
}

// Kind TBD
func (e RawEvent) Kind() *int {
	var kind int
	if err := json.Unmarshal(e["kind"], &kind); err != nil {
		return nil
	}
	return &kind
}

// PublicKey TBD
func (e RawEvent) PublicKey() []byte {
	var publicKey string
	if err := json.Unmarshal(e["pubkey"], &publicKey); err != nil {
		return nil
	}
	return []byte(publicKey)
}

// Set TBD
func (e RawEvent) Set(k string, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	e[k] = data
	return nil
}

// Signature TBD
func (e RawEvent) Signature() []byte {
	var signature []byte
	if err := json.Unmarshal(e["pubkey"], &signature); err != nil {
		return nil
	}
	return signature
}

// Sign TBD
func (e RawEvent) Sign() error {
	return nil
}

// Tags TBD
func (e RawEvent) Tags() []Tag {
	var tags []Tag
	if err := json.Unmarshal(e["tags"], &tags); err != nil {
		return nil
	}
	return tags
}

// Marshal TBD
func (e *RawEvent) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

// Unmarshal TBD
func (e *RawEvent) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &e)
}

// Validate TBD
func (e *RawEvent) Validate() error {
	return nil
}

// NewMetadataEvent TBD
func NewMetadataEvent() Event {
	e := &MetadataEvent{}
	e.Set("kind", EventKindMetadata)
	return e
}

// MetadataEvent TBD
type MetadataEvent = RawEvent

// NewShortTextNoteEvent TBD
func NewShortTextNoteEvent() Event {
	e := &ShortTextNoteEvent{}
	e.Set("kind", EventKindShortTextNote)
	return e
}

// ShortTextNoteEvent TBD
type ShortTextNoteEvent = RawEvent

// NewRecommendRelayEvent TBD
func NewRecommendRelayEvent() Event {
	e := &RecommendRelayEvent{}
	e.Set("kind", EventKindRecommendRelay)
	return e
}

// RecommendRelayEvent TBD
type RecommendRelayEvent = RawEvent

// NewContactsEvent TBD
func NewContactsEvent() Event {
	e := &ContactsEvent{}
	e.Set("kind", EventKindContacts)
	return e
}

// ContactsEvent TBD
type ContactsEvent = RawEvent

// NewEncryptedDirectMessagesEvent TBD
func NewEncryptedDirectMessagesEvent() Event {
	e := &EncryptedDirectMessagesEvent{}
	e.Set("kind", EventKindEncryptedDirectMessages)
	return e
}

// EncryptedDirectMessagesEvent TBD
type EncryptedDirectMessagesEvent = RawEvent

// NewEventDeletionEvent TBD
func NewEventDeletionEvent() Event {
	e := &EventDeletionEvent{}
	e.Set("kind", EventKindEventDeletion)
	return e
}

// EventDeletionEvent TBD
type EventDeletionEvent = RawEvent

// NewEventRepostsEvent TBD
func NewEventRepostsEvent() Event {
	e := &EventRepostsEvent{}
	e.Set("kind", EventKindReposts)
	return e
}

// EventRepostsEvent TBD
type EventRepostsEvent = RawEvent

// NewReactionEvent TBD
func NewReactionEvent() Event {
	e := &ReactionEvent{}
	e.Set("kind", EventKindReaction)
	return e
}

// ReactionEvent TBD
type ReactionEvent = RawEvent

// NewBadgeAwardEvent TBD
func NewBadgeAwardEvent() Event {
	e := &BadgeAwardEvent{}
	e.Set("kind", EventKindBadgeAward)
	return e
}

// BadgeAwardEvent TBD
type BadgeAwardEvent = RawEvent

// NewChannelCreationEvent TBD
func NewChannelCreationEvent() Event {
	e := &ChannelCreationEvent{}
	e.Set("kind", EventKindChannelCreation)
	return e
}

// ChannelCreationEvent TBD
type ChannelCreationEvent = RawEvent

func NewChannelMetadataEvent() Event {
	e := &ChannelMetadataEvent{}
	e.Set("kind", EventKindChannelMetadata)
	return e
}

// ChannelMetadata TBD
type ChannelMetadataEvent = RawEvent

func NewChannelMessageEvent() Event {
	e := &ChannelMessageEvent{}
	e.Set("kind", EventKindChannelMessage)
	return e
}

// ChannelMessage TBD
type ChannelMessageEvent = RawEvent

func NewChannelHideMessageEvent() Event {
	e := &ChannelHideMessageEvent{}
	e.Set("kind", EventKindChannelHideMessage)
	return e
}

// ChannelHideMessage TBD
type ChannelHideMessageEvent = RawEvent

func NewChannelMuteUserEvent() Event {
	e := &ChannelMuteUserEvent{}
	e.Set("kind", EventKindChannelMuteUser)
	return e
}

// ChannelMuteUser TBD
type ChannelMuteUserEvent = RawEvent

func NewReportingEvent() Event {
	e := &ReportingEvent{}
	e.Set("kind", EventKindReporting)
	return e
}

// Reporting TBD
type ReportingEvent = RawEvent

func NewZapRequestEvent() Event {
	e := &ZapRequestEvent{}
	e.Set("kind", EventKindZapRequest)
	return e
}

// ZapRequest TBD
type ZapRequestEvent = RawEvent

func NewZapEvent() Event {
	e := &ZapEvent{}
	e.Set("kind", EventKindZap)
	return e
}

// Zap TBD
type ZapEvent = RawEvent

func NewMuteListEvent() Event {
	e := &MuteListEvent{}
	e.Set("kind", EventKindMuteList)
	return e
}

// MuteList TBD
type MuteListEvent = RawEvent

func NewPinListEvent() Event {
	e := &PinListEvent{}
	e.Set("kind", EventKindPinList)
	return e
}

// PinList TBD
type PinListEvent = RawEvent

func NewRelayListMetadataEvent() Event {
	e := &RelayListMetadataEvent{}
	e.Set("kind", EventKindRelayListMetadata)
	return e
}

// RelayListMetadata TBD
type RelayListMetadataEvent = RawEvent

func NewClientAuthenticationEvent() Event {
	e := &ClientAuthenticationEvent{}
	e.Set("kind", EventKindClientAuthentication)
	return e
}

// ClientAuthentication TBD
type ClientAuthenticationEvent = RawEvent

func NewNostrConnectEvent() Event {
	e := &NostrConnectEvent{}
	e.Set("kind", EventKindNostrConnect)
	return e
}

// NostrConnect TBD
type NostrConnectEvent = RawEvent

func NewCategorizedPeopleListEvent() Event {
	e := &CategorizedPeopleListEvent{}
	e.Set("kind", EventKindCategorizedPeopleList)
	return e
}

// CategorizedPeopleList TBD
type CategorizedPeopleListEvent = RawEvent

func NewCategorizedBookmarkListEvent() Event {
	e := &CategorizedBookmarkListEvent{}
	e.Set("kind", EventKindCategorizedBookmarkList)
	return e
}

// CategorizedBookmarkList TBD
type CategorizedBookmarkListEvent = RawEvent

func NewProfileBadgesEvent() Event {
	e := &ProfileBadgesEvent{}
	e.Set("kind", EventKindProfileBadges)
	return e
}

// ProfileBadges TBD
type ProfileBadgesEvent = RawEvent

func NewBadgeDefinitionEvent() Event {
	e := &BadgeDefinitionEvent{}
	e.Set("kind", EventKindBadgeDefinition)
	return e
}

// BadgeDefinition TBD
type BadgeDefinitionEvent = RawEvent

func NewCreateOrUpdateStallEvent() Event {
	e := &CreateOrUpdateStallEvent{}
	e.Set("kind", EventKindCreateOrUpdateStall)
	return e
}

// CreateOrUpdateStall TBD
type CreateOrUpdateStallEvent = RawEvent

func NewCreateOrUpdateProductEvent() Event {
	e := &CreateOrUpdateProductEvent{}
	e.Set("kind", EventKindCreateOrUpdateProduct)
	return e
}

// CreateOrUpdateProduct TBD
type CreateOrUpdateProductEvent = RawEvent

func NewLongFormContentEvent() Event {
	e := &LongFormContentEvent{}
	e.Set("kind", EventKindLongFormContent)
	return e
}

// LongFormContent TBD
type LongFormContentEvent = RawEvent

func NewApplicationSpecificDataEvent() Event {
	e := &ApplicationSpecificDataEvent{}
	e.Set("kind", EventKindApplicationSpecificData)
	return e
}

// ApplicationSpecificData TBD
type ApplicationSpecificDataEvent = RawEvent
