package utils

import (
	"fmt"
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

type CustomClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Exp      int    `json:"exp"`
	jwt.RegisteredClaims
}

func ParseToken(tokenString string) (interface{}, error) {
	secret := os.Getenv("JWT_SECRET")

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("could not parse claims")
	}

	return claims, nil
}
