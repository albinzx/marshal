package internal

import (
	"errors"
	"log"
	"reflect"
)

// MarshalFunc is function to marshal value to byte array.
type MarshalFunc func(any) ([]byte, error)

// UnmarshalFunc is function to unmarshal byte array to value.
type UnmarshalFunc func([]byte, any) error

// TypeMarshaller is marsheller for specific type.
type TypeMarshaller struct {
	valueType reflect.Type
	marshal   MarshalFunc
	unmarshal UnmarshalFunc
}

// Marshal marshals value to byte array.
func (m *TypeMarshaller) Marshal(value any) ([]byte, error) {
	return m.marshal(value)
}

// Unmarshal unmarshals byte array to value.
func (m *TypeMarshaller) Unmarshal(bytes []byte) (any, error) {
	var value reflect.Value

	isNotPtr := m.valueType.Kind() != reflect.Ptr

	if isNotPtr {
		value = reflect.New(m.valueType)
	} else {
		value = reflect.New(m.valueType.Elem())
	}

	if err := m.unmarshal(bytes, value.Interface()); err != nil {
		log.Printf("error while unmarshalling data: %v", err)
		return nil, err
	}

	if isNotPtr {
		value = reflect.Indirect(value)
	}

	return value.Interface(), nil
}

// New creates a new TypeMarshaller for the given value type.
// It takes a value type, a marshal function, and an unmarshal function as parameters.
// It returns a pointer to the TypeMarshaller and an error if any of the parameters are nil.
// The function ensures that the TypeMarshaller is properly initialized and ready for use.
func New(valueType reflect.Type, marshal MarshalFunc, unmarshal UnmarshalFunc) (*TypeMarshaller, error) {
	// Check if all the parameters are not nil
	if valueType == nil || marshal == nil || unmarshal == nil {
		return nil, errors.New("type, marshal function, unmarshal function cannot be nil")
	}

	return &TypeMarshaller{
		valueType: valueType,
		marshal:   marshal,
		unmarshal: unmarshal,
	}, nil
}
