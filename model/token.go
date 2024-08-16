package model

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(userId int64) (string, error) {
	tokenLifeSpan, err := strconv.Atoi("86400")

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifeSpan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("RDTHelper")))
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("RDTHelper")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}

func GetUserIdFromToken(tokenString string) (int64, error) {
	claims, err := ValidateToken(tokenString)

	if err != nil {
		return 0, err
	}

	userId := int64(claims["user_id"].(float64))

	return userId, nil
}
