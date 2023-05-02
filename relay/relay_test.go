package relay_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-nostr/nostr/relay"
)

func Test_New(t *testing.T) {
	r := relay.New(nil)
	if r == nil {
		t.Fatal("New should not return nil")
	}
}

func TestRelay_WellKnownNostrEndpoint(t *testing.T) {
	r := relay.New(nil)
	ts := httptest.NewServer(r)
	defer ts.Close()
	type args struct {
		names map[string]string
	}
	tests := []struct {
		name   string
		args   args
		expect map[string]string
	}{
		// TODO
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Get(fmt.Sprintf("%s/.well-known/nostr.json", ts.URL))
			if err != nil {
				t.Fatalf("Error fetching nostr.json: %v", err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				t.Fatalf("Expected status 200, got %d", resp.StatusCode)
			}
			if err := json.NewDecoder(resp.Body).Decode(&tt.args.names); err != nil {
				t.Fatalf("Error decoding JSON response: %v", err)
			}
			if reflect.DeepEqual(tt.args.names, tt.expect) {
				t.Fatalf("Expected: %v, got: %v", tt.expect, tt.args.names)
			}
		})
	}
}
