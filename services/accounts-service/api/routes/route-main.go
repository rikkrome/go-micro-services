package routes

import (
	"github.com/gorilla/mux"
	"github.com/rikkrome/go-micro-services/services/accounts-service/api/models"
	"gorm.io/gorm"
)

func CombineRoutes(r *mux.Router, accounts_db *gorm.DB) {
	var accountsModel = models.NewAccountModel(accounts_db)
	RegisterAccountRoutes(r, accountsModel)
}
