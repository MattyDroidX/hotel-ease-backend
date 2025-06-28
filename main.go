package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"hotel-ease-backend/db"
	"hotel-ease-backend/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	db.Connect()

	router := gin.Default()

	// Rotas
	router.GET("/funcionarios", handlers.GetFuncionarios)
	router.POST("/funcionarios", handlers.CreateFuncionario)
	// + outras rotas

	router.Run(":" + os.Getenv("PORT"))
}
