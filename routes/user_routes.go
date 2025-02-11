package routes

import (
	"mercado/acoes/controllers"
)

func InitializeUserRoutes() {
	r := Router
	v1 := r.Group("/api/v1")
	v1.GET("/users/:id", controllers.GetUser)
	v1.GET("/users", controllers.GetAllUser)
	v1.DELETE("/users/:id", controllers.DeleteUser)
	v1.POST("/users", controllers.CreateUser)
	v1.PUT("/users/:id", controllers.UpdateUser)
	v1.GET("users/:id/equities", controllers.UserEquitieStock)
}
