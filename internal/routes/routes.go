package routes

import (
	"os"
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
	Database              *database.Database
	JWTSecretKey          string
	JWTExpiry             time.Duration
	JWTSecretKeyPaySprint string
}

func NewRouter(cfg Config) (*Router, error) {
	// Creating a new Echo Router
	router := echo.New()
	jwt, err := pkg.NewJWTUtils(pkg.Config{
		SecretKey:          cfg.JWTSecretKey,
		Expiry:             cfg.JWTExpiry,
		PaySprintSecretKey: cfg.JWTSecretKeyPaySprint,
	})
	if err != nil {
		return nil, err
	}
	router.Validator = NewValidator()
	paysprintURL := os.Getenv("PAY_SPRINT_URL")

	rtr := Router{EchoRouter: router}

	rtr.EchoRouter.Use(middleware.CORS())
	rtr.authenticationRoutes(cfg.Database, jwt)
	rtr.usersRoutes(cfg.Database, jwt)
	rtr.walletRoutes(cfg.Database, jwt)
	rtr.fundRoutes(cfg.Database, jwt)
	rtr.bankRouter(cfg.Database, jwt)
	rtr.commisionRoutes(cfg.Database, jwt)
	rtr.dmtRouter(cfg.Database, jwt, paysprintURL)

	return &rtr, nil
}
