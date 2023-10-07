package configs

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/rikkrome/go-micro-services/services/user-service/api/models"
)

func InitSQLDatabase() (db *gorm.DB, err error) {
	// dsn := "host=localhost user=postgres password=postgres dbname=user_service port=5432 sslmode=disable"
	dsn := "postgres://postgres:postgres@localhost:5432/user_service?sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	} else {
		log.Print("Connection Successful")
	}
	db.AutoMigrate(&models.User{})
	return db, err
}
