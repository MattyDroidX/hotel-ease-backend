package models

type Funcionario struct {
	ID        string `json:"id" db:"id"`
	Nome      string `json:"nome" db:"nome"`
	Sobrenome string `json:"sobrenome" db:"sobrenome"`
	Email     string `json:"email" db:"email"`
	Telefone  string `json:"telefone" db:"telefone"`
	Cargo     string `json:"cargo" db:"cargo"`
	Ativo     bool   `json:"ativo" db:"ativo"`
}
