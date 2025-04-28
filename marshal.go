package marshal

// Marshaller marshals or unmarshalls value to or from byte array.
type Marshaller interface {
	Marshal(any) ([]byte, error)
	Unmarshal([]byte) (any, error)
}
