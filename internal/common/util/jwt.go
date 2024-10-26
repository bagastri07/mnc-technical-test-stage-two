package util

import (
	"github.com/bagastri07/gin-boilerplate-service/internal/common/constant"
	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken generates a JWT token with custom claims
func GenerateToken[T jwt.Claims](secretKey []byte, claims T) (*string, error) {
	// Create a new token object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

// ValidateJWT validates a JWT token and returns the claims if valid
func ValidateJWT[T jwt.Claims](tokenString string, secretKey []byte, claims T) (T, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return claims, err
	}

	if !token.Valid {
		return claims, constant.ErrUnauthorized
	}

	if c, ok := token.Claims.(T); ok {
		claims = c
	}

	return claims, nil
}
