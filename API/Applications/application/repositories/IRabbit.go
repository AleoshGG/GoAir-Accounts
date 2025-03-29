package repositories

import "GoAir-Accounts/API/Applications/domain"

type IRabbit interface {
	SendRequestPlace(msg domain.RabbitMessage)
}