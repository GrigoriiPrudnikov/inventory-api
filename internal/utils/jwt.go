package utils

import (
	"fmt"
	"os"
	"time"

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
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Exp      uint   `json:"exp"`
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

	if claims.Exp < uint(time.Now().Unix()) {
		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
}
