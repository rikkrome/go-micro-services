package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.rikkrome/tokyo/services/user-service/api/routes"
)

func main() {

	log.SetPrefix("user: ")
	log.Print("...")

	// initialize mux router...
	r := mux.NewRouter()
	routes.CombineRoutes(r)

	//handle for api request...
	http.Handle("/", r)

	fmt.Println("server listining on port 8080")
	http.ListenAndServe("localhost:8080", r)

}
