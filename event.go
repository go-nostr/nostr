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
	Sign() error
	Signature() []byte
	Tags() []Tag
	Unmarshal(data []byte) error
}

// RawEvent TBD
type RawEvent map[string]any

// Content TBD
func (e RawEvent) Content() []byte {
	if content, ok := e["content"].(string); ok {
		return []byte(content)
	}
	return nil
}

// ID TBD
func (e RawEvent) ID() []byte {
	if id, ok := e["id"].(string); ok {
		return []byte(id)
	}
	return nil
}

// Kind TBD
func (e RawEvent) Kind() *int {
	if v, ok := e["kind"].(float64); ok {
		kind := int(v)
		return &kind
	}
	return nil
}

// PublicKey TBD
func (e RawEvent) PublicKey() []byte {
	if pubKey, ok := e["pubkey"].(string); ok {
		return []byte(pubKey)
	}
	return nil
}

// Signature TBD
func (e RawEvent) Signature() []byte {
	if sig, ok := e["sig"].(string); ok {
		return []byte(sig)
	}
	return nil
}

// Sign TBD
func (e RawEvent) Sign() error {
	return nil
}

// Tags TBD
func (e RawEvent) Tags() []Tag {
	var tags []Tag
	if m, ok := e["tags"].([]any); ok {
		for _, n := range m {
			if v, ok := n.([]any); ok {
				tags = append(tags, RawTag(v))
			}
		}
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
	return &MetadataEvent{
		"kind": EventKindMetadata,
	}
}

// MetadataEvent TBD
type MetadataEvent = RawEvent

// NewShortTextNoteEvent TBD
func NewShortTextNoteEvent() Event {
	return &ShortTextNoteEvent{
		"kind": EventKindShortTextNote,
	}
}

// ShortTextNoteEvent TBD
type ShortTextNoteEvent = RawEvent

// NewRecommendRelayEvent TBD
func NewRecommendRelayEvent() Event {
	return &RecommendRelayEvent{
		"kind": EventKindRecommendRelay,
	}
}

// RecommendRelayEvent TBD
type RecommendRelayEvent = RawEvent

// NewContactsEvent TBD
func NewContactsEvent() Event {
	return &ContactsEvent{
		"kind": EventKindContacts,
	}
}

// ContactsEvent TBD
type ContactsEvent = RawEvent

// NewEncryptedDirectMessagesEvent TBD
func NewEncryptedDirectMessagesEvent() Event {
	return &EncryptedDirectMessagesEvent{
		"kind": EventKindEncryptedDirectMessages,
	}
}

// EncryptedDirectMessagesEvent TBD
type EncryptedDirectMessagesEvent = RawEvent

// NewEventDeletionEvent TBD
func NewEventDeletionEvent() Event {
	return &EventDeletionEvent{
		"kind": EventKindEventDeletion,
	}
}

// EventDeletionEvent TBD
type EventDeletionEvent = RawEvent

// NewEventRepostsEvent TBD
func NewEventRepostsEvent() Event {
	return &EventRepostsEvent{
		"kind": EventKindReposts,
	}
}

// EventRepostsEvent TBD
type EventRepostsEvent = RawEvent

// NewReactionEvent TBD
func NewReactionEvent() Event {
	return &ReactionEvent{
		"kind": EventKindReaction,
	}
}

// ReactionEvent TBD
type ReactionEvent = RawEvent

// NewBadgeAwardEvent TBD
func NewBadgeAwardEvent() Event {
	return &BadgeAwardEvent{
		"kind": EventKindBadgeAward,
	}
}

// BadgeAwardEvent TBD
type BadgeAwardEvent = RawEvent

// NewChannelCreationEvent TBD
func NewChannelCreationEvent() Event {
	return &ChannelCreationEvent{
		"kind": EventKindChannelCreation,
	}
}

// ChannelCreationEvent TBD
type ChannelCreationEvent = RawEvent

func NewChannelMetadataEvent() Event {
	return &ChannelMetadataEvent{
		"kind": EventKindChannelMetadata,
	}
}

// ChannelMetadata TBD
type ChannelMetadataEvent = RawEvent

func NewChannelMessageEvent() Event {
	return &ChannelMessageEvent{
		"kind": EventKindChannelMessage,
	}
}

// ChannelMessage TBD
type ChannelMessageEvent = RawEvent

func NewChannelHideMessageEvent() Event {
	return &ChannelHideMessageEvent{
		"kind": EventKindChannelHideMessage,
	}
}

// ChannelHideMessage TBD
type ChannelHideMessageEvent = RawEvent

func NewChannelMuteUserEvent() Event {
	return &ChannelMuteUserEvent{
		"kind": EventKindChannelMuteUser,
	}
}

// ChannelMuteUser TBD
type ChannelMuteUserEvent = RawEvent

func NewReportingEvent() Event {
	return &ReportingEvent{
		"kind": EventKindReporting,
	}
}

// Reporting TBD
type ReportingEvent = RawEvent

func NewZapRequestEvent() Event {
	return &ZapRequestEvent{
		"kind": EventKindZapRequest,
	}
}

// ZapRequest TBD
type ZapRequestEvent = RawEvent

func NewZapEvent() Event {
	return &ZapEvent{
		"kind": EventKindZap,
	}
}

// Zap TBD
type ZapEvent = RawEvent

func NewMuteListEvent() Event {
	return &MuteListEvent{
		"kind": EventKindMuteList,
	}
}

// MuteList TBD
type MuteListEvent = RawEvent

func NewPinListEvent() Event {
	return &PinListEvent{
		"kind": EventKindPinList,
	}
}

// PinList TBD
type PinListEvent = RawEvent

func NewRelayListMetadataEvent() Event {
	return &RelayListMetadataEvent{
		"kind": EventKindRelayListMetadata,
	}
}

// RelayListMetadata TBD
type RelayListMetadataEvent = RawEvent

func NewClientAuthenticationEvent() Event {
	return &ClientAuthenticationEvent{
		"kind": EventKindClientAuthentication,
	}
}

// ClientAuthentication TBD
type ClientAuthenticationEvent = RawEvent

func NewNostrConnectEvent() Event {
	return &NostrConnectEvent{
		"kind": EventKindNostrConnect,
	}
}

// NostrConnect TBD
type NostrConnectEvent = RawEvent

func NewCategorizedPeopleListEvent() Event {
	return &CategorizedPeopleListEvent{
		"kind": EventKindCategorizedPeopleList,
	}
}

// CategorizedPeopleList TBD
type CategorizedPeopleListEvent = RawEvent

func NewCategorizedBookmarkListEvent() Event {
	return &CategorizedBookmarkListEvent{
		"kind": EventKindCategorizedBookmarkList,
	}
}

// CategorizedBookmarkList TBD
type CategorizedBookmarkListEvent = RawEvent

func NewProfileBadgesEvent() Event {
	return &ProfileBadgesEvent{
		"kind": EventKindProfileBadges,
	}
}

// ProfileBadges TBD
type ProfileBadgesEvent = RawEvent

func NewBadgeDefinitionEvent() Event {
	return &BadgeDefinitionEvent{
		"kind": EventKindBadgeDefinition,
	}
}

// BadgeDefinition TBD
type BadgeDefinitionEvent = RawEvent

func NewCreateOrUpdateStallEvent() Event {
	return &CreateOrUpdateStallEvent{
		"kind": EventKindCreateOrUpdateStall,
	}
}

// CreateOrUpdateStall TBD
type CreateOrUpdateStallEvent = RawEvent

func NewCreateOrUpdateProductEvent() Event {
	return &CreateOrUpdateProductEvent{
		"kind": EventKindCreateOrUpdateProduct,
	}
}

// CreateOrUpdateProduct TBD
type CreateOrUpdateProductEvent = RawEvent

func NewLongFormContentEvent() Event {
	return &LongFormContentEvent{
		"kind": EventKindLongFormContent,
	}
}

// LongFormContent TBD
type LongFormContentEvent = RawEvent

func NewApplicationSpecificDataEvent() Event {
	return &ApplicationSpecificDataEvent{
		"kind": EventKindApplicationSpecificData,
	}
}

// ApplicationSpecificData TBD
type ApplicationSpecificDataEvent = RawEvent
