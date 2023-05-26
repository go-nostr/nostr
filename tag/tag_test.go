package tag_test

import (
	"reflect"
	"testing"

	"github.com/go-nostr/nostr/tag"
)

func Test_New(t *testing.T) {
	tests := []struct {
		name   string
		expect tag.Tag
	}{
		{
			name:   "SHOULD construct instance of tag",
			expect: tag.Tag{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tag.New()
			if !reflect.DeepEqual(got, tt.expect) {
				t.Fatalf("expected %v, got %v", tt.expect, got)
			}
			t.Logf("got %v", got)
		})
	}
}

func TestTag_Marshal(t *testing.T) {
	type fields struct {
		tag tag.Tag
	}
	tests := []struct {
		name   string
		expect []byte
		fields fields
	}{
		{
			name: "SHOULD marshal tag",
			fields: fields{
				tag: tag.Tag{"type"},
			},
			expect: []byte("[\"type\"]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := tt.fields.tag.Marshal()
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tt.expect, data) {
				t.Fatalf("expected %v, got %v", tt.expect, data)
			}
			t.Logf("got %v", data)
		})
	}
}

func TestTag_Push(t *testing.T) {
	type args struct {
		v any
	}
	type fields struct {
		tag tag.Tag
	}
	tests := []struct {
		name   string
		args   args
		fields fields
		expect tag.Tag
	}{
		{
			name: "SHOULD push new element into tag",
			args: args{
				v: "type",
			},
			fields: fields{
				tag: tag.Tag{},
			},
			expect: tag.Tag{"type"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.tag.Push(tt.args.v)
			if !reflect.DeepEqual(tt.expect, tt.fields.tag) {
				t.Fatalf("expected %v, got %v", tt.expect, tt.fields.tag)
			}
			t.Logf("got %v", tt.fields.tag)
		})
	}
}

func TestTag_Unmarshal(t *testing.T) {
	type args struct {
		data []byte
	}
	type fields struct {
		tag tag.Tag
	}
	tests := []struct {
		name   string
		args   args
		fields fields
		expect tag.Tag
	}{
		{
			name: "SHOULD unmarshal tag",
			args: args{
				data: []byte("[\"type\"]"),
			},
			fields: fields{
				tag: tag.Tag{},
			},
			expect: tag.Tag{"type"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fields.tag.Unmarshal(tt.args.data)
			if err != nil {
				t.Fatalf("expected %v, got %v", tt.expect, tt.fields.tag)
			}
			if !reflect.DeepEqual(tt.expect, tt.fields.tag) {
				t.Fatalf("expected %v, got %v", tt.expect, tt.fields.tag)
			}
			t.Logf("got %v", tt.fields.tag)
		})
	}
}

func TestTag_Values(t *testing.T) {
	type fields struct {
		tag tag.Tag
	}
	tests := []struct {
		name   string
		fields fields
		expect []any
	}{
		{
			name: "SHOULD get tag values",
			fields: fields{
				tag: tag.Tag{"type"},
			},
			expect: []any{"type"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vals := tt.fields.tag.Values()
			if !reflect.DeepEqual(tt.expect, vals) {
				t.Fatalf("expected %v, got %v", tt.expect, vals)
			}
			t.Logf("got %v", vals)
		})
	}
}
