package models

import "time"

type Tarefa struct {
	ID          string `db:"id" json:"id" example:"123e4567-e89b-12d3-a456-426614174001"`
	Numero      string `db:"numero" json:"numero" example:"T-2025-001"`
	Funcionario string `db:"funcionario" json:"funcionario" example:"123e4567-e89b-12d3-a456-426614174000"`
	Descricao   string `db:"descricao" json:"descricao" example:"Limpar o quarto 204"`
	DataHora    time.Time `db:"data_hora" json:"data_hora" example:"2025-07-05T15:04:05Z"`
	Status      string `db:"status" json:"status" example:"pendente"` 
	Tipo        string `db:"tipo" json:"tipo" example:"Limpeza"`    
}
