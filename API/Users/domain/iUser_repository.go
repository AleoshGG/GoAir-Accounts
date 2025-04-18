package domain

type IUser interface {
	CreateUser(u User) (uint, error)
	DeleteUser(id_user int) (uint, error)
	GetUserByEmail(email string) User
	GetUserById(id_user int) User
	GetPlaces(id_user int) []Place
}