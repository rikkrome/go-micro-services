package routes

import (
	"github.com/gorilla/mux"
	"github.com/rikkrome/go-micro-services/services/user-service/api/controllers"
	"github.com/rikkrome/go-micro-services/services/user-service/api/models"
)

// this function will handle all user routes...
// takes a pointer of *mux.Router
func RegisterUserRoutes(router *mux.Router, userModel *models.UserModel) {
	router.HandleFunc("/users/", controllers.CreateUserHandler(userModel)).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.GetUserHandler(userModel)).Methods("GET")
	router.HandleFunc("/users/all", controllers.GetUsersHandler(userModel)).Methods("GET")
	router.HandleFunc("/users/all", controllers.DeleteUsersHandler(userModel)).Methods("DELETE")
	router.HandleFunc("/users/health", controllers.HealthCheckHandler).Methods("GET")
}
