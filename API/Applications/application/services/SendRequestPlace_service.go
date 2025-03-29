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

func (s *SendRequestPlace) Run(msg domain.RabbitMessage) {
	s.msg.SendRequestPlace(msg)
}