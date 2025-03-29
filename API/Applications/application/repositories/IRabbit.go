package repositories

import "GoAir-Accounts/API/Applications/domain"

type IRabbit interface {
	SendRequestPlace(request domain.Application)
}