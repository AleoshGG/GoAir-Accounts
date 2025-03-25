package services

import "GoAir-Accounts/API/Users/application/repositories"

type ValidatePassword struct {
	hs repositories.Iservices
}

func NewValidatePassword(hs repositories.Iservices) *ValidatePassword {
	return &ValidatePassword{hs: hs}
}

func (s *ValidatePassword) Run(password string, id_user int) bool {
	return s.hs.ValidatePassword(password, id_user)
} 