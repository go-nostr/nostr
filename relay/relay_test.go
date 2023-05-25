package relay_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/go-nostr/nostr/client"
	"github.com/go-nostr/nostr/message"
	"github.com/go-nostr/nostr/message/eosemessage"
	"github.com/go-nostr/nostr/message/okmessage"
	"github.com/go-nostr/nostr/relay"
)

func Test_New(t *testing.T) {
	type args struct {
		opt *relay.Options
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				opt: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rl := relay.New(tt.args.opt)
			if rl == nil {
				t.Logf("expected instance of Relay, got nil")
			}
			t.Logf("got %v", rl)
		})
	}
	r := relay.New(nil)
	if r == nil {
		t.Fatal("New should not return nil")
	}
}

func TestRelay_HandleErrorFunc(t *testing.T) {
	type args struct {
		errFn func(error)
	}
	type fields struct {
		rl *relay.Relay
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "SHOULD handle error",
			fields: fields{
				rl: relay.New(nil),
			},
			args: args{
				errFn: func(err error) {
					t.Log(err)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.rl.HandleErrorFunc(tt.args.errFn)
		})
	}
}

func TestRelay_HandleInformationDocumentFunc(t *testing.T) {
	type args struct {
		informationDocumentFn func() (*relay.InformationDocument, error)
	}
	type fields struct {
		rl *relay.Relay
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "SHOULD handle information document",
			fields: fields{
				rl: relay.New(nil),
			},
			args: args{
				informationDocumentFn: func() (*relay.InformationDocument, error) {
					return &relay.InformationDocument{}, nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.rl.HandleInformationDocumentFunc(tt.args.informationDocumentFn)
		})
	}
}

func TestRelay_HandleInternetIdentifierFunc(t *testing.T) {
	type args struct {
		internetIdentiferFn func(string) (*relay.InternetIdentifier, error)
	}
	type fields struct {
		rl *relay.Relay
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "SHOULD handle information document",
			fields: fields{
				rl: relay.New(nil),
			},
			args: args{
				internetIdentiferFn: func(name string) (*relay.InternetIdentifier, error) {
					return &relay.InternetIdentifier{}, nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.rl.HandleInternetIdentifierFunc(tt.args.internetIdentiferFn)
		})
	}
}

func TestRelay_SendMessage(t *testing.T) {
	type args struct {
		ctx context.Context
		msg message.Message
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "SHOULD send EOSE message",
			args: args{
				ctx: context.TODO(),
				msg: eosemessage.New(),
			},
		},
		{
			name: "SHOULD send EVENT message",
			args: args{
				ctx: context.TODO(),
				msg: eosemessage.New(),
			},
		},
		{
			name: "SHOULD send NOTICE message",
			args: args{
				ctx: context.TODO(),
				msg: eosemessage.New(),
			},
		},
		{
			name: "SHOULD send OK message",
			args: args{
				ctx: context.TODO(),
				msg: okmessage.New("asdf", true, "ok"),
			},
		},
	}
	for _, tt := range tests {
		cl := client.New(nil)
		rl := relay.New(nil)
		ts := httptest.NewServer(rl)
		defer ts.Close()
		cl.Connect(context.TODO(), ts.URL)
		cl.HandleErrorFunc(func(err error) {
			t.Error(err)
		})
		cl.HandleMessageFunc(func(msg message.Message) {
			if !reflect.DeepEqual(tt.args.msg, msg) {
				t.Errorf("expected %v, got %v", tt.args.msg, msg)
			}
			t.Logf("got %v", msg)
		})
		t.Run(tt.name, func(t *testing.T) {
			rl.SendMessage(tt.args.ctx, tt.args.msg)
		})
		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		defer cancel()
		cl.Listen(ctx)
	}
}

func TestRelay_ServeHTTP(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type fields struct {
		rl *relay.Relay
	}
	tests := []struct {
		name   string
		args   args
		fields fields
	}{
		{
			name: "SHOULD get websocket connection",
			args: args{
				w: httptest.NewRecorder(),
				r: (func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/", nil)
					return req
				})(),
			},
			fields: fields{
				rl: relay.New(nil),
			},
		},
		{
			name: "SHOULD get internet identifier",
			args: args{
				w: httptest.NewRecorder(),
				r: (func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/.well-known/nostr.json?name=\"bob\"", nil)
					return req
				})(),
			},
			fields: fields{
				rl: relay.New(nil),
			},
		},
		{
			name: "SHOULD get information document",
			args: args{
				w: httptest.NewRecorder(),
				r: (func() *http.Request {
					req, _ := http.NewRequest(http.MethodGet, "/", nil)
					req.Header.Add("Accept", "application/nostr+json")
					return req
				})(),
			},
			fields: fields{
				rl: relay.New(nil),
			},
		},
	}
	for _, tt := range tests {
		tt.fields.rl.HandleErrorFunc(func(err error) {
			t.Error(err)
		})
		tt.fields.rl.HandleInformationDocumentFunc(func() (*relay.InformationDocument, error) {
			return &relay.InformationDocument{}, nil
		})
		tt.fields.rl.HandleInternetIdentifierFunc(func(name string) (*relay.InternetIdentifier, error) {
			return &relay.InternetIdentifier{
				Names: []string{name},
			}, nil
		})
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.rl.ServeHTTP(tt.args.w, tt.args.r)

			t.Log(tt.args.r)
		})
	}
}
