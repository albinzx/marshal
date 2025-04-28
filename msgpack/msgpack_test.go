package msgpack

import (
	"reflect"
	"testing"

	"github.com/vmihailenco/msgpack/v5"
)

func TestMarshalMessagePack(t *testing.T) {
	type fields struct {
		Type reflect.Type
	}
	type args struct {
		value any
	}
	type Value struct {
		Number int
		Text   string
	}
	value := Value{
		Number: 1,
		Text:   "one",
	}
	bytes, _ := msgpack.Marshal(value)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "marshal struct type to msgpack",
			fields:  fields{Type: reflect.TypeOf(Value{})},
			args:    args{value: value},
			want:    bytes,
			wantErr: false},

		{name: "marshal pointer type to msgpack",
			fields:  fields{Type: reflect.TypeOf(&Value{})},
			args:    args{value: value},
			want:    bytes,
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := New(tt.fields.Type)
			got, err := m.Marshal(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnmarshalMessagePack(t *testing.T) {
	type fields struct {
		Type reflect.Type
	}
	type args struct {
		bytes []byte
	}
	type Value struct {
		Number int
		Text   string
	}
	value := Value{
		Number: 1,
		Text:   "one",
	}
	bytes, _ := msgpack.Marshal(value)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		{
			name:    "unmarshal message pack to struct type",
			fields:  fields{Type: reflect.TypeOf(Value{})},
			args:    args{bytes: bytes},
			want:    value,
			wantErr: false,
		},
		{
			name:    "unmarshal message pack to pointer type",
			fields:  fields{Type: reflect.TypeOf(&Value{})},
			args:    args{bytes: bytes},
			want:    &value,
			wantErr: false,
		},
		{
			name:    "unmarshal to invalid type",
			fields:  fields{Type: reflect.TypeOf(Value{})},
			args:    args{bytes: []byte{}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := New(tt.fields.Type)
			got, err := m.Unmarshal(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unmarshal() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		valueType reflect.Type
	}
	type Value struct {
		Number int
		Text   string
	}
	tests := []struct {
		name    string
		args    args
		wantNil bool
		wantErr bool
	}{
		{
			name:    "valid type",
			args:    args{valueType: reflect.TypeOf(Value{})},
			wantNil: false,
			wantErr: false,
		},
		{
			name:    "nil type",
			args:    args{valueType: nil},
			wantNil: true,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.valueType)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got == nil) != tt.wantNil {
				t.Errorf("New() = %v, wantNil %v", got, tt.wantNil)
			}
		})
	}
}
