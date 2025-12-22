package parser

// Manager is the generic interface for any parsing implementation.
type Manager[T any] interface {
	Write(file string, data T) error
	Read(file string) (T, error)
}
