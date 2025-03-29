package services

import (
	"GoAir-Accounts/API/Applications/application/repositories"
	"GoAir-Accounts/API/Applications/domain"
)

type SendRequestPlace struct {
	msg repositories.IRabbit
}

func NewSendRequestPlace(msg repositories.IRabbit) *SendRequestPlace{
	return &SendRequestPlace{msg: msg}
}

func (s *SendRequestPlace) Run(request domain.Application) {
	s.msg.SendRequestPlace(request)
}