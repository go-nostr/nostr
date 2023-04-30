package nostr_test

import (
	"reflect"
	"testing"

	"github.com/go-nostr/nostr"
)

func Test_NewRawTag(t *testing.T) {
	tests := []struct {
		name   string
		expect nostr.Tag
	}{
		{
			name:   "MUST create a new RawTag with given type",
			expect: &nostr.RawTag{[]byte("\"type\"")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mess := nostr.NewRawTag("type")
			if !reflect.DeepEqual(mess, tt.expect) {
				t.Fatalf("expected %v, got %v", tt.expect, mess)
			}
			t.Logf("got %+v", mess)
		})
	}
}

func TestRawTag_Marshal(t *testing.T) {
	type args struct{}
	type fields struct {
		mess nostr.Tag
	}
	tests := []struct {
		name   string
		args   args
		fields fields
		expect []byte
		err    error
	}{
		{
			name: "MUST create a new RawTag with given type",
			args: args{},
			fields: fields{
				mess: nostr.NewRawTag("type"),
			},
			expect: []byte("[\"type\"]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mess, err := tt.fields.mess.Marshal()
			if err != nil && tt.err == nil {
				t.Fatalf("expected error %v, got error %v", tt.err, err)
			}
			if !reflect.DeepEqual(mess, tt.expect) {
				t.Fatalf("expected %v, got %v", tt.expect, mess)
			}
			t.Logf("got %+v", mess)
		})
	}
}

func TestRawTag_Type(t *testing.T) {
	type fields struct {
		mess nostr.Tag
	}
	tests := []struct {
		name   string
		fields fields
		expect string
		err    error
	}{
		{
			name: "MUST get type for raw message",
			fields: fields{
				mess: nostr.NewRawTag("type"),
			},
			expect: "type",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typ := tt.fields.mess.Type()
			if !reflect.DeepEqual(typ, tt.expect) {
				t.Fatalf("expected %v, got %v", tt.expect, tt.fields.mess)
			}
			t.Logf("got %+v", tt.fields.mess)
		})
	}
}

func TestRawTag_Unmarshal(t *testing.T) {
	type args struct {
		data []byte
	}
	type fields struct {
		mess nostr.Tag
	}
	tests := []struct {
		name   string
		args   args
		fields fields
		expect nostr.Tag
		err    error
	}{
		{
			name: "MUST unmarshal RawTag",
			args: args{
				data: []byte("[\"type\"]"),
			},
			fields: fields{
				mess: &nostr.RawTag{},
			},
			expect: nostr.NewRawTag("type"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fields.mess.Unmarshal(tt.args.data)
			if err != nil && tt.err == nil {
				t.Fatalf("expected error %v, got error %v", tt.err, err)
			}
			if !reflect.DeepEqual(tt.fields.mess, tt.expect) {
				t.Fatalf("expected %v, got %v", tt.expect, tt.fields.mess)
			}
			t.Logf("got %+v", tt.fields.mess)
		})
	}
}

func TestRawTag_Values(t *testing.T) {
	type fields struct {
		mess nostr.Tag
	}
	tests := []struct {
		name   string
		fields fields
		expect []any
		err    error
	}{
		{
			name: "MUST get type for raw message",
			fields: fields{
				mess: nostr.NewRawTag("type"),
			},
			expect: []any{"type"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.fields.mess.Values()
			if !reflect.DeepEqual(v, tt.expect) {
				t.Fatalf("expected %v, got %v", tt.expect, tt.fields.mess)
			}
			t.Logf("got %+v", tt.fields.mess)
		})
	}
}
