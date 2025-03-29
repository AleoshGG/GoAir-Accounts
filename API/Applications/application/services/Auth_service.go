package services

import (
	"GoAir-Accounts/API/Applications/application/repositories"
	"GoAir-Accounts/API/Applications/domain"
)

type Auth struct {
	jwt repositories.IJWT
}

func NewAuth(jwt repositories.IJWT) *Auth {
	return &Auth{jwt: jwt}
}

func (t *Auth) Run(token string) (domain.Claims, error) {
	return t.jwt.Auth(token)
}