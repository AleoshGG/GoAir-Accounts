package usecases

import (
	"GoAir-Accounts/API/Places/domain"
)

type DeletePlace struct {
	db domain.IPlace
}

func NewDeletePlace(db domain.IPlace) *DeletePlace {
	return &DeletePlace{db: db}
}

func (uc *DeletePlace) Run(id_place int) (uint, error) {
	return uc.db.DeletePlace(id_place)
} 