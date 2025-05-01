package config

import (
	"os"
	"path/filepath"
)

// Config holds the application configuration.
type Config struct {
	DataFilePath string // Path to the JSON data file
	// Add other configuration fields here later (e.g., API keys)
}

// DefaultDataDir is the default directory name for storing application data.
const DefaultDataDir = ".smarttask"

// DefaultDataFile is the default filename for the task data.
const DefaultDataFile = "tasks.json"

// LoadConfig loads the application configuration.
// For now, it just returns default values, but can be extended
// to load from a file or environment variables.
func LoadConfig() (*Config, error) {
	// Determine the user's home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err // Cannot proceed without home directory
	}

	// Construct the default data file path
	dataDir := filepath.Join(homeDir, DefaultDataDir)
	dataFilePath := filepath.Join(dataDir, DefaultDataFile)

	// Ensure the data directory exists
	if err := os.MkdirAll(dataDir, 0750); err != nil { // Use appropriate permissions
		return nil, err
	}

	// Create the default config
	cfg := &Config{
		DataFilePath: dataFilePath,
	}

	// TODO: Implement loading from a config file (e.g., YAML/JSON in dataDir)
	// TODO: Implement loading overrides from environment variables

	return cfg, nil
}
