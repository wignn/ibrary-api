package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	UserID   int    `json:"userID"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.StandardClaims
}

func GenerateToken(username string, userID int, isAdmin bool) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	expirationTime := time.Now().Add(10 * time.Hour)

	claims := &Claims{
		Username: username,
		UserID:   userID,
		IsAdmin:  isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv(secretKey)))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return signedToken, nil
}
