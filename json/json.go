package json

import (
	"encoding/json"
	"reflect"

	"github.com/albinzx/marshal/internal"
)

// New creates a new TypeMarshaller for the given value type using JSON encoding.
func New(valueType reflect.Type) (*internal.TypeMarshaller, error) {
	return internal.New(valueType, json.Marshal, json.Unmarshal)
}
