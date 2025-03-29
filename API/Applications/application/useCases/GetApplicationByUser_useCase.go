package usecases

import "GoAir-Accounts/API/Applications/domain"

type GetApplicationByUser struct {
	db domain.IApplication
}

func NewGetApplicationByUser(db domain.IApplication) *GetApplicationByUser {
	return &GetApplicationByUser{db: db}
}

func (uc *GetApplicationByUser) Run(id_user int) []domain.Application {
	return uc.db.GetApplicationByUser(id_user)
}