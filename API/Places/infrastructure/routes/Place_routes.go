package routes

import (
	"GoAir-Accounts/API/Places/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	placeRoutes := r.Group("/places")
	{
		placeRoutes.POST("/", controllers.NewCreatePalceController().AddPlace)
		placeRoutes.DELETE("/:id", controllers.NewDeletePlace().DeletePlace)
	}
}