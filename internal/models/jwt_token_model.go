package models

import "github.com/golang-jwt/jwt/v5"

type JWTTokenModel struct {
	AdminID string `json:"admin_id,omitempty"`
	ID      string `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	jwt.RegisteredClaims
}
