package infrastructure

import "GoAir-Accounts/API/Admin/infrastructure/adapters"

var postgres *adapters.PostgreSQL
var JWT *adapters.JWT

func GoDependences() {
	postgres = adapters.NewPostgreSQL()
	JWT = adapters.NewJWT()
}

func GetPostgreSQL() *adapters.PostgreSQL {
	return postgres
}

func GetJWT() *adapters.JWT {
	return JWT
}