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
