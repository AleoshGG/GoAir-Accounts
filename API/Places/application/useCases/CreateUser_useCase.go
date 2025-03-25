package usecases

import "GoAir-Accounts/API/Places/domain"
type CreatePalce struct {
	db domain.IPlace
}

func NewCreatePlace(db domain.IPlace) *CreatePalce {
	return &CreatePalce{db: db}
}

func (uc CreatePalce) Run(p domain.Place) (uint, error) {
	return uc.db.CreatePlace(p)
}