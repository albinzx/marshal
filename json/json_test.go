package json

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestJSONMarshaller_Marshal(t *testing.T) {
	type fields struct {
		Type reflect.Type
	}
	type args struct {
		value interface{}
	}
	type Data struct {
		Number int
		Name   string
	}
	data := Data{
		Number: 1,
		Name:   "one",
	}

	bytes, _ := json.Marshal(data)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "marshal to json",
			fields:  fields{Type: reflect.TypeOf(Data{})},
			args:    args{value: data},
			want:    bytes,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jm := JSONMarshaller{
				Type: tt.fields.Type,
			}
			got, err := jm.Marshal(tt.args.value)
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

func TestJSONMarshaller_Unmarshal(t *testing.T) {
	type fields struct {
		Type reflect.Type
	}
	type args struct {
		bytes []byte
	}
	type Data struct {
		Number int
		Name   string
	}
	data := Data{
		Number: 1,
		Name:   "one",
	}
	bytes, _ := json.Marshal(data)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:    "unmarshal struct type",
			fields:  fields{Type: reflect.TypeOf(Data{})},
			args:    args{bytes: bytes},
			want:    data,
			wantErr: false,
		},
		{
			name:    "unmarshal pointer type",
			fields:  fields{Type: reflect.TypeOf(&Data{})},
			args:    args{bytes: bytes},
			want:    &data,
			wantErr: false,
		},
		{
			name:    "unmarshal invalid type",
			fields:  fields{Type: reflect.TypeOf(Data{})},
			args:    args{bytes: []byte{}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jm := JSONMarshaller{
				Type: tt.fields.Type,
			}
			got, err := jm.Unmarshal(tt.args.bytes)
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
