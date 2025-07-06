package db

import (
	"fmt"

	"github.com/jmoiron/sqlx" // ← sqlx
)

/*
   Cria/valida as tabelas principais.

   • data_hora como TIMESTAMPTZ
   • chave estrangeira de tarefas → funcionarios
   • UNIQUE em e-mail e DEFAULT em ativo
*/
func CreateTables() {
	const funcionarios = `
CREATE TABLE IF NOT EXISTS funcionarios (
  id         UUID        PRIMARY KEY,
  nome       TEXT        NOT NULL,
  sobrenome  TEXT        NOT NULL,
  email      TEXT        NOT NULL UNIQUE,
  telefone   TEXT        NOT NULL,
  cargo      TEXT        NOT NULL,
  ativo      BOOLEAN     NOT NULL DEFAULT true
);`

	const tarefas = `
CREATE TABLE IF NOT EXISTS tarefas (
  id            UUID         PRIMARY KEY,
  numero        TEXT         NOT NULL,
  funcionario   UUID         NOT NULL
                 REFERENCES funcionarios(id)
                 ON DELETE RESTRICT,
  descricao     TEXT         NOT NULL,
  data_hora     TIMESTAMPTZ  NOT NULL,
  status        TEXT         NOT NULL,
  tipo          TEXT         NOT NULL
);`

	if err := execMany(DB, funcionarios, tarefas); err != nil {
		fmt.Println("❌ Erro ao criar tabelas:", err)
	} else {
		fmt.Println("✅ Tabelas criadas/verificadas com sucesso")
	}
}

// execMany roda todas as queries numa transação.
func execMany(db *sqlx.DB, queries ...string) error {
	tx, err := db.Beginx() // Beginx() devolve *sqlx.Tx
	if err != nil {
		return err
	}
	for _, q := range queries {
		if _, err := tx.Exec(q); err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
