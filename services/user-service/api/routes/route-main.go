package routes

import (
	"github.com/gorilla/mux"
)

func CombineRoutes(r *mux.Router) {
	RegisterUserRoutes(r)
}
