package usecases

import "GoAir-Accounts/API/Applications/domain"

type CreateApplication struct {
	db domain.IApplication
}

func NewCreateApplication(db domain.IApplication) * CreateApplication {
	return &CreateApplication{db: db}
}

func (uc *CreateApplication) Run(id_user int) (uint, error) {
	return uc.db.CreateApplication(id_user)
}