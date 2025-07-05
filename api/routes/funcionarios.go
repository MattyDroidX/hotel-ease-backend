package routes

import (
	"github.com/MattyDroidX/hotel-ease-backend/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterFuncionarioRoutes(r *gin.Engine) {
	funcionarios := r.Group("/funcionarios")
	{
		funcionarios.GET("", handlers.GetFuncionarios)
		funcionarios.GET("/:id", handlers.GetFuncionarioByID)
		funcionarios.POST("", handlers.CreateFuncionario)
		funcionarios.PUT("/:id", handlers.UpdateFuncionario)
		funcionarios.DELETE("/:id", handlers.DeleteFuncionario)
	}
}
