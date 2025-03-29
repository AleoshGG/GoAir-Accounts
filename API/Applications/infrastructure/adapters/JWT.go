package adapters

import (
	"GoAir-Accounts/API/Applications/domain"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type JWT struct{}

func NewJWT() *JWT {
	return &JWT{}
}

func (j *JWT) Auth(tokenString string) (domain.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return domain.Claims{}, err
	}

	claims, ok := token.Claims.(*domain.Claims)

	if !ok || !token.Valid {
		return domain.Claims{}, fmt.Errorf("token inválido")
	}

	return *claims, nil
}