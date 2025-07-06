package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/MattyDroidX/hotel-ease-backend/api/db"
	"github.com/MattyDroidX/hotel-ease-backend/api/middleware"
	"github.com/MattyDroidX/hotel-ease-backend/api/routes"
	"github.com/MattyDroidX/hotel-ease-backend/utils"

	// Swagger
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "github.com/MattyDroidX/hotel-ease-backend/docs"
)

// @title HotelEase API
// @version 1.0
// @description API para gerenciamento de tarefas e funcionários no HotelEase
// @contact.name Suporte HotelEase
// @host localhost:8080
// @BasePath /
func main() {
	// Iniciar logs
	utils.InitLogger()

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", os.ModePerm)
		utils.Info("Diretório de logs criado.")
	}

	// Carregar variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		utils.Warn("⚠️  Arquivo .env não encontrado, usando variáveis padrão...")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		utils.Warn("⚠️  Variável PORT não encontrada, usando porta padrão:", port)
	}

	// Conectar ao banco e criar tabelas
	db.Connect()
	db.CreateTables()

	// Criar roteador
	router := gin.New()
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.GinLogger())
	router.Use(middleware.RequestLogger())

	// Documentação Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Registrar rotas
	routes.RegisterFuncionarioRoutes(router)
	routes.RegisterTarefaRoutes(router)

	utils.Info("Servidor rodando em: http://localhost:%s", port)
	utils.Info("Swagger disponível em: http://localhost:%s/swagger/index.html", port)

	// Rodar servidor
	err = router.Run(":" + port)
	if err != nil {
		utils.Error("Erro ao iniciar servidor: %v", err.Error())
	}
}
