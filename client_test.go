package nostr_test

import (
	"context"
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
			cl := nostr.NewClient(nil)
			if cl == nil {
				t.Fatalf("expected %v, to be not nil", cl)
			}
			t.Logf("got %v", cl)
		})
	}
}

func TestClient_Publish(t *testing.T) {
	messChan := make(chan nostr.Message)
	r := nostr.NewRelay(nil)
	r.HandleMessageFunc(nostr.MessageTypeRequest, func(mess nostr.Message) {
		byt, err := mess.Marshal()
		if err != nil {
			t.Errorf("error handling message request from test %v", err)
			return
		}
		t.Logf("handled from test %s", byt)
		messChan <- mess
	})
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
			cl := nostr.NewClient(nil)
			if cl == nil {
				t.Fatalf("expected %v, to be not nil", cl)
			}
			if err := cl.Subscribe(context.TODO(), ts.URL); err != nil {
				t.Fatalf("error: %v", err)
			}
			if err := cl.Publish(tt.args.mess); err != nil {
				t.Fatalf("error: %v", err)
			}
			mess := <-messChan
			if mess == nil {
				t.Fatalf("expected message to not be nil")
			}
			t.Logf("recieved message %v", mess)
		})
	}
}

func TestClient_Subscribe(t *testing.T) {
	r := nostr.NewRelay(nil)
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
			cl := nostr.NewClient(nil)
			if cl == nil {
				t.Fatalf("expected %v, to be not nil", cl)
			}
			if err := cl.Subscribe(context.TODO(), ts.URL); err != nil {
				t.Fatalf("Error: %v", err)
			}
			t.Log(cl)
		})
	}
}
