package services

import "GoAir-Accounts/API/Admin/domain/entities"

type IServices interface {
	CreateJWT(admin entities.Admin) (string, error)
	Auth(tokenString string) (entities.Claims, error) 
}