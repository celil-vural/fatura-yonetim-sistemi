package helpers

import (
	"errors"
	"fatura-yonetim-sistemi/entity/models"
	"fatura-yonetim-sistemi/infrastructure/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// parseToken is a function to parse token
func parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.ConfigEnv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ExtractClaims is a function to extract claims from token
func ExtractClaims(tokenString string) (jwt.MapClaims, error) {
	token, err := parseToken(tokenString)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token or claims")
}

// CreateClaims is a function to create jwt claims
func CreateClaims(user *models.User) jwt.MapClaims {
	claims := jwt.MapClaims{
		"id":      user.ID,
		"name":    user.Name,
		"surname": user.Surname,
		"email":   user.Email,
		"exp":     jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		"iat":     jwt.NewNumericDate(time.Now()),
	}
	return claims
}

// GenerateJwt is a function to generate jwt token
func GenerateJwt(user *models.User) (string, error) {
	claims := CreateClaims(user)
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	return token.SignedString([]byte(config.ConfigEnv("JWT_SECRET")))
}
