package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MattyDroidX/hotel-ease-backend/api/db"
    "github.com/MattyDroidX/hotel-ease-backend/handlers"
    "github.com/MattyDroidX/hotel-ease-backend/models"
	"github.com/gorilla/mux"
)

func main() {
	db.Connect()

	r := mux.NewRouter()
	r.HandleFunc("/funcionario", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "API Funcionando 👌")
	}).Methods("GET")

	fmt.Println("🚀 Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
