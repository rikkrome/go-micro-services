package controllers

import (
	"encoding/json"
	"net/http"
)

// Model
type User struct {
	Name     string
	Username string
}

type (
	// For Post/Put - /users
	UserResource struct {
		Data User `json:"data"`
	}
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource

	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	user := &dataResource.Data

	if err != nil {
		panic(err)
	}

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"alive": true})

}
