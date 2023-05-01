package nostr

// Filter is a struct that defines a set of criteria for filtering events.
type Filter struct {
	IDs        [][]byte `json:"ids,omitempty"`     // IDs specifies the event IDs to filter.
	Authors    [][]byte `json:"authors,omitempty"` // Authors specifies the event authors to filter.
	Kinds      []int    `json:"kinds,omitempty"`   // Kinds specifies the event kinds to filter.
	EventIDs   [][]byte `json:"#e,omitempty"`      // EventIDs specifies the event IDs to filter by.
	PublicKeys [][]byte `json:"#p,omitempty"`      // PublicKeys specifies the public keys to filter by.
	Since      int      `json:"since,omitempty"`   // Since specifies the starting timestamp for filtering events.
	Until      int      `json:"until,omitempty"`   // Until specifies the ending timestamp for filtering events.
	Limit      int      `json:"limit,omitempty"`   // Limit specifies the maximum number of events to return.
	Search     []byte   `json:"search,omitempty"`  // Search specifies a search term to filter events by.
}
