package models

import "github.com/golang-jwt/jwt/v5"

type JWTTokenModel struct {
	ID   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
	jwt.RegisteredClaims
}
