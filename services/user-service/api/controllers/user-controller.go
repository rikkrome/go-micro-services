package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rikkrome/go-micro-services/services/user-service/api/dto"
	"github.com/rikkrome/go-micro-services/services/user-service/api/models"
)

func CreateUserHandler(m *models.UserModel) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user dto.CreateUser

		// Decode the incoming User json
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// save to DB.
		err := m.Create(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Send response back
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]bool{"success": true})
	})
}

func GetUsersHandler(m *models.UserModel) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users, err := m.GetAllUsers()
		if err != nil {
			http.Error(w, "Could not fetch users", http.StatusInternalServerError)
			return
		}

		// Convert the users slice to JSON
		usersJSON, err := json.Marshal(users)
		if err != nil {
			http.Error(w, "Could not convert users to JSON", http.StatusInternalServerError)
			return
		}

		// Send the JSON response
		w.Header().Set("Content-Type", "application/json")
		w.Write(usersJSON)
	})
}

func DeleteUsersHandler(m *models.UserModel) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := m.DeleteAllUsers()
		if err != nil {
			http.Error(w, "Could not delete users", http.StatusInternalServerError)
			return
		}

		// Send the response
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "All users deleted successfully"}`))
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
