package nostr_test

import (
	"fmt"
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
		// TODO: replace with concrete implementation of Event when available
		// {
		// 	name: "MUST create a new AuthMessage with given challenge and event",
		// 	args: args{
		// 		challenge: "test_challenge",
		// 		event:     nil,
		// 	},
		// 	expect: &nostr.AuthMessage{
		// 		Challenge: "test_challenge",
		// 		Event:     &nostr.Event{ID: "test_event_id"},
		// 	},
		// },
		{
			name: "MUST create a new AuthMessage with given challenge and nil event",
			args: args{
				challenge: "test_challenge",
				event:     nil,
			},
			expect: &nostr.AuthMessage{
				Challenge: "test_challenge",
				Event:     nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authMessage := nostr.NewAuthMessage(tt.args.challenge, tt.args.event)
			authMessageData, err := authMessage.Marshal()
			if err != nil {
				t.Fatal(err)
			}
			expecteData, err := tt.expect.Marshal()
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(authMessageData, expecteData) {
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
			name: "MUST successfully marshal AuthMessage to byte slice",
			args: args{
				challenge: "test_challenge",
				event:     nil,
			},
			expect: []byte("[\"AUTH\",\"test_challenge\"]"),
			err:    nil,
		},
		// {
		// 	name: "SHOULD return an error when marshaling AuthMessage with nil event",
		// 	args: args{
		// 		challenge: "test_challenge",
		// 		event:     nil,
		// 	},
		// 	expect: nil,
		// 	err:    fmt.Errorf("cannot marshal AuthMessage with nil event"),
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authMessage := nostr.NewAuthMessage(tt.args.challenge, tt.args.event)
			data, err := authMessage.Marshal()
			if (err != nil && tt.err == nil) && (err == nil && tt.err != nil) && (err.Error() != tt.err.Error()) {
				t.Fatalf("expected error: %+v, got error: %+v", tt.err, err)
				return
			}
			if !reflect.DeepEqual(data, tt.expect) {
				t.Fatalf("expected: %+v, got: %+v", tt.expect, data)
			}
			t.Logf("got: %+s", data)
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
			name: "MUST successfully unmarshal byte slice to AuthMessage",
			args: args{
				data: []byte("[\"AUTH\",\"test_challenge\"]"),
			},
			expect: &nostr.AuthMessage{
				Type:      nostr.MessageTypeAuth,
				Challenge: "test_challenge",
				Event:     nil,
			},
			err: nil,
		},
		{
			name: "SHOULD return an error when unmarshaling an invalid byte slice",
			args: args{
				data: []byte("invalid_data"),
			},
			expect: &nostr.AuthMessage{
				Type:      nostr.MessageTypeAuth,
				Challenge: "test_challenge",
				Event:     nil,
			},
			err: fmt.Errorf("cannot unmarshal AuthMessage: invalid byte slice"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authMessage := nostr.NewAuthMessage("test_challenge", nil)
			err := authMessage.Unmarshal(tt.args.data)
			if (err != nil && tt.err == nil) && (err == nil && tt.err != nil) && (err.Error() != tt.err.Error()) {
				t.Fatalf("expected error: %+v, got: %+v", tt.err, err)
			}
			if !reflect.DeepEqual(authMessage, tt.expect) {
				t.Fatalf("expected: %+v, got: %+v", tt.expect, authMessage)
			}
			t.Logf("got: %v", authMessage)
		})
	}
}
