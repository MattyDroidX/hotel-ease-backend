package utils

import (
	"errors"
	"github.com/MattyDroidX/hotel-ease-backend/api/db"
)

func FuncionarioIDByNome(nome string) (string, error) {
	var id string
	q := `SELECT id FROM funcionarios
	      WHERE CONCAT(nome,' ',sobrenome) = $1 LIMIT 1`
	if err := db.DB.Get(&id, q, nome); err != nil {
		return "", errors.New("funcionário não encontrado")
	}
	return id, nil
}

func NomeCompletoByID(id string) (string, error) {
	var nome string
	q := `SELECT CONCAT(nome,' ',sobrenome)
	      FROM funcionarios WHERE id=$1`
	if err := db.DB.Get(&nome, q, id); err != nil {
		return "", err
	}
	return nome, nil
}
