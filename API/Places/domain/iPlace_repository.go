package domain

type IPlace interface {
	CreatePlace(p Place) (uint, error)
	DeletePlace(id_place int) (uint, error)
}