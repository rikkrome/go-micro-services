package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rikkrome/go-micro-services/services/user-service/api/dto"
	"github.com/rikkrome/go-micro-services/services/user-service/api/models"
)

// var (
// 	UserModel models.UserModel = models.NewUserModel()
// )

// func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
func CreateUserHandler(m *models.UserModel) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user dto.User

		// Decode the incoming User json
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Validate the DTO (you can add more complex validation here)
		if user.FirstName == "" || user.LastName == "" || user.Username == "" {
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return
		}

		// save to DB.
		err := m.Create(&user)
		if err != nil {
			http.Error(w, "Error with DB", http.StatusBadRequest)
			return
		}
		// Send response back
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]bool{"success": true})
	})
}

/*
HealthCheckHandler
*/
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"alive": true})

}
