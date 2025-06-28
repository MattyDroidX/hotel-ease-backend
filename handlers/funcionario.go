package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/MattyDroidX/hotel-ease-backend/db"
	"github.com/MattyDroidX/hotel-ease-backend/models"
	"github.com/google/uuid"
)

func GetFuncionarios(c *gin.Context) {
	var funcionarios []models.Funcionario
	err := db.DB.Select(&funcionarios, "SELECT * FROM funcionarios")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusOK, funcionarios)
}

func CreateFuncionario(c *gin.Context) {
	var f models.Funcionario
	if err := c.ShouldBindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	f.ID = uuid.New().String()

	_, err := db.DB.Exec(`INSERT INTO funcionarios (id, nome, sobrenome, email, telefone, cargo, ativo)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		f.ID, f.Nome, f.Sobrenome, f.Email, f.Telefone, f.Cargo, f.Ativo,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, f)
}
