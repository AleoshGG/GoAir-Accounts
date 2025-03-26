package repositories

import "GoAir-Accounts/API/Users/domain"

type Iservices interface {
	HashPassword(password string) (string, error)
	ValidatePassword(password, hash string) bool
	CreateJWT(id_user int, email string) (string, error)
	Auth(tokenString string) (domain.Claims, error)
}