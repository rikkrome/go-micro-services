package routes

import (
	"github.com/gorilla/mux"
	"github.com/rikkrome/go-micro-services/services/user-service/api/models"
	"gorm.io/gorm"
)

func CombineRoutes(r *mux.Router, db *gorm.DB) {
	var userModel = models.NewUserModel(db)
	RegisterUserRoutes(r, userModel)
}
