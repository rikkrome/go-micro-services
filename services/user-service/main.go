package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.rikkrome/go-micro-services/services/user-service/api/routes"
	"go.rikkrome/go-micro-services/services/user-service/configs"
)

func main() {

	log.SetPrefix("user-service: ")

	configs.LoadConfigs()
	// initialize mux router...
	r := mux.NewRouter()
	routes.CombineRoutes(r)

	//handle for api request...
	http.Handle("/", r)

	log.Print("server listining on port 8080")
	http.ListenAndServe("localhost:8080", r)

}
