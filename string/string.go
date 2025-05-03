package string

import "fmt"

// Marshaller is a simple marshaller for string values.
type Marshaller struct {
}

// Marshal marshals a string value into a byte slice.
// It returns nil if the value is nil.
// It returns an error if the value is not a string.
// It returns the byte slice representation of the string.
func (m *Marshaller) Marshal(value any) ([]byte, error) {
	if value == nil {
		return nil, nil
	}

	switch v := value.(type) {
	case string:
		return []byte(v), nil
	default:
		return nil, fmt.Errorf("unsupported type: %T", value)
	}
}

// Unmarshal unmarshals a byte slice into a string value.
// It returns nil if the byte slice is nil.
func (m *Marshaller) Unmarshal(bytes []byte) (any, error) {
	if bytes == nil {
		return nil, nil
	}

	return string(bytes), nil
}
