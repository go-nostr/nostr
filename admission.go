package nostr

// Admission represents the admission requirements for accessing a Nostr relay.
type Admission struct {
	Kinds  []int    `json:"kinds,omitempty"`  // Kinds specifies the types of events allowed.
	Amount int      `json:"amount,omitempty"` // Amount indicates the payment required for accessing paid relays.
	Unit   []string `json:"unit,omitempty"`   // Unit denotes the currency unit used for payment.
}
