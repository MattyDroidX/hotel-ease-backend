package routes

import (
	"github.com/MattyDroidX/hotel-ease-backend/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterTarefaRoutes(r *gin.Engine) {
	tarefas := r.Group("/tarefas")
	{
		tarefas.GET("", handlers.GetTarefas)
		tarefas.GET("/:id", handlers.GetTarefaByID)
		tarefas.POST("", handlers.CreateTarefa)
		tarefas.PUT("/:id", handlers.UpdateTarefa)
		tarefas.DELETE("/:id", handlers.DeleteTarefa)
	}
}
