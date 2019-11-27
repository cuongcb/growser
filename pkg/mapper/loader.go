package mapper

import "errors"

var (
	errNotInitialized   = errors.New("uninitialized mapper")
	errNotSupported     = errors.New("unsupported mapper")
	errDuplicatedRecord = errors.New("duplicated record")
	errNotFoundRecord   = errors.New("not found record")
)

// Type defines a kind of storage for mapper
type Type int

const (
	// InMem is to construct im-memory mapper
	InMem Type = 1
	// File is to construct local file mapper
	File Type = 2
	// Database is to create a connection to mapper table
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

// Config for mapper at initialization-stage
type Config struct {
	Type Type
}

// New returns a new clean mapper
func New(cfg *Config) (Mapper, error) {
	switch cfg.Type {
	case InMem:
		return newMemLoader(), nil
	default:
		return nil, errNotSupported
	}
}

func must(l Mapper) error {
	if l == nil {
		return errNotInitialized
	}

	return nil
}
