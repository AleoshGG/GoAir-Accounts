package repositories

type Iservices interface {
	HashPassword(password string) (string, error)
	ValidatePassword(password, hash string) bool
}