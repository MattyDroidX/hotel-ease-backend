package main

import (
	"github.com/MattyDroidX/hotel-ease-backend/api/db"
	"github.com/MattyDroidX/hotel-ease-backend/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	router := gin.Default()

	// Middleware CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Endpoint Funcionarios

	router.GET("/funcionarios", handlers.GetFuncionarios)
	router.GET("/funcionarios/:id", handlers.GetFuncionarioByID)
	router.POST("/funcionarios", handlers.CreateFuncionario)
	router.DELETE("/funcionarios/:id", handlers.DeleteFuncionario)
	router.PUT("/funcionarios/:id", handlers.UpdateFuncionario)

	// Endpoint Tarefas

	router.GET("/tarefas", handlers.GetTarefas)
	router.GET("/tarefas/:id", handlers.GetTarefaByID)
	router.POST("/tarefas", handlers.CreateTarefa)
	router.DELETE("/tarefas/:id", handlers.DeleteTarefa)
	router.PUT("/tarefas/:id", handlers.UpdateTarefa)

	router.Run(":8080")
}
