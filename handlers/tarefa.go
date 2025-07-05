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
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"mensagem": "Erro ao buscar tarefas",
			"detalhe": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"mensagem": "Tarefas carregadas com sucesso",
		"dados": tarefas,
	})
}

func GetTarefaByID(c *gin.Context) {
	id := c.Param("id")
	var t models.Tarefa

	err := db.DB.Get(&t, "SELECT * FROM tarefas WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"mensagem": "Tarefa não encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"mensagem": "Tarefa encontrada",
		"dados": t,
	})
}

func CreateTarefa(c *gin.Context) {
	var t models.Tarefa
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"mensagem": "JSON inválido",
			"detalhe": err.Error(),
		})
		return
	}
	t.ID = uuid.New().String()

	_, err := db.DB.Exec(`
		INSERT INTO tarefas (id, numero, funcionario, descricao, data_hora, status, tipo)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		t.ID, t.Numero, t.Funcionario, t.Descricao, t.DataHora, t.Status, t.Tipo,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"mensagem": "Erro ao salvar tarefa",
			"detalhe": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"mensagem": "Tarefa criada com sucesso",
		"dados": t,
	})
}

func UpdateTarefa(c *gin.Context) {
	id := c.Param("id")
	var t models.Tarefa

	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"mensagem": "JSON inválido",
			"detalhe": err.Error(),
		})
		return
	}
	t.ID = id

	_, err := db.DB.NamedExec(`
		UPDATE tarefas SET
			numero = :numero,
			funcionario = :funcionario,
			descricao = :descricao,
			data_hora = :data_hora,
			status = :status,
			tipo = :tipo
		WHERE id = :id
	`, &t)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"mensagem": "Erro ao atualizar tarefa",
			"detalhe": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"mensagem": "Tarefa atualizada com sucesso",
		"dados": t,
	})
}

func DeleteTarefa(c *gin.Context) {
	id := c.Param("id")

	_, err := db.DB.Exec("DELETE FROM tarefas WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"mensagem": "Erro ao excluir tarefa",
			"detalhe": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"mensagem": "Tarefa excluída com sucesso",
	})
}
