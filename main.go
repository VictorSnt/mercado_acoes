package main

import (
	"mercado/acoes/routes"

	_ "mercado/acoes/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Monitor de Sites API
// @version 1.0
// @description API para verificar status de sites.
// @host localhost:8080
// @BasePath /
func main() {
	routes.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.InitializeUserRoutes()
	routes.InitializeEquitiesRoutes()
	routes.InitializeTransactionRoutes()
	routes.Router.Run(":8080")
}
