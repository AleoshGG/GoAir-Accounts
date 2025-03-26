package repository

import "GoAir-Accounts/API/Admin/domain/entities"

type IAdmin interface {
	GetAdmin() entities.Admin
	CreatePlace(p entities.Place, id_user int) (uint, error)
	SearchUser(last_name string) entities.User 
	CreateId(u entities.User, p entities.Place) (string, error)
	GetIds(id_place int) []entities.Sensor
}