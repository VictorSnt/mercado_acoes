package routes

import (
	"mercado/acoes/controllers"
)

func InitializeEquitiesRoutes() {
	r := Router
	v1 := r.Group("/api/v1")
	v1.GET("/equities", controllers.GetAllEquitie)
	v1.POST("/equities", controllers.CreateEquitie)
	v1.GET("/equities/:id", controllers.GetEquitie)
	v1.PUT("/equities/:id", controllers.UpdateEquitie)
}
