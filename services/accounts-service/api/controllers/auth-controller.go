package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rikkrome/go-micro-services/services/accounts-service/api/dto"
	"github.com/rikkrome/go-micro-services/services/accounts-service/api/models"
)

func SignUpHandler(m *models.AccountModel) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			user dto.EmailSignUp
			err  error
		)
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		newAccount, err := m.Create(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newTokens, err := m.NewToken(newAccount.ID.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tokens := &dto.AuthTokens{
			ID:                 newAccount.ID.String(),
			AccessToken:        newTokens.AccessToken,
			AccessTokenExpiry:  newTokens.AccessTokenExpiry,
			RefreshToken:       newTokens.RefreshToken,
			RefreshTokenExpiry: newTokens.RefreshTokenExpiry,
		}
		tokenJSON, err := json.Marshal(tokens)
		if err != nil {
			http.Error(w, "Could not convert user to JSON", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(tokenJSON)
	})
}

func LoginHandler(m *models.AccountModel) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			user dto.EmailLogin
			err  error
		)
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		newAccount, err := m.GetAccountByEmail(&user)
		if err != nil {
			http.Error(w, "Invalid login credentials. Please try again.", http.StatusUnauthorized)
			return
		}
		newTokens, err := m.NewToken(newAccount.ID.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tokens := &dto.AuthTokens{
			ID:                 newAccount.ID.String(),
			AccessToken:        newTokens.AccessToken,
			AccessTokenExpiry:  newTokens.AccessTokenExpiry,
			RefreshToken:       newTokens.RefreshToken,
			RefreshTokenExpiry: newTokens.RefreshTokenExpiry,
		}
		tokenJSON, err := json.Marshal(tokens)
		if err != nil {
			http.Error(w, "Could not convert user to JSON", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(tokenJSON)
	})
}
