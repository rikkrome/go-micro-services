package databases

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitSQLDatabase() error {
	log.Print("InitSQLDatabase...")
	// dsn := "host=localhost user=postgres password=postgres dbname=user_service port=5432 sslmode=disable"
	dsn := "postgres://postgres:postgres@localhost:5432/user_service?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Connecting DB")
		return err
	} else {
		log.Print("Connection Successful")
		DB = db
	}
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
