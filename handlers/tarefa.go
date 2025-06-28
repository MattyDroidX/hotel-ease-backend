package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/MattyDroidX/hotel-ease-backend/api/db"
	"github.com/MattyDroidX/hotel-ease-backend/models"
)

func GetTarefas(c *gin.Context) {
	var tarefas []models.Tarefa
	err := db.DB.Select(&tarefas, "SELECT * FROM tarefas")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tarefas)
}

func CreateTarefa(c *gin.Context) {
	var t models.Tarefa
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	t.ID = uuid.New().String()

	_, err := db.DB.Exec(`
		INSERT INTO tarefas (id, numero, funcionario, descricao, data_hora, status, tipo)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		t.ID, t.Numero, t.Funcionario, t.Descricao, t.DataHora, t.Status, t.Tipo,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, t)
}
