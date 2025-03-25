package usecases

import "GoAir-Accounts/API/Users/domain"

type CreateUser struct {
	db domain.IUser
}

func NewCreateUser(db domain.IUser) *CreateUser {
	return &CreateUser{db: db}
}

func (uc CreateUser) Run(u domain.User) (uint, error) {
	return uc.db.CreateUser(u)
}