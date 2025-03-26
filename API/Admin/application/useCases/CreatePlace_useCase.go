package usecases

import (
	"GoAir-Accounts/API/Admin/domain/entities"
	"GoAir-Accounts/API/Admin/domain/repository"
)

type CreatePlace struct {
	db repository.IAdmin
}

func NewCreatePlace(db repository.IAdmin) *CreatePlace {
	return &CreatePlace{db: db}
}

func (uc CreatePlace) Run(p entities.Place, id_user int) (uint, error) {
	return uc.db.CreatePlace(p, id_user)
}