package config

import (
	"os"
	"path"

	"github.com/cuongcb/growser/pkg/storage"
)

const (
	defaultConfigDir string = ".growser"
	defaultDBFile    string = "db"
)

var defaultConfigPath string

func init() {
	homePath, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	defaultConfigPath = homePath
}

// Config contains configurations for entire program
type Config struct {
	DBPath      string
	StorageType storage.Type
}

// New returns a new configuration
func New() *Config {
	dbPath := path.Join(defaultConfigPath, defaultConfigDir, defaultDBFile)
	return &Config{DBPath: dbPath}
}

// WithDBPath stores path to db
func (cfg *Config) WithDBPath(path string) *Config {
	cfg.DBPath = path
	return cfg
}

// WithStorage stores storage type
func (cfg *Config) WithStorage(storage storage.Type) *Config {
	cfg.StorageType = storage
	return cfg
}
