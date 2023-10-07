package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rikkrome/go-micro-services/services/user-service/api/routes"
	"github.com/rikkrome/go-micro-services/services/user-service/configs"
)

func main() {

	log.SetPrefix("user-service: ")

	db, _ := configs.InitSQLDatabase()
	if db != nil {
		log.Print("main connection Successful passed")
	}
	// initialize mux router...
	r := mux.NewRouter()
	routes.CombineRoutes(r, db)
	//handle for api request...
	http.Handle("/", r)
	log.Print("server listining on port 8080")
	http.ListenAndServe("localhost:8080", r)

}
