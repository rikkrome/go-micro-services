package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.rikkrome/tokyo/services/auth-service/api/routes"
)

func main() {

	log.SetPrefix("auth: ")
	log.Print("...")
	r := mux.NewRouter()
	routes.CombineRoutes(r)

	http.Handle("/", r)
	log.Print("server listining on port 8080")
	http.ListenAndServe("localhost:8080", r)
}
