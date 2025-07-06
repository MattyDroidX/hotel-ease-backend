package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/MattyDroidX/hotel-ease-backend/api/db"
	"github.com/MattyDroidX/hotel-ease-backend/models"
	"github.com/MattyDroidX/hotel-ease-backend/utils"
)

// GetTarefas retorna a lista de todas as tarefas
func GetTarefas(c *gin.Context) {
	var tarefas []models.Tarefa
	err := db.DB.Select(&tarefas, "SELECT 
		t.id, t.numero, t.descricao, t.data_hora, t.status, t.tipo,
		f.id as funcionario_id, f.nome as funcionario_nome, f.sobrenome as funcionario_sobrenome
		FROM tarefas t
		LEFT JOIN funcionarios f ON f.id = t.funcionario
		")
	if err != nil {
		utils.Error("Erro ao buscar tarefas: " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   "error",
			"mensagem": "Erro ao buscar tarefas",
			"detalhe":  err.Error(),
		})
		return
	}
	utils.Info("Tarefas carregadas: " + string(len(tarefas)))
	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"mensagem": "Tarefas carregadas com sucesso",
		"dados":    tarefas,
	})
}

// GetTarefaByID retorna uma tarefa por ID
func GetTarefaByID(c *gin.Context) {
	id := c.Param("id")
	var t models.Tarefa

	err := db.DB.Get(&t, "SELECT * FROM tarefas WHERE id = $1", id)
	if err != nil {
		utils.Warn("Tarefa não encontrada (id=" + id + "): " + err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "error",
			"mensagem": "Tarefa não encontrada",
			"detalhe":  err.Error(),
		})
		return
	}

	utils.Info("Tarefa encontrada: " + t.ID)
	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"mensagem": "Tarefa encontrada",
		"dados":    t,
	})
}

// CreateTarefa cria uma nova tarefa
func CreateTarefa(c *gin.Context) {
	var t models.Tarefa
	if err := c.ShouldBindJSON(&t); err != nil {
		utils.Error("JSON inválido para criação de tarefa: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "error",
			"mensagem": "JSON inválido",
			"detalhe":  err.Error(),
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
		utils.Error("Erro ao salvar tarefa (id=" + t.ID + "): " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   "error",
			"mensagem": "Erro ao salvar tarefa",
			"detalhe":  err.Error(),
		})
		return
	}

	utils.Info("Tarefa criada com sucesso (id=" + t.ID + ")")
	c.JSON(http.StatusCreated, gin.H{
		"status":   "success",
		"mensagem": "Tarefa criada com sucesso",
		"dados":    t,
	})
}

// UpdateTarefa atualiza os dados de uma tarefa
func UpdateTarefa(c *gin.Context) {
	id := c.Param("id")
	var t models.Tarefa

	if err := c.ShouldBindJSON(&t); err != nil {
		utils.Error("JSON inválido para atualizar tarefa (id=" + id + "): " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "error",
			"mensagem": "JSON inválido",
			"detalhe":  err.Error(),
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
		utils.Error("Erro ao atualizar tarefa (id=" + id + "): " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   "error",
			"mensagem": "Erro ao atualizar tarefa",
			"detalhe":  err.Error(),
		})
		return
	}

	utils.Info("Tarefa atualizada com sucesso (id=" + id + ")")
	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"mensagem": "Tarefa atualizada com sucesso",
		"dados":    t,
	})
}

// DeleteTarefa exclui uma tarefa
func DeleteTarefa(c *gin.Context) {
	id := c.Param("id")

	_, err := db.DB.Exec("DELETE FROM tarefas WHERE id = $1", id)
	if err != nil {
		utils.Error("Erro ao excluir tarefa (id=" + id + "): " + err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   "error",
			"mensagem": "Erro ao excluir tarefa",
			"detalhe":  err.Error(),
		})
		return
	}

	utils.Info("Tarefa excluída com sucesso (id=" + id + ")")
	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"mensagem": "Tarefa excluída com sucesso",
	})
}
