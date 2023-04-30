package nostr_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/go-nostr/nostr"
)

func Test_NewRawEvent(t *testing.T) {
	tests := []struct {
		name   string
		expect nostr.Event
	}{
		{
			name:   "MUST create new raw event",
			expect: &nostr.RawEvent{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := nostr.NewRawEvent()
			if !reflect.DeepEqual(tt.expect, event) {
				t.Fatalf("expected %v, got %v", tt.expect, event)
			}
			t.Logf("got: %v", event)
		})
	}
}

func TestRawEvent_Content(t *testing.T) {
	type fields struct {
		event nostr.Event
	}
	tests := []struct {
		name   string
		fields fields
		expect []byte
	}{
		{
			name: "MUST get empty content for raw event with no content",
			fields: fields{
				event: nostr.NewRawEvent(),
			},
			expect: []byte(""),
		},
		{
			name: "MUST get content for raw event with content",
			fields: fields{
				event: &nostr.RawEvent{
					"content": json.RawMessage(`"Hello, World!"`),
				},
			},
			expect: []byte("Hello, World!"),
		},
		{
			name: "MUST return empty content for malformed JSON",
			fields: fields{
				event: &nostr.RawEvent{
					"content": json.RawMessage(`"unclosed_string`),
				},
			},
			expect: []byte(""),
		},
		{
			name: "MUST return empty content for raw event without content key",
			fields: fields{
				event: &nostr.RawEvent{
					"other_key": json.RawMessage(`"Hello, World!"`),
				},
			},
			expect: []byte(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			content := tt.fields.event.Content()
			if !reflect.DeepEqual(tt.expect, content) {
				t.Fatalf("expected %v, got %v", tt.expect, content)
			}
			t.Logf("got: %v", content)
		})
	}
}

func TestRawEvent_Get(t *testing.T) {
	type args struct {
		key string
	}
	type fields struct {
		event nostr.Event
	}
	tests := []struct {
		name   string
		args   args
		fields fields
		expect any
	}{
		{
			name: "MUST get nil value for unset key",
			args: args{
				key: "key",
			},
			fields: fields{
				event: nostr.NewRawEvent(),
			},
			expect: nil,
		},
		{
			name: "MUST get correct value for existing key",
			args: args{
				key: "key",
			},
			fields: fields{
				event: &nostr.RawEvent{
					"key": json.RawMessage(`"value"`),
				},
			},
			expect: "value",
		},
		{
			name: "MUST get nil value for malformed JSON",
			args: args{
				key: "key",
			},
			fields: fields{
				event: &nostr.RawEvent{
					"key": json.RawMessage(`"unclosed_string`),
				},
			},
			expect: nil,
		},
		{
			name: "MUST get correct value for numeric key",
			args: args{
				key: "number",
			},
			fields: fields{
				event: &nostr.RawEvent{
					"number": json.RawMessage(`42`),
				},
			},
			expect: float64(42),
		},
		{
			name: "MUST get correct value for boolean key",
			args: args{
				key: "bool",
			},
			fields: fields{
				event: &nostr.RawEvent{
					"bool": json.RawMessage(`true`),
				},
			},
			expect: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val := tt.fields.event.Get(tt.args.key)
			if !reflect.DeepEqual(tt.expect, val) {
				t.Fatalf("expected %v, got %v", tt.expect, val)
			}
			t.Logf("got: %v", val)
		})
	}
}

func TestRawEvent_ID(t *testing.T) {
	events := map[int]nostr.Event{
		1: nostr.NewRawEvent(),
		2: nostr.NewRawEvent(),
		3: nostr.NewRawEvent(),
	}
	events[1].Set("id", "asdf-1234")
	events[3].Set("id", struct{}{})

	type fields struct {
		event nostr.Event
	}
	tests := []struct {
		name   string
		fields fields
		expect []byte
	}{
		{
			name: "MUST get ID value",
			fields: fields{
				event: events[1],
			},
			expect: []byte("asdf-1234"),
		},
		{
			name: "MUST return empty ID for event without ID key",
			fields: fields{
				event: events[2],
			},
			expect: []byte(""),
		},
		{
			name: "MUST return empty ID for event with malformed ID value",
			fields: fields{
				event: events[3],
			},
			expect: []byte(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := tt.fields.event.ID()
			if !reflect.DeepEqual(tt.expect, id) {
				t.Fatalf("expected %v, got %v", tt.expect, id)
			}
			t.Logf("got: %v", id)
		})
	}
}

func TestRawEvent_Keys(t *testing.T) {
	events := map[int]nostr.Event{
		1: nostr.NewRawEvent(),
		2: nostr.NewRawEvent(),
		3: nostr.NewRawEvent(),
	}
	events[1].Set("id", "asdf-1234")
	events[1].Set("content", "example content")
	events[2].Set("timestamp", "1631234567")
	events[3].Set("id", "qwer-5678")
	events[3].Set("content", "another example content")

	type fields struct {
		event nostr.Event
	}
	tests := []struct {
		name   string
		fields fields
		expect []string
	}{
		{
			name: "MUST get keys for event with multiple keys",
			fields: fields{
				event: events[1],
			},
			expect: []string{"id", "content"},
		},
		{
			name: "MUST get keys for event with single key",
			fields: fields{
				event: events[2],
			},
			expect: []string{"timestamp"},
		},
		{
			name: "MUST get keys for event with multiple keys",
			fields: fields{
				event: events[3],
			},
			expect: []string{"id", "content"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keys := tt.fields.event.Keys()
			if !equalStringSlicesUnordered(tt.expect, keys) {
				t.Fatalf("expected %v, got %v", tt.expect, keys)
			}
			t.Logf("got: %v", keys)
		})
	}
}

func equalStringSlicesUnordered(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[string]int)
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		m[v]--
		if m[v] < 0 {
			return false
		}
	}
	return true
}

func TestRawEvent_Kind(t *testing.T) {
	events := map[int]nostr.Event{
		1: nostr.NewRawEvent(),
		2: nostr.NewRawEvent(),
		3: nostr.NewRawEvent(),
	}
	events[1].Set("kind", 1)
	events[3].Set("kind", "invalid")

	type fields struct {
		event nostr.Event
	}
	tests := []struct {
		name   string
		fields fields
		expect int
	}{
		{
			name: "MUST get kind value",
			fields: fields{
				event: events[1],
			},
			expect: 1,
		},
		{
			name: "MUST return -1 for event without kind key",
			fields: fields{
				event: events[2],
			},
			expect: -1,
		},
		{
			name: "MUST return -1 for event with malformed kind value",
			fields: fields{
				event: events[3],
			},
			expect: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kind := tt.fields.event.Kind()
			if tt.expect != kind {
				t.Fatalf("expected %v, got %v", tt.expect, kind)
			}
			t.Logf("got: %v", kind)
		})
	}
}

func TestRawEvent_PublicKey(t *testing.T) {
	events := map[int]nostr.Event{
		1: nostr.NewRawEvent(),
		2: nostr.NewRawEvent(),
		3: nostr.NewRawEvent(),
	}
	events[1].Set("pubkey", "abcd1234")
	events[3].Set("pubkey", struct{}{})

	type fields struct {
		event nostr.Event
	}
	tests := []struct {
		name   string
		fields fields
		expect []byte
	}{
		{
			name: "MUST get PublicKey value",
			fields: fields{
				event: events[1],
			},
			expect: []byte("abcd1234"),
		},
		{
			name: "MUST return empty PublicKey for event without pubkey key",
			fields: fields{
				event: events[2],
			},
			expect: []byte(""),
		},
		{
			name: "MUST return empty PublicKey for event with malformed pubkey value",
			fields: fields{
				event: events[3],
			},
			expect: []byte(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pubKey := tt.fields.event.PublicKey()
			if !reflect.DeepEqual(tt.expect, pubKey) {
				t.Fatalf("expected %v, got %v", tt.expect, pubKey)
			}
			t.Logf("got: %v", pubKey)
		})
	}
}

func TestRawEvent_Signature(t *testing.T) {
	events := map[int]nostr.Event{
		1: nostr.NewRawEvent(),
		2: nostr.NewRawEvent(),
		3: nostr.NewRawEvent(),
	}
	events[1].Set("sig", "abcd-5678")
	events[3].Set("sig", struct{}{})

	type fields struct {
		event nostr.Event
	}
	tests := []struct {
		name   string
		fields fields
		expect []byte
	}{
		{
			name: "MUST get signature value",
			fields: fields{
				event: events[1],
			},
			expect: []byte("abcd-5678"),
		},
		{
			name: "MUST return empty signature for event without sig key",
			fields: fields{
				event: events[2],
			},
			expect: []byte(""),
		},
		{
			name: "MUST return empty signature for event with malformed sig value",
			fields: fields{
				event: events[3],
			},
			expect: []byte(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sig := tt.fields.event.Signature()
			if !reflect.DeepEqual(tt.expect, sig) {
				t.Fatalf("expected %v, got %v", tt.expect, sig)
			}
			t.Logf("got: %v", sig)
		})
	}
}

// func TestRawEvent_Tags(t *testing.T) {
// 	events := map[int]nostr.Event{
// 		1: nostr.NewRawEvent(),
// 		2: nostr.NewRawEvent(),
// 		3: nostr.NewRawEvent(),
// 	}

// 	events[1].Set("tags", []nostr.Tag{
// 		nostr.NewRawTag("tag1"),
// 		nostr.NewRawTag("tag2"),
// 	})
// 	events[3].Set("tags", struct{}{})

// 	type fields struct {
// 		event nostr.Event
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		expect []nostr.Tag
// 	}{
// 		{
// 			name: "MUST get tags value",
// 			fields: fields{
// 				event: events[1],
// 			},
// 			expect: []nostr.Tag{&nostr.RawTag{[]byte("tag1")}, &nostr.RawTag{[]byte("tag2")}},
// 		},
// 		{
// 			name: "MUST return empty tags for event without tags key",
// 			fields: fields{
// 				event: events[2],
// 			},
// 			expect: []nostr.Tag{},
// 		},
// 		{
// 			name: "MUST return empty tags for event with malformed tags value",
// 			fields: fields{
// 				event: events[3],
// 			},
// 			expect: []nostr.Tag{},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tags := tt.fields.event.Tags()
// 			if !reflect.DeepEqual(tt.expect, tags) {
// 				t.Fatalf("expected %v, got %v", tt.expect, tags)
// 			}
// 			t.Logf("got: %v", tags)
// 		})
// 	}
// }

// For Go types:
// type RawEvent map[string]json.RawMessage
//
// And methods:
//
//	func (e *RawEvent) Marshal() ([]byte, error) {
//		return json.Marshal(e)
//	}
//
//	func (e *RawEvent) Unmarshal(data []byte) error {
//		return json.Unmarshal(data, &e)
//	}
//
// Follow the instructions as '// PROMPT: <INSTRUCTION>'
//
// PROMPT: Add test functions for both methods in a similar code style to the follow
func TestRawEvent_Tags(t *testing.T) {
	events := map[int]nostr.Event{
		1: nostr.NewRawEvent(),
		2: nostr.NewRawEvent(),
		3: nostr.NewRawEvent(),
	}

	events[1].Set("tags", []nostr.Tag{nostr.NewRawTag("tag1"), nostr.NewRawTag("tag2")})
	events[3].Set("tags", struct{}{})

	type fields struct {
		event nostr.Event
	}
	tests := []struct {
		name   string
		fields fields
		expect []nostr.Tag
	}{
		{
			name: "MUST get tags value",
			fields: fields{
				event: events[1],
			},
			expect: []nostr.Tag{nostr.NewRawTag("tag1"), nostr.NewRawTag("tag2")},
		},
		{
			name: "MUST return empty tags for event without tags key",
			fields: fields{
				event: events[2],
			},
			expect: []nostr.Tag{},
		},
		{
			name: "MUST return empty tags for event with malformed tags value",
			fields: fields{
				event: events[3],
			},
			expect: []nostr.Tag{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tags := tt.fields.event.Tags()
			if !reflect.DeepEqual(tt.expect, tags) {
				t.Fatalf("expected %v, got %v", tt.expect, tags)
			}
			t.Logf("got: %v", tags)
		})
	}
}

func TestRawEvent_Marshal(t *testing.T) {
	event := nostr.NewRawEvent()
	event.Set("id", "test-id")
	event.Set("tags", []nostr.Tag{nostr.NewRawTag("tag1"), nostr.NewRawTag("tag2")})

	type fields struct {
		event nostr.Event
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "MUST marshal event without error",
			fields: fields{
				event: event,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.fields.event.Marshal()
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRawEvent_Unmarshal(t *testing.T) {
	type fields struct {
		event nostr.Event
	}
	tests := []struct {
		name    string
		fields  fields
		data    []byte
		wantErr bool
	}{
		{
			name: "MUST unmarshal event without error",
			fields: fields{
				event: nostr.NewRawEvent(),
			},
			data:    []byte(`{"id": "test-id", "tags": [["tag1"], ["tag2"]]}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fields.event.Unmarshal(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRawEvent_Values(t *testing.T) {
	events := map[int]nostr.Event{
		1: nostr.NewRawEvent(),
		2: nostr.NewRawEvent(),
		3: nostr.NewRawEvent(),
	}

	events[1].Set("key1", "value1")
	events[1].Set("key2", "value2")
	events[2].Set("key3", "value3")
	events[3].Set("key4", struct{}{})

	type fields struct {
		event nostr.Event
	}
	tests := []struct {
		name   string
		fields fields
		expect []any
	}{
		{
			name: "MUST get values",
			fields: fields{
				event: events[1],
			},
			expect: []any{json.RawMessage(`"value1"`), json.RawMessage(`"value2"`)},
		},
		{
			name: "MUST get a single value",
			fields: fields{
				event: events[2],
			},
			expect: []any{json.RawMessage(`"value3"`)},
		},
		{
			name: "MUST return empty values for event with malformed value",
			fields: fields{
				event: events[3],
			},
			expect: []any{json.RawMessage("{}")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values := tt.fields.event.Values()
			if !reflect.DeepEqual(tt.expect, values) {
				t.Fatalf("expected %v, got %v", tt.expect, values)
			}
			t.Logf("got: %v", values)
		})
	}
}
