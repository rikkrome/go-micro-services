package routes

import (
	"github.com/gorilla/mux"
	"github.com/rikkrome/go-micro-services/services/user-service/api/controllers"
	"github.com/rikkrome/go-micro-services/services/user-service/api/models"
)

// this function will handle all user routes...
// takes a pointer of *mux.Router
func RegisterUserRoutes(router *mux.Router, userModel *models.UserModel) {
	router.HandleFunc("/user/", controllers.CreateUserHandler(userModel)).Methods("POST")
	router.HandleFunc("/user/health", controllers.HealthCheckHandler).Methods("GET")
}
