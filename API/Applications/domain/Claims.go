package domain

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Id_user int    `json:"id_user"`
	Email   string `json:"email"`
	jwt.RegisteredClaims
}