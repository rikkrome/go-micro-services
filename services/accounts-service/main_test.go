package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rikkrome/go-micro-services/services/accounts-service/api/controllers"
	"github.com/rikkrome/go-micro-services/services/accounts-service/api/dto"
	"github.com/rikkrome/go-micro-services/services/accounts-service/api/models"
	"github.com/rikkrome/go-micro-services/services/accounts-service/configs"
	"gorm.io/gorm"
)

var accounts_db *gorm.DB
var tokens *dto.AuthTokens
var accountsModel *models.AccountModel

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := configs.InitSQLDatabase()
	if err != nil {
		return
	}
	accounts_db = db
	accountsModel = models.NewAccountModel(accounts_db)
	// router := mux.NewRouter()
	// routes.CombineRoutes(router, accounts_db)

	// Run the tests
	code := m.Run()

	os.Exit(code)
}

func TestSignup(t *testing.T) {
	// Create a request to send to the above route
	payload := map[string]string{
		"firstname": "test",
		"lastname":  "user",
		"username":  "testuser",
		"email":     "test@romeroricky.com",
		"password":  "123456789",
	}
	payloadBytes, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "/accounts/signup", bytes.NewBuffer(payloadBytes))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.SignUpHandler(accountsModel))

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder
	handler.ServeHTTP(responseRecorder, req)

	// Check the status code is what we expect
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// var tokens dto.AuthTokens
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &tokens)
	if err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	// Assert that the ID field is set and not empty (as a basic test of the response)
	if tokens.ID == "" {
		t.Errorf("handler returned empty ID field")
	}

	// Similarly, check if the AccessToken and RefreshToken are not empty.
	if tokens.AccessToken == "" || tokens.RefreshToken == "" {
		t.Errorf("handler returned empty tokens")
	}
}

func TestDeleteAccount(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/accounts/delete", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+tokens.AccessToken)

	// Create a ResponseRecorder to record the response
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.DeleteMineHandler(accountsModel))

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder
	handler.ServeHTTP(responseRecorder, req)

	// Check the status code is what we expect
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response dto.SuccessResponse
	err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	if response.Success == false {
		t.Errorf("handler returned Success false")
	}
}
