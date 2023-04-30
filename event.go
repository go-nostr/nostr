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
	Content() []byte
	Get(key string) any
	ID() []byte
	Keys() []string
	Kind() int
	Marshal() ([]byte, error)
	PublicKey() []byte
	Set(key string, val any) error
	Sign() error
	Signature() []byte
	Tags() []Tag
	Unmarshal(data []byte) error
	Values() []any
}

func NewRawEvent() Event {
	return &RawEvent{}
}

// RawEvent TBD
type RawEvent map[string]json.RawMessage

// Content TBD
func (e *RawEvent) Content() []byte {
	var content string
	if err := json.Unmarshal((*e)["content"], &content); err != nil {
		return []byte{}
	}
	return []byte(content)
}

// Get TBD
func (e *RawEvent) Get(key string) any {
	var val any
	if err := json.Unmarshal((*e)[key], &val); err != nil {
		return nil
	}
	return val
}

// ID TBD
func (e *RawEvent) ID() []byte {
	var id string
	if err := json.Unmarshal((*e)["id"], &id); err != nil {
		return []byte{}
	}
	return []byte(id)
}

// Keys TBD
func (e *RawEvent) Keys() []string {
	var keys []string
	for key := range *e {
		keys = append(keys, key)
	}
	return keys
}

// Kind TBD
func (e *RawEvent) Kind() int {
	var kind int
	if err := json.Unmarshal((*e)["kind"], &kind); err != nil {
		return -1
	}
	return kind
}

// PublicKey TBD
func (e *RawEvent) PublicKey() []byte {
	var pubKey string
	if err := json.Unmarshal((*e)["pubkey"], &pubKey); err != nil {
		return []byte{}
	}
	return []byte(pubKey)
}

// Set TBD
func (e *RawEvent) Set(key string, val any) error {
	data, err := json.Marshal(val)
	if err != nil {
		return err
	}
	(*e)[key] = data
	return nil
}

// Signature TBD
func (e *RawEvent) Signature() []byte {
	var sig string
	if err := json.Unmarshal((*e)["sig"], &sig); err != nil {
		return []byte{}
	}
	return []byte(sig)
}

// Sign TBD
func (e *RawEvent) Sign() error {
	// TODO
	return nil
}

// Tags TBD
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

// Values TBD
func (e *RawEvent) Values() []any {
	var vals []any
	for _, v := range *e {
		vals = append(vals, v)
	}
	return vals
}

// NewMetadataEvent TBD
func NewMetadataEvent() Event {
	event := &MetadataEvent{}
	event.Set("kind", EventKindMetadata)
	return event
}

// MetadataEvent TBD
type MetadataEvent struct {
	*RawEvent
}

// NewShortTextNoteEvent TBD
func NewShortTextNoteEvent() Event {
	event := &ShortTextNoteEvent{}
	event.Set("kind", EventKindShortTextNote)
	return event
}

// ShortTextNoteEvent TBD
type ShortTextNoteEvent struct {
	*RawEvent
}

// NewRecommendRelayEvent TBD
func NewRecommendRelayEvent() Event {
	event := &RecommendRelayEvent{}
	event.Set("kind", EventKindRecommendRelay)
	return event
}

// RecommendRelayEvent TBD
type RecommendRelayEvent struct {
	*RawEvent
}

// NewContactsEvent TBD
func NewContactsEvent() Event {
	event := &ContactsEvent{}
	event.Set("kind", EventKindContacts)
	return event
}

// ContactsEvent TBD
type ContactsEvent struct {
	*RawEvent
}

// NewEncryptedDirectMessagesEvent TBD
func NewEncryptedDirectMessagesEvent() Event {
	event := &EncryptedDirectMessagesEvent{}
	event.Set("kind", EventKindEncryptedDirectMessages)
	return event
}

// EncryptedDirectMessagesEvent TBD
type EncryptedDirectMessagesEvent struct {
	*RawEvent
}

// NewEventDeletionEvent TBD
func NewEventDeletionEvent() Event {
	event := &EventDeletionEvent{}
	event.Set("kind", EventKindEventDeletion)
	return event
}

// EventDeletionEvent TBD
type EventDeletionEvent struct {
	*RawEvent
}

// NewEventRepostsEvent TBD
func NewEventRepostsEvent() Event {
	event := &EventRepostsEvent{}
	event.Set("kind", EventKindReposts)
	return event
}

// EventRepostsEvent TBD
type EventRepostsEvent struct {
	*RawEvent
}

// NewReactionEvent TBD
func NewReactionEvent() Event {
	event := &ReactionEvent{}
	event.Set("kind", EventKindReaction)
	return event
}

// ReactionEvent TBD
type ReactionEvent struct {
	*RawEvent
}

// NewBadgeAwardEvent TBD
func NewBadgeAwardEvent() Event {
	event := &BadgeAwardEvent{}
	event.Set("kind", EventKindBadgeAward)
	return event
}

// BadgeAwardEvent TBD
type BadgeAwardEvent struct {
	*RawEvent
}

// NewChannelCreationEvent TBD
func NewChannelCreationEvent() Event {
	event := &ChannelCreationEvent{}
	event.Set("kind", EventKindChannelCreation)
	return event
}

// ChannelCreationEvent TBD
type ChannelCreationEvent struct {
	*RawEvent
}

func NewChannelMetadataEvent() Event {
	event := &ChannelMetadataEvent{}
	event.Set("kind", EventKindChannelMetadata)
	return event
}

// ChannelMetadata TBD
type ChannelMetadataEvent struct {
	*RawEvent
}

func NewChannelMessageEvent() Event {
	event := &ChannelMessageEvent{}
	event.Set("kind", EventKindChannelMessage)
	return event
}

// ChannelMessage TBD
type ChannelMessageEvent struct {
	*RawEvent
}

func NewChannelHideMessageEvent() Event {
	event := &ChannelHideMessageEvent{}
	event.Set("kind", EventKindChannelHideMessage)
	return event
}

// ChannelHideMessage TBD
type ChannelHideMessageEvent struct {
	*RawEvent
}

func NewChannelMuteUserEvent() Event {
	event := &ChannelMuteUserEvent{}
	event.Set("kind", EventKindChannelMuteUser)
	return event
}

// ChannelMuteUser TBD
type ChannelMuteUserEvent struct {
	*RawEvent
}

func NewReportingEvent() Event {
	event := &ReportingEvent{}
	event.Set("kind", EventKindReporting)
	return event
}

// Reporting TBD
type ReportingEvent struct {
	*RawEvent
}

func NewZapRequestEvent() Event {
	event := &ZapRequestEvent{}
	event.Set("kind", EventKindZapRequest)
	return event
}

// ZapRequest TBD
type ZapRequestEvent struct {
	*RawEvent
}

func NewZapEvent() Event {
	event := &ZapEvent{}
	event.Set("kind", EventKindZap)
	return event
}

// Zap TBD
type ZapEvent struct {
	*RawEvent
}

func NewMuteListEvent() Event {
	event := &MuteListEvent{}
	event.Set("kind", EventKindMuteList)
	return event
}

// MuteList TBD
type MuteListEvent struct {
	*RawEvent
}

func NewPinListEvent() Event {
	event := &PinListEvent{}
	event.Set("kind", EventKindPinList)
	return event
}

// PinList TBD
type PinListEvent struct {
	*RawEvent
}

func NewRelayListMetadataEvent() Event {
	event := &RelayListMetadataEvent{}
	event.Set("kind", EventKindRelayListMetadata)
	return event
}

// RelayListMetadata TBD
type RelayListMetadataEvent struct {
	*RawEvent
}

func NewClientAuthenticationEvent() Event {
	event := &ClientAuthenticationEvent{}
	event.Set("kind", EventKindClientAuthentication)
	return event
}

// ClientAuthentication TBD
type ClientAuthenticationEvent struct {
	*RawEvent
}

func NewNostrConnectEvent() Event {
	event := &NostrConnectEvent{}
	event.Set("kind", EventKindNostrConnect)
	return event
}

// NostrConnect TBD
type NostrConnectEvent struct {
	*RawEvent
}

func NewCategorizedPeopleListEvent() Event {
	event := &CategorizedPeopleListEvent{}
	event.Set("kind", EventKindCategorizedPeopleList)
	return event
}

// CategorizedPeopleList TBD
type CategorizedPeopleListEvent struct {
	*RawEvent
}

func NewCategorizedBookmarkListEvent() Event {
	event := &CategorizedBookmarkListEvent{}
	event.Set("kind", EventKindCategorizedBookmarkList)
	return event
}

// CategorizedBookmarkList TBD
type CategorizedBookmarkListEvent struct {
	*RawEvent
}

func NewProfileBadgesEvent() Event {
	event := &ProfileBadgesEvent{}
	event.Set("kind", EventKindProfileBadges)
	return event
}

// ProfileBadges TBD
type ProfileBadgesEvent struct {
	*RawEvent
}

func NewBadgeDefinitionEvent() Event {
	event := &BadgeDefinitionEvent{}
	event.Set("kind", EventKindBadgeDefinition)
	return event
}

// BadgeDefinition TBD
type BadgeDefinitionEvent struct {
	*RawEvent
}

func NewCreateOrUpdateStallEvent() Event {
	event := &CreateOrUpdateStallEvent{}
	event.Set("kind", EventKindCreateOrUpdateStall)
	return event
}

// CreateOrUpdateStall TBD
type CreateOrUpdateStallEvent struct {
	*RawEvent
}

func NewCreateOrUpdateProductEvent() Event {
	event := &CreateOrUpdateProductEvent{}
	event.Set("kind", EventKindCreateOrUpdateProduct)
	return event
}

// CreateOrUpdateProduct TBD
type CreateOrUpdateProductEvent struct {
	*RawEvent
}

func NewLongFormContentEvent() Event {
	event := &LongFormContentEvent{}
	event.Set("kind", EventKindLongFormContent)
	return event
}

// LongFormContent TBD
type LongFormContentEvent struct {
	*RawEvent
}

func NewApplicationSpecificDataEvent() Event {
	event := &ApplicationSpecificDataEvent{}
	event.Set("kind", EventKindApplicationSpecificData)
	return event
}

// ApplicationSpecificData TBD
type ApplicationSpecificDataEvent struct {
	*RawEvent
}
