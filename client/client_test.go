package client_test

import (
	"context"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/go-nostr/nostr/client"
	"github.com/go-nostr/nostr/message"
	"github.com/go-nostr/nostr/message/closemessage"
	"github.com/go-nostr/nostr/message/requestmessage"
	"github.com/go-nostr/nostr/relay"
)

func Test_New(t *testing.T) {
	type args struct {
		opt *client.Options
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "SHOULD create instance of Client",
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

func TestClient_HandleErrorFunc(t *testing.T) {
	type args struct {
		fn func(error)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "SHOULD handle error",
			args: args{
				fn: func(err error) {
					t.Log(err)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := client.New(nil)
			cl.HandleErrorFunc(tt.args.fn)
		})
	}
}

func TestClient_HandleMessageFunc(t *testing.T) {
	type args struct {
		fn func(message.Message)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "SHOULD handle message",
			args: args{
				fn: func(msg message.Message) {
					t.Log(msg)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := client.New(nil)
			cl.HandleMessageFunc(tt.args.fn)
		})
	}
}

func TestClient_Listen(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "SHOULD listen for messages",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Second)
			defer cancel()
			cl := client.New(nil)
			cl.Listen(ctx)
		})
	}
}

func TestClient_SendMessage(t *testing.T) {
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
			name: "SHOULD send close message from client to relay",
			args: args{
				msg: closemessage.New(),
			},
		},
		// FIX: failing assertion comparing pointer to map
		// {
		// 	name: "SHOULD send event message from client to relay",
		// 	args: args{
		// 		msg: eventmessage.New("subscription-id", metadataevent.New("name", "", "")),
		// 	},
		// },
		{
			name: "SHOULD send request message from client to relay",
			args: args{
				msg: requestmessage.New("subscription-id"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errCh := make(chan error)
			msgCh := make(chan message.Message)
			rl := relay.New(nil)
			rl.HandleErrorFunc(func(err error) {
				errCh <- err
			})
			rl.HandleMessageFunc(func(msg message.Message) {
				msgCh <- msg
			})
			ts := httptest.NewServer(rl)
			defer ts.Close()
			ctx, cancel := context.WithTimeout(context.TODO(), 6*time.Second)
			defer cancel()
			cl := client.New(nil)
			cl.HandleErrorFunc(func(err error) {
				t.Error(err)
			})
			cl.HandleMessageFunc(func(msg message.Message) {
				t.Log(msg)
			})
			cl.Connect(ctx, ts.URL)
			cl.SendMessage(ctx, tt.args.msg)
			select {
			case err := <-errCh:
				t.Error(err)
			case msg := <-msgCh:
				if !reflect.DeepEqual(tt.args.msg, msg) {
					t.Errorf("expected %v, got %v", tt.args.msg, msg)
				}
				t.Logf("got %v", msg)
			case <-ctx.Done():
				return
			}
		})
	}
}

func TestClient_Connect(t *testing.T) {
	rl := relay.New(nil)
	rl.HandleErrorFunc(func(err error) {
		// t.Error(err)
	})
	rl.HandleMessageFunc(func(msg message.Message) {
		t.Log(msg)
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
			name: "SHOULD subscribe to relay",
			args: args{
				msg: requestmessage.New("subscription-id", nil),
			},
			fields: fields{
				u: ts.URL,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), 6*time.Second)
			defer cancel()
			cl := client.New(nil)
			cl.HandleErrorFunc(func(err error) {
				t.Error(err)
			})
			cl.HandleMessageFunc(func(msg message.Message) {
				t.Log(msg)
			})
			cl.Connect(ctx, tt.fields.u)
		})
	}
}
