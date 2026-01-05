package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Database DatabaseConfig
	JWT      JWTConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	DatabaseURL string
}

type JWTConfig struct {
	JWTSecretKey          string
	JWTSecretKeyPaySprint string
	JWTExpiry             time.Duration
}

type ServerConfig struct {
	ServerENV  string
	ServerPort string
}

func Load() (*Config, error) {
	// Loading the .env variables
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load env: %w", err)
	}
	// Returning all the configurations
	return &Config{
		Database: DatabaseConfig{
			DatabaseURL: os.Getenv("DATABASE_URL"),
		},
		JWT: JWTConfig{
			JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
			JWTSecretKeyPaySprint: os.Getenv("PAY_SPRINT_JWT_KEY"),
			JWTExpiry:    24 * time.Hour,
		},
		Server: ServerConfig{
			ServerENV:  os.Getenv("SERVER_ENVIRONMENT"),
			ServerPort: os.Getenv("SERVER_PORT"),
		},
	}, nil
}
