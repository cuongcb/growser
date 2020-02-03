package config

import (
	"os"

	"github.com/cuongcb/growser/pkg/storage"
)

const (
	defaultConfigDir string = ".growser"
	defaultDBFile    string = "gr.db"
)

var defaultConfigPath string

func init() {
	if defaultConfigPath, err := os.UserHomeDir(); err != nil {
		panic(err)
	}
}

// Config contains configurations for entire program
type Config struct {
	ConfigPath string
	storage.Config
}

// New returns a new configuration
func New(cfgPath string, storageType storage.Type) Config {
	path := cfgPath
	if path == "" {
		path = defaultConfigPath
	}

	return Config{ConfigPath: path, storage.Config: storage.Config{storageType}}
}
