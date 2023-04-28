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
		event     nostr.Event
	}
	tests := []struct {
		name   string
		args   args
		expect *nostr.AuthMessage
	}{
		{
			name: "MUST create a new AuthMessage with given challenge and nil event",
			args: args{
				challenge: "test-challenge",
				event:     nil,
			},
			expect: &nostr.AuthMessage{nostr.BaseMessage{string(nostr.MessageTypeAuth), "test-challenge"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authMessage := nostr.NewAuthMessage(tt.args.challenge, tt.args.event)
			if !reflect.DeepEqual(authMessage, tt.expect) {
				t.Errorf("expected %s, got %s", authMessage, tt.expect)
			}
			t.Logf("got %+v", authMessage)
		})
	}
}

func TestAuthMessage_Marshal(t *testing.T) {
	type args struct {
		challenge string
		event     nostr.Event
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
				challenge: "test-challenge",
				event:     nil,
			},
			expect: []byte("[\"AUTH\",\"test-challenge\"]"),
			err:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authMessage := nostr.NewAuthMessage(tt.args.challenge, tt.args.event)
			data, err := authMessage.Marshal()
			if (err != nil && tt.err == nil) && (err == nil && tt.err != nil) && (err.Error() != tt.err.Error()) {
				t.Fatalf("expected error: %s, got error: %s", tt.err, err)
				return
			}
			if !reflect.DeepEqual(data, tt.expect) {
				t.Fatalf("expected: %s, got: %s", tt.expect, data)
			}
			t.Logf("got: %+s", data)
		})
	}
}

func TestAuthMessage_Unmarshal(t *testing.T) {
	type args struct {
		data []byte
	}
	type fields struct {
		challenge string
		event     nostr.Event
	}
	tests := []struct {
		name   string
		args   args
		fields fields
		expect *nostr.AuthMessage
		err    error
	}{
		{
			name: "MUST successfully unmarshal byte slice to AuthMessage",
			args: args{
				data: []byte("[\"AUTH\",\"test-challenge\"]"),
			},
			fields: fields{
				challenge: "test-challenge",
				event:     nil,
			},
			expect: &nostr.AuthMessage{nostr.BaseMessage{string(nostr.MessageTypeAuth), "test-challenge"}},
			err:    nil,
		},
		{
			name: "SHOULD return an error when unmarshaling an invalid byte slice",
			args: args{
				data: []byte("invalid_data"),
			},
			fields: fields{
				challenge: "test-challenge",
				event:     nil,
			},
			expect: &nostr.AuthMessage{nostr.BaseMessage{string(nostr.MessageTypeAuth), "test-challenge"}},
			err:    fmt.Errorf("cannot unmarshal AuthMessage: invalid byte slice"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authMessage := nostr.NewAuthMessage(tt.fields.challenge, tt.fields.event)
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

func TestAuthMessage_Validate(t *testing.T) {
	type fields struct {
		mess *nostr.AuthMessage
	}
	tests := []struct {
		name   string
		fields fields
		err    error
	}{
		{
			name: "MUST successfully validate AuthMessage",
			fields: fields{
				mess: &nostr.AuthMessage{nostr.BaseMessage{string(nostr.MessageTypeAuth), "test-challenge"}},
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authMessage := nostr.NewAuthMessage("test-challenge", nil)
			err := authMessage.Validate()
			if (err != nil && tt.err == nil) && (err == nil && tt.err != nil) && (err.Error() != tt.err.Error()) {
				t.Fatalf("expected error: %+v, got: %+v", tt.err, err)
			}
			t.Logf("got: %v", err)
		})
	}
}

func Test_NewCloseMessage(t *testing.T) {
	type args struct {
		subscriptionID string
	}
	tests := []struct {
		name   string
		args   args
		expect *nostr.CloseMessage
	}{
		{
			name: "MUST create a new CloseMessage with given subscription ID",
			args: args{
				subscriptionID: "subscription-id",
			},
			expect: &nostr.CloseMessage{nostr.BaseMessage{string(nostr.MessageTypeClose), "subscription-id"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			closeMessage := nostr.NewCloseMessage(tt.args.subscriptionID)
			if !reflect.DeepEqual(closeMessage, tt.expect) {
				t.Errorf("expected %s, got %s", closeMessage, tt.expect)
			}
			t.Logf("got %+v", closeMessage)
		})
	}
}

func TestCloseMessage_Marshal(t *testing.T) {
	type args struct {
		subscriptionID string
	}
	tests := []struct {
		name   string
		args   args
		expect []byte
		err    error
	}{
		{
			name: "MUST successfully marshal CloseMessage to byte slice",
			args: args{
				subscriptionID: "subscription-id",
			},
			expect: []byte("[\"CLOSE\",\"subscription-id\"]"),
			err:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			closeMessage := nostr.NewCloseMessage(tt.args.subscriptionID)
			data, err := closeMessage.Marshal()
			if (err != nil && tt.err == nil) && (err == nil && tt.err != nil) && (err.Error() != tt.err.Error()) {
				t.Fatalf("expected error: %s, got error: %s", tt.err, err)
				return
			}
			if !reflect.DeepEqual(data, tt.expect) {
				t.Fatalf("expected: %s, got: %s", tt.expect, data)
			}
			t.Logf("got: %+s", data)
		})
	}
}

func TestCloseMessage_Unmarshal(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		args   args
		expect *nostr.CloseMessage
		err    error
	}{
		{
			name: "MUST successfully unmarshal byte slice to CloseMessage",
			args: args{
				data: []byte("[\"CLOSE\",\"subscription-id\"]"),
			},
			expect: &nostr.CloseMessage{nostr.BaseMessage{string(nostr.MessageTypeClose), "subscription-id"}},
			err:    nil,
		},
		{
			name: "SHOULD return an error when unmarshaling an invalid byte slice",
			args: args{
				data: []byte("invalid_data"),
			},
			expect: &nostr.CloseMessage{nostr.BaseMessage{string(nostr.MessageTypeClose), "subscription-id"}},
			err:    fmt.Errorf("cannot unmarshal CloseMessage: invalid byte slice"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			closeMessage := nostr.NewCloseMessage("subscription-id")
			err := closeMessage.Unmarshal(tt.args.data)
			if (err != nil && tt.err == nil) && (err == nil && tt.err != nil) && (err.Error() != tt.err.Error()) {
				t.Fatalf("expected error: %+v, got: %+v", tt.err, err)
			}
			if !reflect.DeepEqual(closeMessage, tt.expect) {
				t.Fatalf("expected: %+v, got: %+v", tt.expect, closeMessage)
			}
			t.Logf("got: %v", closeMessage)
		})
	}
}

func TestCloseMessage_Validate(t *testing.T) {
	type fields struct {
		mess *nostr.CloseMessage
	}
	tests := []struct {
		name   string
		fields fields
		err    error
	}{
		{
			name: "MUST successfully validate CloseMessage",
			fields: fields{
				mess: &nostr.CloseMessage{nostr.BaseMessage{string(nostr.MessageTypeClose), "subscription-id"}},
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			closeMessage := nostr.NewCloseMessage("subscription-id")
			err := closeMessage.Validate()
			if (err != nil && tt.err == nil) && (err == nil && tt.err != nil) && (err.Error() != tt.err.Error()) {
				t.Fatalf("expected error: %+v, got: %+v", tt.err, err)
			}
			t.Logf("got: %v", err)
		})
	}
}

func Test_NewCountMessage(t *testing.T) {
	type args struct {
		subscriptionID string
		count          *nostr.Count
	}
	tests := []struct {
		name   string
		args   args
		expect *nostr.CountMessage
	}{
		{
			name: "MUST create a new CountMessage with given subscription ID and count",
			args: args{
				subscriptionID: "subscription-id",
				count:          &nostr.Count{Count: 10},
			},
			expect: &nostr.CountMessage{nostr.BaseMessage{string(nostr.MessageTypeCount), "subscription-id", &nostr.Count{Count: 10}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			countMessage := nostr.NewCountMessage(tt.args.subscriptionID, tt.args.count)
			if !reflect.DeepEqual(countMessage, tt.expect) {
				t.Errorf("expected %v, got %v", tt.expect, countMessage)
			}
			t.Logf("got %+v", countMessage)
		})
	}
}

func TestCountMessage_Marshal(t *testing.T) {
	type args struct {
		subscriptionID string
		count          *nostr.Count
	}
	tests := []struct {
		name   string
		args   args
		expect []byte
		err    error
	}{
		{
			name: "MUST successfully marshal CountMessage to byte slice",
			args: args{
				subscriptionID: "subscription-id",
				count:          &nostr.Count{Count: 10},
			},
			expect: []byte("[\"COUNT\",\"subscription-id\",{\"count\":10}]"),
			err:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			countMessage := nostr.NewCountMessage(tt.args.subscriptionID, tt.args.count)
			data, err := countMessage.Marshal()
			if (err != nil && tt.err == nil) && (err == nil && tt.err != nil) && (err.Error() != tt.err.Error()) {
				t.Fatalf("expected error: %s, got error: %s", tt.err, err)
				return
			}
			if !reflect.DeepEqual(data, tt.expect) {
				t.Fatalf("expected: %s, got: %s", tt.expect, data)
			}
			t.Logf("got: %+s", data)
		})
	}
}
