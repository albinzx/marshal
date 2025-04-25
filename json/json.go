package json

import (
	"encoding/json"
	"log"
	"reflect"
)

// JSONMarshaller marshal or unmarshal value in json format.
type JSONMarshaller struct {
	ValueType reflect.Type
}

// Marshal marshals value type to json.
func (jm JSONMarshaller) Marshal(value interface{}) ([]byte, error) {
	return json.Marshal(value)
}

// Unmarshal unmarshals json to value type.
func (jm JSONMarshaller) Unmarshal(bytes []byte) (interface{}, error) {
	var value reflect.Value

	if jm.ValueType.Kind() != reflect.Ptr {
		value = reflect.New(jm.ValueType)
	} else {
		value = reflect.New(jm.ValueType.Elem())
	}

	if err := json.Unmarshal(bytes, value.Interface()); err != nil {
		log.Printf("error while unmarshalling data: %v", err)
		return nil, err
	}

	if jm.ValueType.Kind() != reflect.Ptr {
		value = reflect.Indirect(value)
	}

	return value.Interface(), nil
}
