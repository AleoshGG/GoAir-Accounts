package usecases

import "GoAir-Accounts/API/Users/domain"

type GetUserByEmail struct {
	db domain.IUser
}

func NewGetUserByEmail(db domain.IUser) *GetUserByEmail {
	return &GetUserByEmail{db: db}
}

func (uc GetUserByEmail) Run(email string) domain.User {
	return uc.db.GetUserByEmail(email)
}