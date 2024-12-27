package api

import (
	"net/http"
	"shelfie-demo/internal/config"
	"shelfie-demo/pkg/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// _ "github.com/byron-ojua/starter-project/internal/api/docs"

//go:generate swag init --parseDependency --parseInternal

// @title Shelfie
// @version 1.0
// @description This is the API for the Shelfie application.
// @termsOfService TBD
//
// @contact.name Byron Ojua-Nice
// @contact.url http://firstlaunch.dev
// @contact.email byronojua@firstlaunch.dev
//
// @license.name TBD
// @license.url TBD
//
// @host localhost:8080
// @BasePath /
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Use the format: `Bearer <your_token>` (Bearer must be added before the token)
//
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

const API_VERSION = "1.0"

// error declarations
const (
	//the error returned to the user when bad user token being provided
	ErrUnauthorized = "bad user token"

	//the error returned to the user when no token is provided
	ErrNoToke = "no authentication token provided"

	//the error returned to the user when a user is not authorized to perform an action
	ErrForbidden = "forbidden"

	//the error returned when no resource ID is provided for a request pertaining to a specific resource
	ErrorNoResourceID = "no resource id provided"

	//the error returned to the user when the requested record in out of date
	ErrCalibrationOutOfDate = "out of date record"

	//error signifying that a feature is not supported for a specific product
	ErrUnsupportedFeature = "unsupported product feature"

	//err signifying that a vehicle does not have an attached device
	ErrNoVehicleDevice = "no device attached to vehicle"
)

// ErrorsResponse represents the structure for API error responses.
type ErrorsResponse struct {
	Errors []string `json:"errors"` // Array of error messages
}

// Api is the interface for the API
type Api interface {
	RunLocal() error // Run the API locally
}

// env is the environment for the API
type env struct {
	logger    *zap.SugaredLogger  `validate:"required"` // Logger
	validator *validator.Validate `validate:"required"` // Validator
	config    *config.Config      `validate:"required"` // Config
	db        db.Db               `validate:"required"` // Database
	api       *gin.Engine         `validate:"required"` // API
}

// RunLocal runs the API locally
func (e *env) RunLocal() error {
	return http.ListenAndServe("localhost:8080", e.api)
}

// New creates a new API instance
func New(logger *zap.SugaredLogger, cfg *config.Config, dbInstance db.Db) (Api, error) {
	// Create a new validator
	validator := validator.New()

	e := &env{
		validator: validator,
		logger:    logger,
		config:    cfg,
		db:        dbInstance,
	}

	r := gin.Default()

	//CORS setup
	r.Use(cors.New(cors.Config{
		AllowHeaders:    []string{"Authorization", "Content-Type"},
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	}))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup the routes. Make inline return with message "hello from server 2"
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello from server 2",
		})
	})

	e.api = r

	return e, nil
}
