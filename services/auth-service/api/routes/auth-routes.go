package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rikkrome/go-micro-services/services/auth-service/api/controllers"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("auth!\n"))
}

// this function will handle all user routes...
// takes a pointer of *mux.Router
func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/auth/", handler).Methods("POST")
	router.HandleFunc("/auth/health", controllers.HealthCheckHandler).Methods("GET")
}
