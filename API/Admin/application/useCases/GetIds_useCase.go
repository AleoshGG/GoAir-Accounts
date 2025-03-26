package usecases

import (
	"GoAir-Accounts/API/Admin/domain/entities"
	"GoAir-Accounts/API/Admin/domain/repository"
)
type GetIds struct {
	db repository.IAdmin
}

func NewGetIds(db repository.IAdmin) *GetIds {
	return &GetIds{db: db}
}

func (uc *GetIds) Run(id_place int) []entities.Sensor {
	return uc.db.GetIds(id_place)
}