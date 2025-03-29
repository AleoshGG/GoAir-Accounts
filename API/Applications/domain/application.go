package domain

type StatusApplication string

const (
	Requested StatusApplication = "requested"
	Pending   StatusApplication = "pending"
	Complete  StatusApplication = "complete"
)

type Application struct {
	Id_Application     int
	Status_application StatusApplication
	Id_user            int
}