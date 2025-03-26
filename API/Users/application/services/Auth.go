package services

import (
	"GoAir-Accounts/API/Users/application/repositories"
	"GoAir-Accounts/API/Users/domain"
)

type Auth struct {
	jwt repositories.Iservices
}

func NewAuth(jwt repositories.Iservices) *Auth {
	return &Auth{jwt: jwt}
}

func (t *Auth) Run(token string) (domain.Claims, error) {
	return t.jwt.Auth(token)
}