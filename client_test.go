package nostr_test

import (
	"net/http/httptest"
	"testing"

	"github.com/go-nostr/nostr"
)

func Test_NewClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "SHOULD create client",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mess := nostr.NewRequestMessage("asdf", &nostr.Filter{})
			cl := nostr.NewClient()
			if cl == nil {
				t.Fatalf("expected %v, to be not nil", cl)
			}
			t.Logf("got %v", cl)
		})
	}
}

func TestClient_Publish(t *testing.T) {
	r := nostr.NewRelay()
	ts := httptest.NewServer(r)
	defer ts.Close()
	type args struct {
		mess nostr.Message
	}
	type fields struct {
		addr string
	}
	tests := []struct {
		name   string
		args   args
		fields fields
	}{
		{
			name: "SHOULD publish message to relay",
			args: args{
				mess: nostr.NewRequestMessage("asdf", &nostr.Filter{}),
			},
			fields: fields{
				addr: "ws://0.0.0.0:3001",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := nostr.NewClient()
			if cl == nil {
				t.Fatalf("expected %v, to be not nil", cl)
			}
			if err := cl.Subscribe(ts.URL); err != nil {
				t.Fatalf("error: %v", err)
			}
			if err := cl.Publish(tt.args.mess); err != nil {
				t.Fatalf("error: %v", err)
			}
			t.Logf("published mess: %v", tt.args.mess)
			t.Logf("%v", cl)
		})
	}
}

func TestClient_Subscribe(t *testing.T) {
	r := nostr.NewRelay()
	ts := httptest.NewServer(r)
	defer ts.Close()
	tests := []struct {
		name string
	}{
		{
			name: "SHOULD subscribe client to relay",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := nostr.NewClient()
			if cl == nil {
				t.Fatalf("expected %v, to be not nil", cl)
			}
			if err := cl.Subscribe(ts.URL); err != nil {
				t.Fatalf("Error: %v", err)
			}
			t.Log(cl)
		})
	}
}
