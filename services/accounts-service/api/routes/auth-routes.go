package routes

import (
	"github.com/gorilla/mux"
	"github.com/rikkrome/go-micro-services/services/accounts-service/api/controllers"
	"github.com/rikkrome/go-micro-services/services/accounts-service/api/models"
)

// this function will handle all user routes...
func RegisterAccountRoutes(router *mux.Router, accountsModel *models.AccountModel) {
	router.HandleFunc("/accounts/signup", controllers.SignUpHandler(accountsModel)).Methods("POST")
	router.HandleFunc("/accounts/login", controllers.LoginHandler(accountsModel)).Methods("POST")
	router.HandleFunc("/accounts/mine", controllers.MineHandler(accountsModel)).Methods("GET")
	router.HandleFunc("/accounts/delete", controllers.DeleteMineHandler(accountsModel)).Methods("DELETE")
}
