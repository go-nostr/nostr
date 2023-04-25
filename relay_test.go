package nostr_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-nostr/nostr"
)

func Test_New(t *testing.T) {
	r := nostr.NewRelay()
	if r == nil {
		t.Fatal("New should not return nil")
	}
}

func TestRelay_WellKnownNostrEndpoint(t *testing.T) {
	r := nostr.NewRelay()
	ts := httptest.NewServer(r)
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/.well-known/nostr.json", ts.URL))
	if err != nil {
		t.Fatalf("Error fetching nostr.json: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", resp.StatusCode)
	}

	var data struct {
		Names  []string `json:"names,omitempty"`
		Relays []string `json:"relays,omitempty"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		t.Fatalf("Error decoding JSON response: %v", err)
	}

	if len(data.Names) != 1 || data.Names[0] != "bob" {
		t.Fatalf("Expected names to be [\"bob\"], got %v", data.Names)
	}
}
