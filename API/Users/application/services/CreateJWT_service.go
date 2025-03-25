package services

import "GoAir-Accounts/API/Users/application/repositories"

type CreateJWT struct {
	jwt repositories.Iservices
}

func NewCreateJWT(jwt repositories.Iservices) *CreateJWT {
	return &CreateJWT{jwt: jwt}
}

func (t *CreateJWT) Run(id_user int, email string) (string, error) {
	return t.jwt.CreateJWT(id_user, email)
}