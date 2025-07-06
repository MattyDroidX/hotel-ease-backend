package models

import "time"

type TarefaDTO struct {
	ID            string    `json:"id"            db:"id"`
	Numero        string    `json:"numero"        db:"numero"`
	Descricao     string    `json:"descricao"     db:"descricao"`
	DataHora      time.Time `json:"dataHora"      db:"data_hora"`
	Status        string    `json:"status"        db:"status"`
	Tipo          string    `json:"tipo"          db:"tipo"`
	Funcionario   string    `json:"funcionario"   db:"funcionario_nome"`
	FuncionarioID string    `json:"funcionarioId" db:"funcionario_id"`
}

type TarefaIn struct {
	Numero        string `json:"numero"        binding:"required"`
	FuncionarioID string `json:"funcionarioId" binding:"required,uuid"`
	Descricao     string `json:"descricao"     binding:"required"`
	DataHora      string `json:"dataHora"      binding:"required"`
	Status        string `json:"status"        binding:"required"`
	Tipo          string `json:"tipo"          binding:"required"`
}