package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/rikkrome/go-micro-services/services/user-service/api/dto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	User struct {
		ID             uuid.UUID `gorm:"type:uuid;primary_key;"`
		FirstName      string    `gorm:"column:first_name"`
		LastName       string    `gorm:"column:last_name"`
		Username       string    `gorm:"column:username;unique"`
		Email          string    `gorm:"uniqueIndex,column:email;unique"`
		HashedPassword string    `gorm:"column:hashed_password"`
		CreatedAt      time.Time
	}

	UserModel struct {
		DB *gorm.DB
	}
)

// constructor func
func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{DB: db}
}

func (m *UserModel) Create(user *dto.CreateUser) error {
	// Validate the DTO
	validate := validator.New()
	err := validate.Struct(user)
	errorMessage := ""
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()
			errorMessage += fmt.Sprintf("%s is %s\n", field, tag)
		}
		return errors.New(errorMessage)
	}
	newUUID, err := uuid.NewRandom() // Generate a new UUID
	if err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		// handle error
		log.Printf("Error GenerateFromPassword: %v", err)
		return err
	}
	newUser := &User{
		ID:             newUUID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Username:       user.Username,
		HashedPassword: string(hashedPassword),
	}
	result := m.DB.Create(newUser)
	return result.Error
}

func (m *UserModel) GetAllUsers() ([]User, error) {
	var users []User
	result := m.DB.Find(&users)
	return users, result.Error
}

func (m *UserModel) DeleteAllUsers() error {
	// result := m.DB.Delete(&User{})
	result := m.DB.Exec("DELETE FROM users")
	if result.Error != nil {
		log.Printf("Error deleting all users: %v", result.Error)
	}
	return result.Error
}
