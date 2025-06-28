package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error
	connStr := "user=hotelease_user password=SUA_SENHA dbname=hotelease sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Erro ao verificar conexão:", err)
	}

	fmt.Println("✅ Banco de dados conectado com sucesso")
}
