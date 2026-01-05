package pkg

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/levionstudio/eddoswipe-backend/internal/models"
)

type JWTUtils struct {
	secretKey          string
	duration           time.Duration
	paySprintSecretKey string
}

type Config struct {
	SecretKey          string
	Expiry             time.Duration
	PaySprintSecretKey string
}

func NewJWTUtils(cfg Config) (*JWTUtils, error) {
	if cfg.SecretKey == "" || cfg.Expiry == 0 || cfg.PaySprintSecretKey == "" {
		return nil, fmt.Errorf("failed to create jwt utils fields are empty")
	}
	return &JWTUtils{
		secretKey:          cfg.SecretKey,
		duration:           cfg.Expiry,
		paySprintSecretKey: cfg.PaySprintSecretKey,
	}, nil
}

func (ju *JWTUtils) GenerateTokenForPaysprint() (string, error) {
	// Payload (same as PHP associative array)
	claims := jwt.MapClaims{
		"timestamp": time.Now().Unix(),
		"partnerId": "PS002746",
		"reqid":     uuid.NewString(),
	}

	// Create token with HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token
	tokenString, err := token.SignedString([]byte(ju.paySprintSecretKey))
	if err != nil {
		return "", fmt.Errorf("failed to generate token")
	}

	return tokenString, nil
}

func (ju *JWTUtils) GenerateToken(data models.JWTTokenModel) (string, error) {
	now := time.Now()

	data.RegisteredClaims = jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(ju.duration)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	signedToken, err := token.SignedString([]byte(ju.secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (ju *JWTUtils) ValidateToken(tokenString string) (*models.JWTTokenModel, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&models.JWTTokenModel{},
		func(token *jwt.Token) (any, error) {
			// Ensure signing method is HMAC
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(ju.secretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.JWTTokenModel)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
