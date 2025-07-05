package models

type Funcionario struct {
	ID        string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Nome      string `json:"nome" example:"João"`
	Sobrenome string `json:"sobrenome" example:"Silva"`
	Email     string `json:"email" example:"joao.silva@email.com"`
	Telefone  string `json:"telefone" example:"+55 11 91234-5678"`
	Cargo     string `json:"cargo" example:"Recepcionista"`
	Ativo     bool   `json:"ativo" example:"true"`
}
