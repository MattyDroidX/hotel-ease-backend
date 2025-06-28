package models

type Funcionario struct {
	ID        string `json:"id"`
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Email     string `json:"email"`
	Telefone  string `json:"telefone"`
	Cargo     string `json:"cargo"`
	Ativo     bool   `json:"ativo"`
}
