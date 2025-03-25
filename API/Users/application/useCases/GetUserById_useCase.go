package usecases

import "GoAir-Accounts/API/Users/domain"

type GetUserById struct {
	db domain.IUser
}

func NewGetUserById(db domain.IUser) *GetUserById {
	return &GetUserById{db: db}
}

func (uc *GetUserById) Run(id_user int) domain.User {
	return uc.db.GetUserById(id_user)
}