package nostr

type NIP int64

const (
	NIP1BasicProtocolFlow                NIP = 1  // NIP1BasicProtocolFlow represents NIP-01: Basic protocol flow description
	NIP2ContactListAndPetnames           NIP = 2  // NIP2ContactListAndPetnames represents NIP-02: Contact List and Petnames
	NIP3OpenTimestampsAttestations       NIP = 3  // NIP3OpenTimestampsAttestations represents NIP-03: OpenTimestamps Attestations for Events
	NIP4EncryptedDirectMessage           NIP = 4  // NIP4EncryptedDirectMessage represents NIP-04: Encrypted Direct Message
	NIP5MappingNostrKeysToDNS            NIP = 5  // NIP5MappingNostrKeysToDNS represents NIP-05: Mapping Nostr keys to DNS-based internet identifiers
	NIP6BasicKeyDerivation               NIP = 6  // NIP6BasicKeyDerivation represents NIP-06: Basic key derivation from mnemonic seed phrase
	NIP7WindowNostrCapability            NIP = 7  // NIP7WindowNostrCapability represents NIP-07: `window.nostr` capability for web browsers
	NIP8HandlingMentionsDeprecated       NIP = 8  // NIP8HandlingMentionsDeprecated represents NIP-08: Handling Mentions](08.md) --- **unrecommended**: deprecated in favor of [NIP-27
	NIP9EventDeletion                    NIP = 9  // NIP9EventDeletion represents NIP-09: Event Deletion
	NIP10ConventionsForClients           NIP = 10 // NIP10ConventionsForClients represents NIP-10: Conventions for clients' use of `e` and `p` tags in text events
	NIP11RelayInformationDocument        NIP = 11 // NIP11RelayInformationDocument represents NIP-11: Relay Information Document
	NIP12GenericTagQueries               NIP = 12 // NIP12GenericTagQueries represents NIP-12: Generic Tag Queries
	NIP13ProofOfWork                     NIP = 13 // NIP13ProofOfWork represents NIP-13: Proof of Work
	NIP14SubjectTagInTextEvents          NIP = 14 // NIP14SubjectTagInTextEvents represents NIP-14: Subject tag in text events.
	NIP15NostrMarketplace                NIP = 15 // NIP15NostrMarketplace represents NIP-15: Nostr Marketplace (for resilient marketplaces)
	NIP16EventTreatment                  NIP = 16 // NIP16EventTreatment represents NIP-16: Event Treatment
	NIP18Reposts                         NIP = 18 // NIP18Reposts represents NIP-18: Reposts
	NIP19Bech32EncodedEntities           NIP = 19 // NIP19Bech32EncodedEntities represents NIP-19: bech32-encoded entities
	NIP20CommandResults                  NIP = 20 // NIP20CommandResults represents NIP-20: Command Results
	NIP21NostrURLScheme                  NIP = 21 // NIP21NostrURLScheme represents NIP-21: `nostr:` URL scheme
	NIP22EventCreatedAtLimits            NIP = 22 // NIP22EventCreatedAtLimits represents NIP-22: Event `created_at` Limits
	NIP23LongFormContent                 NIP = 23 // NIP23LongFormContent represents NIP-23: Long-form Content
	NIP25Reactions                       NIP = 25 // NIP25Reactions represents NIP-25: Reactions
	NIP26DelegatedEventSigning           NIP = 26 // NIP26DelegatedEventSigning represents NIP-26: Delegated Event Signing
	NIP27TextNoteReferences              NIP = 27 // NIP27TextNoteReferences represents NIP-27: Text Note References
	NIP28PublicChat                      NIP = 28 // NIP28PublicChat represents NIP-28: Public Chat
	NIP33ParameterizedReplaceableEvents  NIP = 33 // NIP33ParameterizedReplaceableEvents represents NIP-33: Parameterized Replaceable Events
	NIP36SensitiveContent                NIP = 36 // NIP36SensitiveContent represents NIP-36: Sensitive Content
	NIP39ExternalIdentitiesInProfiles    NIP = 39 // NIP39ExternalIdentitiesInProfiles represents NIP-39: External Identities in Profiles
	NIP40ExpirationTimestamp             NIP = 40 // NIP40ExpirationTimestamp represents NIP-40: Expiration Timestamp
	NIP42AuthenticationOfClientsToRelays NIP = 42 // NIP42AuthenticationOfClientsToRelays represents NIP-42: Authentication of clients to relays
	NIP45CountingResults                 NIP = 45 // NIP45CountingResults represents NIP-45: Counting results
	NIP46NostrConnect                    NIP = 46 // NIP46NostrConnect represents NIP-46: Nostr Connect
	NIP50KeywordsFilter                  NIP = 50 // NIP50KeywordsFilter represents NIP-50: Keywords filter
	NIP51Lists                           NIP = 51 // NIP51Lists represents NIP-51: Lists
	NIP56Reporting                       NIP = 56 // NIP56Reporting represents NIP-56: Reporting
	NIP57LightningZaps                   NIP = 57 // NIP57LightningZaps represents NIP-57: Lightning Zaps
	NIP58Badges                          NIP = 58 // NIP58Badges represents NIP-58: Badges
	NIP65RelayListMetadata               NIP = 65 // NIP65RelayListMetadata represents NIP-65: Relay List Metadata
	NIP78ApplicationSpecificData         NIP = 78 // NIP78ApplicationSpecificData represents NIP-78: Application-specific data
	NIP94FileMetadata                    NIP = 94 // NIP94FileMetadata represents NIP-94: File Metadata
)
