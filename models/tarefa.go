package models

type Tarefa struct {
	ID          string `json:"id" db:"id"`
	Numero      string `json:"numero" db:"numero"`
	Funcionario string `json:"funcionario" db:"funcionario"`
	Descricao   string `json:"descricao" db:"descricao"`
	DataHora    string `json:"dataHora" db:"data_hora"`
	Status      string `json:"status" db:"status"`
	Tipo        string `json:"tipo" db:"tipo"`
}
