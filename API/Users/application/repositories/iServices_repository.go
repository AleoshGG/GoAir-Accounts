package repositories

type Iservices interface {
	HashPassword(password string) (string, error)
	ValidatePassword(password, hash string) bool
	CreateJWT(id_user int, email string) (string, error)
}