package msgpack

import (
	"log"
	"reflect"

	"github.com/vmihailenco/msgpack/v5"
)

// MessagePackMarshaller marshal or unmarshal value in message pack format.
type MessagePackMarshaller struct {
	Type reflect.Type
}

// Marshal marshals value type to message pack.
func (m MessagePackMarshaller) Marshal(value interface{}) ([]byte, error) {
	return msgpack.Marshal(value)
}

// Unmarshal unmarshals message pack to value type.
func (m MessagePackMarshaller) Unmarshal(bytes []byte) (interface{}, error) {
	var value reflect.Value

	if m.Type.Kind() != reflect.Ptr {
		value = reflect.New(m.Type)
	} else {
		value = reflect.New(m.Type.Elem())
	}

	if err := msgpack.Unmarshal(bytes, value.Interface()); err != nil {
		log.Printf("error while unmarshalling data: %v", err)
		return nil, err
	}

	if m.Type.Kind() != reflect.Ptr {
		value = reflect.Indirect(value)
	}

	return value.Interface(), nil
}
