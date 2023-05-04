package client_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/go-nostr/nostr/client"
	"github.com/go-nostr/nostr/message"
	"github.com/go-nostr/nostr/message/requestmessage"
	"github.com/go-nostr/nostr/relay"
)

func Test_NewClient(t *testing.T) {
	type args struct {
		opt *client.Options
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "SHOULD create client",
			args: args{
				opt: &client.Options{
					ReadLimit: 6.4e+7,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := client.New(tt.args.opt)
			if cl == nil {
				t.Errorf("got nil: %v", cl)
			}
			t.Logf("got: %v", cl)
		})
	}
}

func TestClient_Publish(t *testing.T) {
	errChan := make(chan error)
	msgChan := make(chan message.Message)
	rl := relay.New(nil)
	rl.HandleErrorFunc(func(err error) {
		errChan <- err
	})
	rl.HandleMessageFunc(func(msg message.Message) {
		msgChan <- msg
	})
	ts := httptest.NewServer(rl)
	defer ts.Close()
	type args struct {
		msg message.Message
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
				msg: requestmessage.New("asdf", &requestmessage.Filter{}),
			},
			fields: fields{
				addr: "ws://0.0.0.0:3001",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()
			cl := client.New(nil)
			if cl == nil {
				t.Fatalf("expected %v, to be not nil", cl)
			}
			cl.Subscribe(ctx, ts.URL)
			cl.Publish(ctx, tt.args.msg)
			select {
			case err := <-errChan:
				t.Fatal(err)
			case msg := <-msgChan:
				if msg == nil {
					t.Fatalf("expected message to not be nil")
				}
				t.Logf("recieved message %v", msg)
			}
		})
	}
}

// func TestClient_Subscribe(t *testing.T) {
// 	r := nostr.NewRelay(nil)
// 	ts := httptest.NewServer(r)
// 	defer ts.Close()
// 	tests := []struct {
// 		name string
// 	}{
// 		{
// 			name: "SHOULD subscribe client to relay",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			cl := nostr.NewClient(nil)
// 			if cl == nil {
// 				t.Fatalf("expected %v, to be not nil", cl)
// 			}
// 			if err := cl.Subscribe(context.TODO(), ts.URL); err != nil {
// 				t.Fatalf("Error: %v", err)
// 			}
// 			t.Log(cl)
// 		})
// 	}
// }
