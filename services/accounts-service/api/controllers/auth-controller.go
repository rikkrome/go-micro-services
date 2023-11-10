package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
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
			account dto.EmailLogin
			err     error
		)
		if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		newAccount, err := m.GetAccountByEmail(&account)
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
		cookie := &http.Cookie{
			Name:     "accessToken",
			Value:    tokens.AccessToken,
			Path:     "/",
			HttpOnly: true, // HttpOnly helps mitigate the risk of client side script accessing the protected cookie.
			Secure:   true, // Secure ensures the browser sends the cookie only over HTTPS.
		}
		http.SetCookie(w, cookie)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(tokenJSON)
	})
}

func MineHandler(m *models.AccountModel) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			account dto.Account
		)
		authorizationHeader := r.Header.Get("Authorization")
		parts := strings.Split(authorizationHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "Authorization header format must be Bearer {token}", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]

		base64Key := os.Getenv("SECRET_AUTH_KEY")
		// Decode the Base64 encoded key
		key, err := base64.StdEncoding.DecodeString(base64Key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// Now parse the token TokenClaims
		parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
		if err != nil {
			fmt.Println("Error parsing token: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
			if accountIdValue, ok := claims["id"]; ok {
				if accountId, ok := accountIdValue.(string); ok {
					account, err = m.GetMine(&accountId)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
				}
			}
		} else {
			http.Error(w, "Invalid token", http.StatusInternalServerError)
		}

		accountJSON, err := json.Marshal(account)
		if err != nil {
			http.Error(w, "Could not convert user to JSON", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(accountJSON)
	})
}

func DeleteMineHandler(m *models.AccountModel) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		parts := strings.Split(authorizationHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "Authorization header format must be Bearer {token}", http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]

		base64Key := os.Getenv("SECRET_AUTH_KEY")
		// Decode the Base64 encoded key
		key, err := base64.StdEncoding.DecodeString(base64Key)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// Now parse the token TokenClaims
		parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})
		if err != nil {
			fmt.Println("Error parsing token: ", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
			if accountIdValue, ok := claims["id"]; ok {
				if accountId, ok := accountIdValue.(string); ok {
					err = m.DeleteAccount(&accountId)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
				}
			}
		} else {
			http.Error(w, "Invalid token", http.StatusInternalServerError)
		}
		response := dto.SuccessResponse{Success: true}
		responseJSON, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Could not convert user to JSON", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	})
}
