package usecases

import (
	"GoAir-Accounts/API/Users/domain"
)

type GetPlaces struct {
	db domain.IUser
}

func NewGetPlaces(db domain.IUser) *GetPlaces {
	return &GetPlaces{db: db}
}

func (uc *GetPlaces) Run(id_user int) []domain.Place {
	return uc.db.GetPlaces(id_user)
}