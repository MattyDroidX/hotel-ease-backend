package main

import (
	"github.com/MattyDroidX/hotel-ease-backend/api/db"
	"github.com/MattyDroidX/hotel-ease-backend/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	router := gin.Default()

	// Funcionários
	router.GET("/funcionarios", handlers.GetFuncionarios)
	router.POST("/funcionarios", handlers.CreateFuncionario)

	// Tarefas
	router.GET("/tarefas", handlers.GetTarefas)
	router.POST("/tarefas", handlers.CreateTarefa)

	router.Run(":8080")
}
