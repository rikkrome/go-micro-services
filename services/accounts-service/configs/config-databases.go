package configs

import (
	"log"
	"os"
	"time"

	"github.com/rikkrome/go-micro-services/services/accounts-service/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitSQLDatabase() (accounts_db *gorm.DB, err error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// connectionString := "mongodb://rootuser:rootpass@localhost:27017"
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	// if err != nil {
	// 	log.Fatalf("Error connecting to MongoDB: %v", err)
	// } else {
	// 	log.Print("Auth Connection Successful")
	// }

	// // Check the connection
	// err = client.Ping(ctx, nil)
	// if err != nil {
	// 	log.Fatalf("Error pinging Auth DB: %v", err)
	// }

	// auth_db = client.Database("auth_service")

	// user_serive postgres
	dsn := os.Getenv("SQL_DSN")
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Warn, // Log level
			Colorful:      false,       // Disable color
		},
	)

	accounts_db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("Could not connect to the users database:", err)
	} else {
		log.Print("Users Database Connection Successful")
	}
	accounts_db.AutoMigrate(&models.Account{})

	return accounts_db, err
}
