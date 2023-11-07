package test

import (
	"os"
	"testing"

	"github.com/rikkrome/go-micro-services/services/accounts-service/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	dsn := os.Getenv("SQL_DSN")
	accounts_db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrate the schema
	accounts_db.AutoMigrate(&models.Account{})

	// Run the tests
	code := m.Run()
	// You can also clean up the database here if needed

	os.Exit(code)

	// // Set up the router
	// r := mux.NewRouter()
	// r.HandleFunc("/login", LoginHandler).Methods("POST")

	// // Create a request to send to the above route
	// payload := map[string]string{"username": "testuser", "password": "testpass"}
	// payloadBytes, _ := json.Marshal(payload)
	// req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payloadBytes))
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// // Create a ResponseRecorder to record the response
	// rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(LoginHandler)

	// // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// // directly and pass in our Request and ResponseRecorder
	// handler.ServeHTTP(rr, req)

	// // Check the status code is what we expect
	// if status := rr.Code; status != http.StatusOK {
	// 	t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	// }

	// // Check the response body is what we expect
	// expected := `{"success":true}`
	// if rr.Body.String() != expected {
	// 	t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	// }
}
