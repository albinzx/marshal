package string

import "fmt"

type Marshaller struct {
}

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

func (m *Marshaller) Unmarshal(bytes []byte) (any, error) {
	if bytes == nil {
		return nil, nil
	}

	return string(bytes), nil
}
