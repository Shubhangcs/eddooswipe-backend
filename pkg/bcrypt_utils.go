package pkg

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hashPassword, normalPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(normalPassword))
}

func GenerateHashedPassword(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password")
	}
	return string(pass), nil
}
