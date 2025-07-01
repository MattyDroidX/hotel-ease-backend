package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"

	"github.com/MattyDroidX/hotel-ease-backend/api/db"
	"github.com/MattyDroidX/hotel-ease-backend/models"
)

func GetFuncionarios(c *gin.Context) {
	var funcionarios []models.Funcionario
	err := db.DB.Select(&funcionarios, "SELECT * FROM funcionarios")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao buscar funcionários"})
		return
	}
	c.JSON(http.StatusOK, funcionarios)
}

func GetFuncionarioByID(c *gin.Context) {
	id := c.Param("id")

	var f models.Funcionario
	err := db.DB.Get(&f, "SELECT * FROM funcionarios WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Funcionário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, f)
}

func CreateFuncionario(c *gin.Context) {
	var f models.Funcionario
	if err := c.ShouldBindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}
	f.ID = uuid.New().String()

	_, err := db.DB.NamedExec(`
		INSERT INTO funcionarios (id, nome, sobrenome, email, telefone, cargo, ativo)
		VALUES (:id, :nome, :sobrenome, :email, :telefone, :cargo, :ativo)
	`, &f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao salvar no banco"})
		return
	}

	c.JSON(http.StatusCreated, f)

}

func DeleteFuncionario(c *gin.Context) {
	id := c.Param("id")

	_, err := db.DB.Exec("DELETE FROM funcionarios WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao excluir funcionário"})
		return
	}

	c.Status(http.StatusNoContent)
}

func UpdateFuncionario(c *gin.Context) {
	id := c.Param("id")

	var f models.Funcionario
	if err := c.ShouldBindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}

	f.ID = id

	_, err := db.DB.NamedExec(`
		UPDATE funcionarios 
		SET nome = :nome, sobrenome = :sobrenome, email = :email, telefone = :telefone, cargo = :cargo, ativo = :ativo 
		WHERE id = :id
	`, &f)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao atualizar funcionário"})
		return
	}

	c.JSON(http.StatusOK, f)
}