package marshal

// Marshaller marshals or unmarshalls value to or from byte array.
type Marshaller interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte) (interface{}, error)
}
