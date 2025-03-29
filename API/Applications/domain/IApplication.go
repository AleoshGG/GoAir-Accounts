package domain

type IApplication interface {
	CreateApplication(id_user int) (RabbitMessage, error)
	GetApplicationByUser(id_user int) []Application
}