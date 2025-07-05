package db

import "fmt"

func CreateTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS funcionarios (
			id UUID PRIMARY KEY,
			nome TEXT NOT NULL,
			sobrenome TEXT NOT NULL,
			email TEXT NOT NULL,
			telefone TEXT NOT NULL,
			cargo TEXT NOT NULL,
			ativo BOOLEAN NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS tarefas (
			id UUID PRIMARY KEY,
			numero TEXT NOT NULL,
			funcionario TEXT NOT NULL,
			descricao TEXT NOT NULL,
			data_hora TEXT NOT NULL,
			status TEXT NOT NULL,
			tipo TEXT NOT NULL
		)`,
	}

	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			fmt.Println("❌ Erro ao criar tabela:", err)
		}
	}

	fmt.Println("✅ Tabelas criadas/verificadas com sucesso")
}
