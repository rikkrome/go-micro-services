package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.rikkrome/go-micro-services/services/user-service/api/controllers"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

// this function will handle all user routes...
// takes a pointer of *mux.Router
func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/user/", handler).Methods("POST")
	router.HandleFunc("/user/health", controllers.HealthCheckHandler).Methods("GET")
}
