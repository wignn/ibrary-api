package utils

import (
	"log"
	"os"
	"time"
	"fmt"
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
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("error signing token: %w", err)
	}

	return signedToken, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	secretKey := os.Getenv("SECRET_KEY")
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func (token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		log.Printf("Error parsing token: %v\n", err)
		return nil, err
	}

	if !token.Valid {
		log.Println("Token is invalid")
		return nil, fmt.Errorf("token is invalid")
	}

	log.Printf("Token is valid. Claims: %+v\n", claims)
	return claims, nil
}
