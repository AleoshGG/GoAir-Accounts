package usecases

import "GoAir-Accounts/API/Users/domain"

type DeleteUser struct {
	db domain.IUser
}

func NewDeleteUser(db domain.IUser) *DeleteUser {
	return &DeleteUser{db: db}
}

func (uc *DeleteUser) Run(id_user int) (uint, error) {
	return uc.db.DeleteUser(id_user)
} 