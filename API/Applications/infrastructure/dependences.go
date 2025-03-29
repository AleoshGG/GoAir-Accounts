package infrastructure

import "GoAir-Accounts/API/Applications/infrastructure/adapters"

var postgres *adapters.PostgreSQL
var jwt 	 *adapters.JWT
var rabbitmq *adapters.RabbitMQ

func GoDependences() {
	postgres = adapters.NewPostgreSQL()
	jwt = adapters.NewJWT()
	rabbitmq = adapters.NewRabbitMQ()
}

func GetPostgreSQL() *adapters.PostgreSQL {
	return postgres
}

func GetRabbitMQ() *adapters.RabbitMQ {
	return rabbitmq
}

func GetJWT() *adapters.JWT {
	return jwt
}
