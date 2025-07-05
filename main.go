package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/MattyDroidX/hotel-ease-backend/api/db"
	"github.com/MattyDroidX/hotel-ease-backend/api/middleware"
	"github.com/MattyDroidX/hotel-ease-backend/api/routes"

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
	// Configurar logs
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Carregar variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Arquivo .env não encontrado, usando variáveis padrão...")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("⚠️  Variável PORT não encontrada, usando porta padrão:", port)
	}

	// Conectar ao banco e criar tabelas
	db.Connect()
	db.CreateTables()

	// Criar roteador
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	// Documentação Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Registrar rotas
	routes.RegisterFuncionarioRoutes(router)
	routes.RegisterTarefaRoutes(router)

	log.Printf("✅ Servidor rodando em: http://localhost:%s", port)
	log.Printf("📚 Swagger disponível em: http://localhost:%s/swagger/index.html", port)

	// Rodar servidor
	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("❌ Erro ao iniciar servidor: %v", err)
	}
}
