package services

import "GoAir-Accounts/API/Users/application/repositories"

type HashPassword struct {
	hs repositories.Iservices
}

func NewHashPassword(hs repositories.Iservices) *HashPassword {
	return &HashPassword{hs: hs}
}

func (s *HashPassword) Run(password string) (string, error) {
	return s.hs.HashPassword(password)
}