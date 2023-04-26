package nostr_test

import (
	"reflect"
	"testing"

	"github.com/go-nostr/nostr"
)

func Test_NewAuthMessage(t *testing.T) {
	type args struct {
		challenge string
		event     *nostr.Event
	}
	tests := []struct {
		name   string
		args   args
		expect *nostr.AuthMessage
	}{
		{
			name: "SHOULD create instance of AuthMessage",
			args: args{
				challenge: "abc",
			},
			expect: &nostr.AuthMessage{
				Type:      nostr.MessageTypeAuth,
				Challenge: "abc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authMessage := nostr.NewAuthMessage(tt.args.challenge, tt.args.event)

			if !reflect.DeepEqual(authMessage, tt.expect) {
				t.Errorf("expected %+v, got %+v", authMessage, tt.expect)
			}

			t.Logf("got %+v", authMessage)
		})
	}
}

func TestAuthMessage_Marshal(t *testing.T) {
	type args struct {
		challenge string
		event     *nostr.Event
	}
	tests := []struct {
		name   string
		args   args
		expect []byte
		err    error
	}{
		{
			name: "SHOULD marshl AuthMessage",
			args: args{
				challenge: "abc",
			},
			expect: []byte("[\"AUTH\",\"abc\"]"),
			err:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authMessage := nostr.NewAuthMessage(tt.args.challenge, tt.args.event)
			data, err := authMessage.Marshal()
			if (err != nil && tt.err == nil) && (err == nil && tt.err != nil) && (err.Error() != tt.err.Error()) {
				t.Fatalf("expected error: %+v, got error: %+v", tt.err, err)
			}
			if !reflect.DeepEqual(data, tt.expect) {
				t.Fatalf("expected: %+v, got: %+v", tt.expect, data)
			}
			t.Logf("got: %+v", data)
		})
	}
}

func TestAuthMessage_Unmarshal(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		args   args
		expect *nostr.AuthMessage
		err    error
	}{
		{
			name: "SHOULD unmarshal AuthMessage",
			args: args{
				data: []byte("[\"AUTH\",\"abc\"]"),
			},
			expect: &nostr.AuthMessage{
				Type:      nostr.MessageTypeAuth,
				Challenge: "abc",
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authMessage := &nostr.AuthMessage{}
			err := authMessage.Unmarshal(tt.args.data)
			if (err != nil && tt.err == nil) && (err == nil && tt.err != nil) && (err.Error() != tt.err.Error()) {
				t.Fatalf("expected error: %+v, got: %+v", tt.err, err)
			}
			if !reflect.DeepEqual(authMessage, tt.expect) {
				t.Fatalf("expected error: %+v, got: %+v", tt.expect, authMessage)
			}
			t.Logf("got: %+v", authMessage)
		})
	}
}
