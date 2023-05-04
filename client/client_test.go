package client_test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

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
		fmt.Println(msg)
		msgChan <- msg
	})
	ts := httptest.NewServer(rl)
	defer ts.Close()
	type args struct {
		msg message.Message
	}
	type fields struct {
		u string
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
				u: ts.URL,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
			defer cancel()
			cl := client.New(nil)
			if cl == nil {
				t.Fatalf("expected %v, to be not nil", cl)
			}
			cl.Subscribe(ctx, tt.fields.u)
			cl.Publish(ctx, tt.args.msg)
			if err := cl.Listen(ctx); err != nil {
				t.Error(err)
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
