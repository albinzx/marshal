package internal

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
)

func TestTypeMarshaller_Marshal(t *testing.T) {
	type fields struct {
		valueType reflect.Type
		marshal   MarshalFunc
		unmarshal UnmarshalFunc
	}
	type args struct {
		value any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "marshal nil",
			fields: fields{
				valueType: reflect.TypeOf(""),
				marshal: func(v any) ([]byte, error) {
					return json.Marshal(v)
				},
			},
			args: args{
				value: nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "successful marshal",
			fields: fields{
				valueType: reflect.TypeOf(""),
				marshal: func(v any) ([]byte, error) {
					return json.Marshal(v)
				},
			},
			args: args{
				value: "test",
			},
			want:    []byte(`"test"`),
			wantErr: false,
		},
		{
			name: "marshal error",
			fields: fields{
				valueType: reflect.TypeOf(""),
				marshal: func(v any) ([]byte, error) {
					return nil, errors.New("marshal error")
				},
			},
			args: args{
				value: "test",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "marshal pointer type",
			fields: fields{
				valueType: reflect.TypeOf(new(string)),
				marshal: func(v any) ([]byte, error) {
					return json.Marshal(v)
				},
			},
			args: args{
				value: new(string),
			},
			want:    []byte(`""`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &TypeMarshaller{
				valueType: tt.fields.valueType,
				marshal:   tt.fields.marshal,
				unmarshal: tt.fields.unmarshal,
			}
			got, err := m.Marshal(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("TypeMarshaller.Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TypeMarshaller.Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypeMarshaller_Unmarshal(t *testing.T) {
	type fields struct {
		valueType reflect.Type
		marshal   MarshalFunc
		unmarshal UnmarshalFunc
	}
	type args struct {
		bytes []byte
	}
	str := "test"
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "unmarshal nil",
			fields: fields{
				valueType: reflect.TypeOf(""),
				unmarshal: func(data []byte, v any) error {
					return json.Unmarshal(data, v)
				},
			},
			args: args{
				bytes: nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "successful unmarshal",
			fields: fields{
				valueType: reflect.TypeOf(""),
				unmarshal: func(data []byte, v any) error {
					return json.Unmarshal(data, v)
				},
			},
			args: args{
				bytes: []byte(`"test"`),
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "unmarshal error",
			fields: fields{
				valueType: reflect.TypeOf(""),
				unmarshal: func(data []byte, v any) error {
					return errors.New("unmarshal error")
				},
			},
			args: args{
				bytes: []byte(`"test"`),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unmarshal invalid type",
			fields: fields{
				valueType: reflect.TypeOf(""),
				unmarshal: func(data []byte, v any) error {
					return json.Unmarshal(data, v)
				},
			},
			args: args{
				bytes: []byte(`{"invalid": "data"}`),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unmarshal pointer type",
			fields: fields{
				valueType: reflect.TypeOf(new(string)),
				unmarshal: func(data []byte, v any) error {
					return json.Unmarshal(data, v)
				},
			},
			args: args{
				bytes: []byte(`"test"`),
			},
			want:    &str,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &TypeMarshaller{
				valueType: tt.fields.valueType,
				marshal:   tt.fields.marshal,
				unmarshal: tt.fields.unmarshal,
			}
			got, err := m.Unmarshal(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("TypeMarshaller.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TypeMarshaller.Unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		valueType reflect.Type
		marshal   MarshalFunc
		unmarshal UnmarshalFunc
	}
	tests := []struct {
		name    string
		args    args
		wantNil bool
		wantErr bool
	}{
		{
			name: "successful creation",
			args: args{
				valueType: reflect.TypeOf(""),
				marshal: func(v any) ([]byte, error) {
					return json.Marshal(v)
				},
				unmarshal: func(data []byte, v any) error {
					return json.Unmarshal(data, v)
				},
			},
			wantNil: false,
			wantErr: false,
		},
		{
			name: "nil value type",
			args: args{
				valueType: nil,
				marshal: func(v any) ([]byte, error) {
					return json.Marshal(v)
				},
				unmarshal: func(data []byte, v any) error {
					return json.Unmarshal(data, v)
				},
			},
			wantNil: true,
			wantErr: true,
		},
		{
			name: "nil marshal function",
			args: args{
				valueType: reflect.TypeOf(""),
				marshal:   nil,
				unmarshal: func(data []byte, v any) error {
					return json.Unmarshal(data, v)
				},
			},
			wantNil: true,
			wantErr: true,
		},
		{
			name: "nil unmarshal function",
			args: args{
				valueType: reflect.TypeOf(""),
				marshal: func(v any) ([]byte, error) {
					return json.Marshal(v)
				},
				unmarshal: nil,
			},
			wantNil: true,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.valueType, tt.args.marshal, tt.args.unmarshal)
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
