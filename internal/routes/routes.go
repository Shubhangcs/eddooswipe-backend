package routes

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/levionstudio/eddoswipe-backend/internal/database"
	"github.com/levionstudio/eddoswipe-backend/pkg"
)

type Router struct {
	EchoRouter *echo.Echo
}

type Config struct {
	Database     *database.Database
	JWTSecretKey string
	JWTExpiry    time.Duration
}

func NewRouter(cfg Config) (*Router, error) {
	// Creating a new Echo Router
	router := echo.New()
	jwt, err := pkg.NewJWTUtils(pkg.Config{
		SecretKey: cfg.JWTSecretKey,
		Expiry:    cfg.JWTExpiry,
	})
	if err != nil {
		return nil, err
	}
	router.Validator = NewValidator()

	rtr := Router{EchoRouter: router}

	rtr.EchoRouter.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	rtr.authenticationRoutes(cfg.Database, jwt)
	rtr.usersRoutes(cfg.Database, jwt)

	return &rtr, nil
}
