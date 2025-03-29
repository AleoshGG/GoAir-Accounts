package repositories

import "GoAir-Accounts/API/Applications/domain"

type IJWT interface {
	Auth(tokenString string) (domain.Claims, error)
}