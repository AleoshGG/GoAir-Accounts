package usecases

import (
	"GoAir-Accounts/API/Admin/domain/entities"
	"GoAir-Accounts/API/Admin/domain/repository"
)

type CreateId struct {
	db repository.IAdmin
}

func NewCreateId(db repository.IAdmin) *CreateId {
	return &CreateId{db: db}
}

func (uc *CreateId) Run(p entities.Place, u entities.User) (string, error) {
	return uc.db.CreateId(u, p)
}