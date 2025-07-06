package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/MattyDroidX/hotel-ease-backend/api/db"
	"github.com/MattyDroidX/hotel-ease-backend/models"
	"github.com/MattyDroidX/hotel-ease-backend/utils"
	"github.com/MattyDroidX/hotel-ease-backend/helpers"
)

const qDTO = `
SELECT t.id, t.numero, t.descricao, t.data_hora,
       t.status, t.tipo,
       concat(f.nome,' ',f.sobrenome) AS funcionario_nome,
       f.id                           AS funcionario_id
  FROM tarefas t
  LEFT JOIN funcionarios f ON f.id = t.funcionario
`

func GetTarefas(c *gin.Context) {
	var out []models.TarefaDTO
	if err := db.DB.Select(&out, qDTO+" ORDER BY t.data_hora DESC"); err != nil {
		utils.Error("Erro ao buscar tarefas: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "db"})
		return
	}
	utils.Info("Tarefas carregadas: %d", len(out))
	c.JSON(http.StatusOK, gin.H{"dados": out})
}

func GetTarefaByID(c *gin.Context) {
	id := c.Param("id")

	var dto models.TarefaDTO
	if err := db.DB.Get(&dto, qDTO+" WHERE t.id=$1 LIMIT 1", id); err != nil {
		utils.Warn("Tarefa não encontrada: %s", id)
		c.JSON(http.StatusNotFound, gin.H{"erro": "não encontrada"})
		return
	}

	c.JSON(http.StatusOK, dto)
}

func CreateTarefa(c *gin.Context) {
    var in models.TarefaIn
    if err := c.ShouldBindJSON(&in); err != nil {
        utils.Error("JSON inválido (tarefa): %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"erro": "json inválido"})
        return
    }

    var ok bool
    if err := db.DB.Get(&ok,
        "SELECT true FROM funcionarios WHERE id=$1 LIMIT 1", in.FuncionarioID); err != nil || !ok {
        utils.Warn("Funcionário não encontrado: %s", in.FuncionarioID)
        c.JSON(http.StatusBadRequest, gin.H{"erro": "funcionário inexistente"})
        return
    }

    t, err := helpers.ParseHTMLDateTime(in.DataHora)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"erro": "dataHora inválida"})
        return
    }

    id := uuid.NewString()
    if _, err := db.DB.Exec(`
       INSERT INTO tarefas (id,numero,funcionario,descricao,data_hora,status,tipo)
       VALUES ($1,$2,$3,$4,$5,$6,$7)`,
        id, in.Numero, in.FuncionarioID, in.Descricao, t, in.Status, in.Tipo); err != nil {
        utils.Error("Erro insert tarefa: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"erro": "db"})
        return
    }
    utils.Info("Tarefa criada %s", id)
    c.JSON(http.StatusCreated, gin.H{"id": id})
}

func UpdateTarefa(c *gin.Context) {
    id := c.Param("id")
    var in models.TarefaIn
    if err := c.ShouldBindJSON(&in); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"erro": "json inválido"})
        return
    }
    var ok bool
    if err := db.DB.Get(&ok,
        "SELECT true FROM funcionarios WHERE id=$1 LIMIT 1", in.FuncionarioID); err != nil || !ok {
        c.JSON(http.StatusBadRequest, gin.H{"erro": "funcionário inexistente"})
        return
    }

    t, err := helpers.ParseHTMLDateTime(in.DataHora)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"erro": "dataHora inválida"})
        return
    }

    if _, err := db.DB.Exec(`
        UPDATE tarefas SET
          numero=$1, funcionario=$2, descricao=$3, data_hora=$4, status=$5, tipo=$6
        WHERE id=$7`,
        in.Numero, in.FuncionarioID, in.Descricao, t, in.Status, in.Tipo, id); err != nil {
        utils.Error("Erro update tarefa: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"erro": "db"})
        return
    }
    c.Status(http.StatusOK)
}


func DeleteTarefa(c *gin.Context) {
	id := c.Param("id")
	if _, err := db.DB.Exec(`DELETE FROM tarefas WHERE id=$1`, id); err != nil {
		utils.Error("Erro delete tarefa: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "db"})
		return
	}
	c.Status(http.StatusOK)
}
