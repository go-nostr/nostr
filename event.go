package nostr

import "encoding/json"

// EventKind represents the different types of events available.
type EventKind int64

const (
	EventKindMetadata                EventKind = 0     // Event for setting metadata
	EventKindShortTextNote           EventKind = 1     // Event for storing a text note
	EventKindRecommendRelay          EventKind = 2     // Event for recommending a server
	EventKindContacts                EventKind = 3     // Event for contact List and petnames
	EventKindEncryptedDirectMessages EventKind = 4     // Event for encrypted direct messages
	EventKindEventDeletion           EventKind = 5     // Event for deleting events
	EventKindReposts                 EventKind = 6     // Event for signaling to followers that another event is worth reading
	EventKindReaction                EventKind = 7     // Event for reacting to other notes
	EventKindBadgeAward              EventKind = 8     // Event for awarding badges to users
	EventKindChannelCreation         EventKind = 40    // Event for creating new channels
	EventKindChannelMetadata         EventKind = 41    // Event for setting channel metadata
	EventKindChannelMessage          EventKind = 42    // Event for posting messages in a channel
	EventKindChannelHideMessage      EventKind = 43    // Event for hiding messages in a channel
	EventKindChannelMuteUser         EventKind = 44    // Event for muting a user in a channel
	EventKindReporting               EventKind = 1984  // Event for reporting content or users
	EventKindZapRequest              EventKind = 9734  // Event for requesting a Zap action
	EventKindZap                     EventKind = 9735  // Event for performing a Zap action
	EventKindMuteList                EventKind = 10000 // Event for managing a mute list
	EventKindPinList                 EventKind = 10001 // Event for managing a pin list
	EventKindRelayListMetadata       EventKind = 10002 // Event for managing relay list metadata
	EventKindClientAuthentication    EventKind = 22242 // Event for client authentication process
	EventKindNostrConnect            EventKind = 24133 // Event for Nostr client connection
	EventKindCategorizedPeopleList   EventKind = 30000 // Event for managing categorized people list
	EventKindCategorizedBookmarkList EventKind = 30001 // Event for managing categorized bookmark list
	EventKindProfileBadges           EventKind = 30008 // Event for profile badges management
	EventKindBadgeDefinition         EventKind = 30009 // Event for defining badges
	EventKindCreateOrUpdateStall     EventKind = 30017 // Event for creating or updating a stall
	EventKindCreateOrUpdateProduct   EventKind = 30018 // Event for creating or updating a product
	EventKindLongFormContent         EventKind = 30023 // Event for posting long-form content
	EventKindApplicationSpecificData EventKind = 30078 // Event for managing application-specific data
)

// Event represents an event with a specified kind.
type Event interface {
	Marshal() ([]byte, error)
	Unmarshal(data []byte) error
	Validate() error
}

type BaseEvent struct {
	ID        string    `json:"id,omitempty"`
	PubKey    string    `json:"pub_key,omitempty"`
	CreatedAt int64     `json:"created_at,omitempty"`
	Kind      EventKind `json:"kind,omitempty"`
	Tags      []Tag     `json:"tags,omitempty"`
	Content   string    `json:"content,omitempty"`
	Signature string    `json:"signature,omitempty"`
}

// Marshal TBD
func (e *BaseEvent) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

// Unmarshal TBD
func (e *BaseEvent) Unmarshal(data []byte) error {
	return json.Unmarshal(data, e)
}

// Validate TBD
func (e *BaseEvent) Validate() error {
	return nil
}

// { TBD
// NewMetadataEvent TBD
func NewMetadataEvent() Event {
	return &MetadataEvent{&BaseEvent{
		Kind: EventKindMetadata,
	}}
}

// MetadataEvent TBD
type MetadataEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *MetadataEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *MetadataEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *MetadataEvent) Validate() error {
	return e.BaseEvent.Validate()
}

// NewShortTextNoteEvent TBD
func NewShortTextNoteEvent() Event {
	return &ShortTextNoteEvent{&BaseEvent{
		Kind: EventKindShortTextNote,
	}}
}

// ShortTextNoteEvent TBD
type ShortTextNoteEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *ShortTextNoteEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ShortTextNoteEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ShortTextNoteEvent) Validate() error {
	return e.BaseEvent.Validate()
}

// NewRecommendRelayEvent TBD
func NewRecommendRelayEvent() Event {
	return &RecommendRelayEvent{&BaseEvent{
		Kind: EventKindRecommendRelay,
	}}
}

// RecommendRelayEvent TBD
type RecommendRelayEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *RecommendRelayEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *RecommendRelayEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *RecommendRelayEvent) Validate() error {
	return e.BaseEvent.Validate()
}

// NewContactsEvent TBD
func NewContactsEvent() Event {
	return &ContactsEvent{&BaseEvent{
		Kind: EventKindContacts,
	}}
}

// ContactsEvent TBD
type ContactsEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *ContactsEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ContactsEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ContactsEvent) Validate() error {
	return e.BaseEvent.Validate()
}

// NewEncryptedDirectMessagesEvent TBD
func NewEncryptedDirectMessagesEvent() Event {
	return &EncryptedDirectMessagesEvent{&BaseEvent{
		Kind: EventKindEncryptedDirectMessages,
	}}
}

// EncryptedDirectMessagesEvent TBD
type EncryptedDirectMessagesEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *EncryptedDirectMessagesEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *EncryptedDirectMessagesEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *EncryptedDirectMessagesEvent) Validate() error {
	return e.BaseEvent.Validate()
}

// NewEventDeletionEvent TBD
func NewEventDeletionEvent() Event {
	return &EventDeletionEvent{&BaseEvent{
		Kind: EventKindEventDeletion,
	}}
}

// EventDeletionEvent TBD
type EventDeletionEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *EventDeletionEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *EventDeletionEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *EventDeletionEvent) Validate() error {
	return e.BaseEvent.Validate()
}

// NewEventRepostsEvent TBD
func NewEventRepostsEvent() Event {
	return &EventRepostsEvent{&BaseEvent{
		Kind: EventKindReposts,
	}}
}

// EventRepostsEvent TBD
type EventRepostsEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *EventRepostsEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *EventRepostsEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *EventRepostsEvent) Validate() error {
	return e.BaseEvent.Validate()
}

// NewReactionEvent TBD
func NewReactionEvent() Event {
	return &ReactionEvent{&BaseEvent{
		Kind: EventKindReaction,
	}}
}

// ReactionEvent TBD
type ReactionEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *ReactionEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ReactionEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ReactionEvent) Validate() error {
	return e.BaseEvent.Validate()
}

// NewBadgeAwardEvent TBD
func NewBadgeAwardEvent() Event {
	return &BadgeAwardEvent{&BaseEvent{
		Kind: EventKindBadgeAward,
	}}
}

// BadgeAwardEvent TBD
type BadgeAwardEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *BadgeAwardEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *BadgeAwardEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *BadgeAwardEvent) Validate() error {
	return e.BaseEvent.Validate()
}

// NewChannelCreationEvent TBD
func NewChannelCreationEvent() Event {
	return &ChannelCreationEvent{&BaseEvent{
		Kind: EventKindChannelCreation,
	}}
}

// ChannelCreationEvent TBD
type ChannelCreationEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *ChannelCreationEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ChannelCreationEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ChannelCreationEvent) Validate() error {
	return e.BaseEvent.Validate()
}

func NewChannelMetadataEvent() Event {
	return &ChannelMetadataEvent{&BaseEvent{
		Kind: EventKindChannelMetadata,
	}}
}

// ChannelMetadata TBD
type ChannelMetadataEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *ChannelMetadataEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ChannelMetadataEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ChannelMetadataEvent) Validate() error {
	return e.BaseEvent.Validate()
}

func NewChannelMessageEvent() Event {
	return &ChannelMessageEvent{&BaseEvent{
		Kind: EventKindChannelMessage,
	}}
}

// ChannelMessage TBD
type ChannelMessageEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *ChannelMessageEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ChannelMessageEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ChannelMessageEvent) Validate() error {
	return e.BaseEvent.Validate()
}

func NewChannelHideMessageEvent() Event {
	return &ChannelHideMessageEvent{&BaseEvent{
		Kind: EventKindChannelHideMessage,
	}}
}

// ChannelHideMessage TBD
type ChannelHideMessageEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *ChannelHideMessageEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ChannelHideMessageEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ChannelHideMessageEvent) Validate() error {
	return e.BaseEvent.Validate()
}

func NewChannelMuteUserEvent() Event {
	return &ChannelMuteUserEvent{&BaseEvent{
		Kind: EventKindChannelMuteUser,
	}}
}

// ChannelMuteUser TBD
type ChannelMuteUserEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *ChannelMuteUserEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ChannelMuteUserEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ChannelMuteUserEvent) Validate() error {
	return e.BaseEvent.Validate()
}

func NewReportingEvent() Event {
	return &ReportingEvent{&BaseEvent{
		Kind: EventKindReporting,
	}}
}

// Reporting TBD
type ReportingEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *ReportingEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ReportingEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ReportingEvent) Validate() error {
	return e.BaseEvent.Validate()
}

func NewZapRequestEvent() Event {
	return &ZapRequestEvent{&BaseEvent{
		Kind: EventKindZapRequest,
	}}
}

// ZapRequest TBD
type ZapRequestEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *ZapRequestEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ZapRequestEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ZapRequestEvent) Validate() error {
	return e.BaseEvent.Validate()
}

func NewZapEvent() Event {
	return &ZapEvent{&BaseEvent{
		Kind: EventKindZap,
	}}
}

// Zap TBD
type ZapEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *ZapEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ZapEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ZapEvent) Validate() error {
	return e.BaseEvent.Validate()
}

func NewMuteListEvent() Event {
	return &MuteListEvent{&BaseEvent{
		Kind: EventKindMuteList,
	}}
}

// MuteList TBD
type MuteListEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *MuteListEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *MuteListEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *MuteListEvent) Validate() error {
	return e.BaseEvent.Validate()
}

func NewPinListEvent() Event {
	return &PinListEvent{&BaseEvent{
		Kind: EventKindPinList,
	}}
}

// PinList TBD
type PinListEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *PinListEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *PinListEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *PinListEvent) Validate() error {
	return e.BaseEvent.Validate()
}

func NewRelayListMetadataEvent() Event {
	return &RelayListMetadataEvent{&BaseEvent{
		Kind: EventKindRelayListMetadata,
	}}
}

// RelayListMetadata TBD
type RelayListMetadataEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *RelayListMetadataEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *RelayListMetadataEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *RelayListMetadataEvent) Validate() error {
	return e.BaseEvent.Validate()
}

func NewClientAuthenticationEvent() Event {
	return &ClientAuthenticationEvent{&BaseEvent{
		Kind: EventKindClientAuthentication,
	}}
}

// ClientAuthentication TBD
type ClientAuthenticationEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *ClientAuthenticationEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ClientAuthenticationEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ClientAuthenticationEvent) Validate() error {
	return e.BaseEvent.Validate()
}

func NewNostrConnectEvent() Event {
	return &NostrConnectEvent{&BaseEvent{
		Kind: EventKindNostrConnect,
	}}
}

// NostrConnect TBD
type NostrConnectEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *NostrConnectEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *NostrConnectEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *NostrConnectEvent) Validate() error {
	return e.BaseEvent.Validate()
}

func NewCategorizedPeopleListEvent() Event {
	return &CategorizedPeopleListEvent{&BaseEvent{
		Kind: EventKindCategorizedPeopleList,
	}}
}

// CategorizedPeopleList TBD
type CategorizedPeopleListEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *CategorizedPeopleListEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *CategorizedPeopleListEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *CategorizedPeopleListEvent) Validate() error {
	return e.BaseEvent.Validate()
}

func NewCategorizedBookmarkListEvent() Event {
	return &CategorizedBookmarkListEvent{&BaseEvent{
		Kind: EventKindCategorizedBookmarkList,
	}}
}

// CategorizedBookmarkList TBD
type CategorizedBookmarkListEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *CategorizedBookmarkListEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *CategorizedBookmarkListEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *CategorizedBookmarkListEvent) Validate() error {
	return e.BaseEvent.Validate()
}

func NewProfileBadgesEvent() Event {
	return &ProfileBadgesEvent{&BaseEvent{
		Kind: EventKindProfileBadges,
	}}
}

// ProfileBadges TBD
type ProfileBadgesEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *ProfileBadgesEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ProfileBadgesEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ProfileBadgesEvent) Validate() error {
	return e.BaseEvent.Validate()
}

func NewBadgeDefinitionEvent() Event {
	return &BadgeDefinitionEvent{&BaseEvent{
		Kind: EventKindBadgeDefinition,
	}}
}

// BadgeDefinition TBD
type BadgeDefinitionEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *BadgeDefinitionEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *BadgeDefinitionEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *BadgeDefinitionEvent) Validate() error {
	return e.BaseEvent.Validate()
}

func NewCreateOrUpdateStallEvent() Event {
	return &CreateOrUpdateStallEvent{&BaseEvent{
		Kind: EventKindCreateOrUpdateStall,
	}}
}

// CreateOrUpdateStall TBD
type CreateOrUpdateStallEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *CreateOrUpdateStallEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *CreateOrUpdateStallEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *CreateOrUpdateStallEvent) Validate() error {
	return e.BaseEvent.Validate()
}

func NewCreateOrUpdateProductEvent() Event {
	return &CreateOrUpdateProductEvent{&BaseEvent{
		Kind: EventKindCreateOrUpdateProduct,
	}}
}

// CreateOrUpdateProduct TBD
type CreateOrUpdateProductEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *CreateOrUpdateProductEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *CreateOrUpdateProductEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *CreateOrUpdateProductEvent) Validate() error {
	return e.BaseEvent.Validate()
}

func NewLongFormContentEvent() Event {
	return &LongFormContentEvent{&BaseEvent{
		Kind: EventKindLongFormContent,
	}}
}

// LongFormContent TBD
type LongFormContentEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *LongFormContentEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *LongFormContentEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *LongFormContentEvent) Validate() error {
	return e.BaseEvent.Validate()
}

func NewApplicationSpecificDataEvent() Event {
	return &ApplicationSpecificDataEvent{&BaseEvent{
		Kind: EventKindApplicationSpecificData,
	}}
}

// ApplicationSpecificData TBD
type ApplicationSpecificDataEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *ApplicationSpecificDataEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ApplicationSpecificDataEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ApplicationSpecificDataEvent) Validate() error {
	return e.BaseEvent.Validate()
}
