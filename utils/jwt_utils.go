package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// GenerateToken generates a JWT token for a user.
func GenerateToken(userID string, userRole string) (string, error) {
	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   userID,
		"userRole": userRole,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	})

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// VerifyToken verifies a JWT token and returns the claims.
func VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
