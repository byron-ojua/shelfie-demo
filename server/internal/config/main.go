package config

import (
	_ "embed"
	"encoding/json"
	"os"

	"go.uber.org/zap"
)

//go:embed config.json
var configJSON []byte

const (
	// MAX_PAGE_SIZE is the maximum page size for calls to the database
	MAX_PAGE_SIZE = 500

	PRODUCTION_ENV = "PRODUCTION"
)

// Config is the configuration for the application
type Config struct {
	MaxPageSize int      `json:"max_page_size"`
	Database    Database `json:"cosmos"`
}

// Holds info for the Azure Cosmos DB
type Database struct {
	Current     DatabaseParams
	Production  DatabaseParams `json:"production"`
	Development DatabaseParams `json:"development"`
}

type DatabaseParams struct {
	Endpoint string `json:"endpoint"`
	Key      string `json:"key"`
}

// New creates a new Config instance
func New(logger *zap.SugaredLogger) (*Config, error) {
	logger.Info("Setting up configuration")

	var c Config

	err := json.Unmarshal(configJSON, &c)
	if err != nil {
		logger.Errorf("failed to unmarshal config: %v", err)
		return nil, err
	}

	c.MaxPageSize = MAX_PAGE_SIZE

	// Get env variable APP_ENV to determine which database to use
	// Default to development
	env := os.Getenv("APP_ENV")

	if env == PRODUCTION_ENV {
		logger.Debug("Running with production configuration")
		c.Database.Current = c.Database.Production
	} else {
		logger.Debug("Running with development configuration")
		c.Database.Current = c.Database.Development
	}

	return &c, nil
}
