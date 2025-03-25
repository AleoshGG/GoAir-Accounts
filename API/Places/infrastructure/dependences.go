package infrastructure

import "GoAir-Accounts/API/Users/infrastructure/adapters"

var postgres *adapters.PostgreSQL
var bcrypt *adapters.Bcrypt

func GoDependences() {
	postgres = adapters.NewPostgreSQL()
	bcrypt = adapters.NewBcrypt()
}

func GetPostgreSQL() *adapters.PostgreSQL {
	return postgres
}

func GetBcrypt() *adapters.Bcrypt {
	return bcrypt
}