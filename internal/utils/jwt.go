package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(payload jwt.Claims) *string {
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil
	}
	return &signed
}
