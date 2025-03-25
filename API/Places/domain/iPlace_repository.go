package domain

type IPlaces interface {
	CreatePlace(p Places) (uint, error)
	DeletePlace(id_place int) (uint, error)
	GetPlaceById(id_place int) []Places
}