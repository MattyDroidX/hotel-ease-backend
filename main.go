package main

import (
	"github.com/gin-gonic/gin"
	"github.com/MattyDroidX/hotel-ease-backend/api/db"
	"github.com/MattyDroidX/hotel-ease-backend/api/middleware"
	"github.com/MattyDroidX/hotel-ease-backend/api/routes"
)

func main() {
	db.Connect()
	db.CreateTables()

	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	routes.RegisterFuncionarioRoutes(router)
	routes.RegisterTarefaRoutes(router)

	router.Run(":8080")
}
