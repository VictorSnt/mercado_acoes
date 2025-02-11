package routes

import (
	"mercado/acoes/controllers"
)

func InitializeTransactionRoutes() {
	r := Router
	v1 := r.Group("/api/v1")
	v1.GET("/transactions/user/:id", controllers.GetAllUserTransaction)
	v1.POST("/transactions", controllers.CreateTransaction)
}
