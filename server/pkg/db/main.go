package db

import (
	"shelfie-demo/internal/config"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// Db is the public interface for the database client
type Db interface {
}

// env is the internal environment for the database client
type env struct {
	logger    *zap.SugaredLogger  `validate:"required"` // Logger
	validator *validator.Validate `validate:"required"` // Validator
	client    *azcosmos.Client    `validate:"required"` // Cosmos DB client
}

func New(logger *zap.SugaredLogger, cfg *config.Config) (Db, error) {
	logger.Info("Creating new database client")

	// Create new validator
	validate := validator.New()

	// Create new Cosmos DB client
	cred, err := azcosmos.NewKeyCredential(cfg.Database.Current.Key)
	if err != nil {
		return nil, err
	}

	client, err := azcosmos.NewClientWithKey(cfg.Database.Current.Endpoint, cred, nil)
	if err != nil {
		return nil, err
	}

	// Create new environment
	e := &env{
		logger:    logger,
		validator: validate,
		client:    client,
	}

	return e, nil
}
