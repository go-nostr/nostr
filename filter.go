package nostr

// Filter is a struct that defines a set of criteria for filtering events.
type Filter struct {
	IDs        [][]byte `json:"ids,omitempty"`     // The IDs of the events to filter
	Authors    [][]byte `json:"authors,omitempty"` // The authors of the events to filter
	Kinds      []int    `json:"kinds,omitempty"`   // The kinds of the events to filter
	EventIDs   [][]byte `json:"#e,omitempty"`      // The event IDs to filter by
	PublicKeys [][]byte `json:"#p,omitempty"`      // The public keys to filter by
	Since      int      `json:"since,omitempty"`   // The starting timestamp for filtering events
	Until      int      `json:"until,omitempty"`   // The ending timestamp for filtering events
	Limit      int      `json:"limit,omitempty"`   // The maximum number of events to return
	Search     []byte   `json:"search,omitempty"`  // TBD
}
