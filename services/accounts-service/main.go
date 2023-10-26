package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
	"github.com/rikkrome/go-micro-services/services/accounts-service/api/routes"
	"github.com/rikkrome/go-micro-services/services/accounts-service/configs"
)

func main() {

	log.SetPrefix("accounts: ")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accounts_db, err := configs.InitSQLDatabase()
	if err != nil {
		return
	}

	r := mux.NewRouter()
	routes.CombineRoutes(r, accounts_db)

	http.Handle("/", r)
	log.Print("server listining on port 8080")
	http.ListenAndServe("localhost:8080", r)
}
