package nostr

// Admission TBD
type Admission struct {
	Kinds  int      `json:"kinds,omitempty"`
	Amount int      `json:"amount,omitempty"`
	Unit   []string `json:"unit,omitempty"`
}
