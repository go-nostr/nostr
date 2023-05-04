package relay

type InternetIdentifier struct {
	Names  []string `json:"names,omitempty"`
	Relays []string `json:"relays,omitempty"`
}
