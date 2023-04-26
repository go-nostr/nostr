package nostr

type NIP int64

const (
	NIP1BasicProtocolFlow                NIP = 1
	NIP2ContactListAndPetnames           NIP = 2
	NIP3OpenTimestampsAttestations       NIP = 3
	NIP4EncryptedDirectMessage           NIP = 4
	NIP5MappingNostrKeysToDNS            NIP = 5
	NIP6BasicKeyDerivation               NIP = 6
	NIP7WindowNostrCapability            NIP = 7
	NIP8HandlingMentionsDeprecated       NIP = 8
	NIP9EventDeletion                    NIP = 9
	NIP10ConventionsForClients           NIP = 10
	NIP11RelayInformationDocument        NIP = 11
	NIP12GenericTagQueries               NIP = 12
	NIP13ProofOfWork                     NIP = 13
	NIP14SubjectTagInTextEvents          NIP = 14
	NIP15NostrMarketplace                NIP = 15
	NIP16EventTreatment                  NIP = 16
	NIP18Reposts                         NIP = 18
	NIP19Bech32EncodedEntities           NIP = 19
	NIP20CommandResults                  NIP = 20
	NIP21NostrURLScheme                  NIP = 21
	NIP22EventCreatedAtLimits            NIP = 22
	NIP23LongFormContent                 NIP = 23
	NIP25Reactions                       NIP = 25
	NIP26DelegatedEventSigning           NIP = 26
	NIP27TextNoteReferences              NIP = 27
	NIP28PublicChat                      NIP = 28
	NIP33ParameterizedReplaceableEvents  NIP = 33
	NIP36SensitiveContent                NIP = 36
	NIP39ExternalIdentitiesInProfiles    NIP = 39
	NIP40ExpirationTimestamp             NIP = 40
	NIP42AuthenticationOfClientsToRelays NIP = 42
	NIP45CountingResults                 NIP = 45
	NIP46NostrConnect                    NIP = 46
	NIP50KeywordsFilter                  NIP = 50
	NIP51Lists                           NIP = 51
	NIP56Reporting                       NIP = 56
	NIP57LightningZaps                   NIP = 57
	NIP58Badges                          NIP = 58
	NIP65RelayListMetadata               NIP = 65
	NIP78ApplicationSpecificData         NIP = 78
	NIP94FileMetadata                    NIP = 94
)
