package routes

import (
	"GoAir-Accounts/API/Applications/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	applicationRouter := r.Group("/applications")
	{
		applicationRouter.POST("/", controllers.NewCreCreateApplicationController().CreateApplication)
	}
}