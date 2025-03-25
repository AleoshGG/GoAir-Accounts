package infrastructure

import "GoAir-Accounts/API/Places/infrastructure/adapters"

var postgres *adapters.PostgreSQL

func GoDependences() {
	postgres = adapters.NewPostgreSQL()
}

func GetPostgreSQL() *adapters.PostgreSQL {
	return postgres
}
