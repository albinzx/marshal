package json

import (
	"encoding/json"
	"log"
	"reflect"
)

// JSONMarshaller marshal or unmarshal value in json format.
type JSONMarshaller struct {
	Type reflect.Type
}

// Marshal marshals value type to json.
func (m JSONMarshaller) Marshal(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

// Unmarshal unmarshals json to value type.
func (m JSONMarshaller) Unmarshal(bytes []byte) (interface{}, error) {
	var value reflect.Value

	if m.Type.Kind() != reflect.Ptr {
		value = reflect.New(m.Type)
	} else {
		value = reflect.New(m.Type.Elem())
	}

	if err := json.Unmarshal(bytes, value.Interface()); err != nil {
		log.Printf("error while unmarshalling data: %v", err)
		return nil, err
	}

	if m.Type.Kind() != reflect.Ptr {
		value = reflect.Indirect(value)
	}

	return value.Interface(), nil
}
