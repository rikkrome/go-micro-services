package models

import (
	"log"

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
	err := dto.Validate(user)
	if err != nil {
		return err
	}
	newUUID, err := uuid.NewRandom() // Generate a new UUID
	if err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	newUser := &User{
		ID:             newUUID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: string(hashedPassword),
	}
	result := m.DB.Create(newUser)
	return result.Error
}

func (m *UserModel) GetAllUsers() ([]dto.User, error) {
	var users []dto.User
	result := m.DB.Find(&users)
	return users, result.Error
}

func (m *UserModel) GetUserById(id string) (dto.User, error) {
	var user dto.User
	result := m.DB.First(&user, "id = ?", id)
	return user, result.Error
}

func (m *UserModel) DeleteAllUsers() error {
	// result := m.DB.Delete(&User{})
	result := m.DB.Exec("DELETE FROM users")
	if result.Error != nil {
		log.Printf("Error deleting all users: %v", result.Error)
	}
	return result.Error
}
