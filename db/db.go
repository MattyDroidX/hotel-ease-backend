package db

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Connect() {
	var err error
	DB, err = sqlx.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}
}
