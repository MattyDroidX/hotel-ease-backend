package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/MattyDroidX/hotel-ease-backend/api/db"
	"github.com/MattyDroidX/hotel-ease-backend/models"
	"github.com/MattyDroidX/hotel-ease-backend/utils"
)

// GetFuncionarios retorna a lista de todos os funcionários
// @Summary Lista de funcionários
// @Description Retorna todos os funcionários cadastrados
// @Tags Funcionários
// @Produce json
// @Success 200 {array} models.Funcionario
// @Failure 500 {object} map[string]string
// @Router /funcionarios [get]
func GetFuncionarios(c *gin.Context) {
	var funcionarios []models.Funcionario
	err := db.DB.Select(&funcionarios, "SELECT * FROM funcionarios")
	if err != nil {
		utils.Error("Erro ao buscar funcionários: %v", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   "error",
			"mensagem": "Erro ao buscar funcionários",
			"detalhe":  err.Error(),
		})
		return
	}

	utils.Info("Funcionarios carregados: " + string(len(funcionarios)))
	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"mensagem": "Funcionários encontrados com sucesso",
		"dados":    funcionarios,
	})
}

// GetFuncionarioByID retorna um funcionário por ID
// @Summary Busca funcionário por ID
// @Description Retorna todos os dados de um funcionário específico
// @Tags Funcionários
// @Produce json
// @Param id path string true "ID do Funcionário"
// @Success 200 {object} models.Funcionario
// @Failure 404 {object} map[string]string
// @Router /funcionarios/{id} [get]
func GetFuncionarioByID(c *gin.Context) {
	id := c.Param("id")
	var f models.Funcionario

	err := db.DB.Get(&f, "SELECT * FROM funcionarios WHERE id = $1", id)
	if err != nil {
		utils.Warn("Funcionário não encontrado (id=%s): %v", id, err.Error())
		c.JSON(http.StatusNotFound, gin.H{
			"status":   "error",
			"mensagem": "Funcionário não encontrado",
			"detalhe":  err.Error(),
		})
		return
	}

	utils.Info("Funcionario encontrada: " + f.ID)
	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"mensagem": "Funcionário encontrado com sucesso",
		"dados":    f,
	})
}

// CreateFuncionario cria o funcionário
// @Summary Criação de funcionário
// @Description Cadastro de Funcionário
// @Tags Funcionários
// @Accept json
// @Produce json
// @Success 201 {object} models.Funcionario
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /funcionarios [post]
func CreateFuncionario(c *gin.Context) {
	var f models.Funcionario
	if err := c.ShouldBindJSON(&f); err != nil {
		utils.Error("JSON inválido para criação de funcionário: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "error",
			"mensagem": "JSON inválido",
			"detalhe":  err.Error(),
		})
		return
	}
	f.ID = uuid.New().String()

	_, err := db.DB.NamedExec(`
		INSERT INTO funcionarios (id, nome, sobrenome, email, telefone, cargo, ativo)
		VALUES (:id, :nome, :sobrenome, :email, :telefone, :cargo, :ativo)
	`, &f)
	if err != nil {
		utils.Error("Erro ao salvar funcionário (id=%s): %v", f.ID, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   "error",
			"mensagem": "Erro ao salvar funcionário",
			"detalhe":  err.Error(),
		})
		return
	}

	utils.Info("Funcionario criado com sucesso (id=" + f.ID + ")")
	c.JSON(http.StatusCreated, gin.H{
		"status":   "success",
		"mensagem": "Funcionário criado com sucesso",
		"dados":    f,
	})
}

// UpdateFuncionario atualiza os dados de um funcionário
// @Summary Atualiza dados de um funcionário
// @Description Atualiza todos os campos de um funcionário existente
// @Tags Funcionários
// @Accept json
// @Produce json
// @Param id path string true "ID do Funcionário"
// @Success 200 {object} models.Funcionario
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /funcionarios/{id} [put]
func UpdateFuncionario(c *gin.Context) {
	id := c.Param("id")
	var f models.Funcionario

	if err := c.ShouldBindJSON(&f); err != nil {
		utils.Error("JSON inválido para atualizar funcionário (id=%s): %v", id, err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status":   "error",
			"mensagem": "JSON inválido",
			"detalhe":  err.Error(),
		})
		return
	}
	f.ID = id

	_, err := db.DB.NamedExec(`
		UPDATE funcionarios SET
			nome = :nome,
			sobrenome = :sobrenome,
			email = :email,
			telefone = :telefone,
			cargo = :cargo,
			ativo = :ativo
		WHERE id = :id
	`, &f)
	if err != nil {
		utils.Error("Erro ao atualizar funcionário (id=%s): %v", id, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   "error",
			"mensagem": "Erro ao atualizar funcionário",
			"detalhe":  err.Error(),
		})
		return
	}

	utils.Info("Dados de funcionario atualizado com sucesso (id=" + id + ")")
	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"mensagem": "Funcionário atualizado com sucesso",
		"dados":    f,
	})
}

// DeleteFuncionario apaga um funcionário
// @Summary Exclui um funcionário
// @Description Remove permanentemente um funcionário do sistema
// @Tags Funcionários
// @Produce json
// @Param id path string true "ID do Funcionário"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /funcionarios/{id} [delete]
func DeleteFuncionario(c *gin.Context) {
	id := c.Param("id")
	_, err := db.DB.Exec("DELETE FROM funcionarios WHERE id = $1", id)
	if err != nil {
		utils.Error("Erro ao excluir funcionário (id=%s): %v", id, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":   "error",
			"mensagem": "Erro ao excluir funcionário",
			"detalhe":  err.Error(),
		})
		return
	}
	
	utils.Info("Funcionario excluído com sucesso (id=" + id + ")")
	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"mensagem": "Funcionário excluído com sucesso",
	})
}
