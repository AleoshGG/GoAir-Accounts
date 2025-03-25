package repositories

type Iservices interface {
	HashPassword(password string) string
	ValidatePassword(password string, id_user int) bool
}