package msgpack

import (
	"reflect"

	"github.com/albinzx/marshal/internal"
	"github.com/vmihailenco/msgpack/v5"
)

// New creates a new TypeMarshaller for the given value type using MessagePack encoding.
func New(valueType reflect.Type) (*internal.TypeMarshaller, error) {
	return internal.New(valueType, msgpack.Marshal, msgpack.Unmarshal)
}
