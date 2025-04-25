package msgpack

import (
	"reflect"
	"testing"

	"github.com/vmihailenco/msgpack/v5"
)

func TestMessagePackMarshaller_Marshal(t *testing.T) {
	type fields struct {
		ValueType reflect.Type
	}
	type args struct {
		value interface{}
	}
	type Data struct {
		Num  int
		Name string
	}
	data := Data{
		Num:  1,
		Name: "one",
	}
	bytes, _ := msgpack.Marshal(data)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "marshal",
			fields:  fields{ValueType: reflect.TypeOf(Data{})},
			args:    args{value: data},
			want:    bytes,
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MessagePackMarshaller{
				ValueType: tt.fields.ValueType,
			}
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

func TestMessagePackMarshaller_Unmarshal(t *testing.T) {
	type fields struct {
		ValueType reflect.Type
	}
	type args struct {
		bytes []byte
	}
	type Data struct {
		Num  int
		Name string
	}
	data := Data{
		Num:  1,
		Name: "one",
	}
	bytes, _ := msgpack.Marshal(data)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:    "unmarshal",
			fields:  fields{ValueType: reflect.TypeOf(Data{})},
			args:    args{bytes: bytes},
			want:    data,
			wantErr: false,
		},
		{
			name:    "unmarshal-pointer",
			fields:  fields{ValueType: reflect.TypeOf(&Data{})},
			args:    args{bytes: bytes},
			want:    &data,
			wantErr: false,
		},
		{
			name:    "unmarshal-invalid",
			fields:  fields{ValueType: reflect.TypeOf(Data{})},
			args:    args{bytes: []byte{}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MessagePackMarshaller{
				ValueType: tt.fields.ValueType,
			}
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
