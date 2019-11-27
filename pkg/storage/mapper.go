package storage

import "errors"

var (
	errNotInitialized   = errors.New("uninitialized storage")
	errNotSupported     = errors.New("unsupported storage")
	errDuplicatedRecord = errors.New("duplicated record")
	errNotFoundRecord   = errors.New("not found record")
)

// Type defines a kind of storage for storage
type Type int

const (
	// InMem is to construct im-memory storage
	InMem Type = 1
	// File is to construct local file storage
	File Type = 2
	// Database is to create a connection to storage table
	Database Type = 3
)

// Mapper provides the interface of db behavior
type Mapper interface {
	List() (map[string]string, error)
	Add(string, string) error
	Update(string, string) error
	Remove(string) error
	Clean() error
	Info(string) (string, error)
}

// Config for storage at initialization-stage
type Config struct {
	Type Type
}

// NewMapper returns a new clean storage
func NewMapper(cfg *Config) (Mapper, error) {
	switch cfg.Type {
	case InMem:
		return newMemMapper(), nil
	case File:
		return newFileMapper(), nil
	default:
		return nil, errNotSupported
	}
}

func must(m Mapper) error {
	if m == nil {
		return errNotInitialized
	}

	return nil
}
