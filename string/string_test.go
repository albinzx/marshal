package string

import (
	"reflect"
	"testing"
)

func TestMarshaller_Marshal(t *testing.T) {
	type args struct {
		value any
	}
	tests := []struct {
		name    string
		m       *Marshaller
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "string",
			m:    &Marshaller{},
			args: args{
				value: "test",
			},
			want:    []byte("test"),
			wantErr: false,
		},
		{
			name: "empty string",
			m:    &Marshaller{},
			args: args{
				value: "",
			},
			want:    []byte(""),
			wantErr: false,
		},
		{
			name: "nil",
			m:    &Marshaller{},
			args: args{
				value: nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "unsupported type",
			m:    &Marshaller{},
			args: args{
				value: 123,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Marshaller{}
			got, err := m.Marshal(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshaller.Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshaller.Marshal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarshaller_Unmarshal(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		m       *Marshaller
		args    args
		want    any
		wantErr bool
	}{
		{
			name: "string",
			m:    &Marshaller{},
			args: args{
				bytes: []byte("test"),
			},
			want:    "test",
			wantErr: false,
		},
		{
			name: "empty bytes",
			m:    &Marshaller{},
			args: args{
				bytes: []byte{},
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "nil",
			m:    &Marshaller{},
			args: args{
				bytes: nil,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Marshaller{}
			got, err := m.Unmarshal(tt.args.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshaller.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Marshaller.Unmarshal() = %v, want %v", got, tt.want)
			}
		})
	}
}
