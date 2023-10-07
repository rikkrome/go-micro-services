package models

import (
	"github.com/rikkrome/go-micro-services/services/user-service/api/dto"
	"gorm.io/gorm"
)

type (
	User struct {
		// gorm.Model
		// FirstName string `json:"firstname"`
		// LastName  string `json:"lastname"`
		// Username  string `json:"username"`
		// Password  string `json:"password"`

		ID        uint   `gorm:"primaryKey"`
		FirstName string `gorm:"column:first_name"`
		LastName  string `gorm:"column:last_name"`
		Username  string `gorm:"column:username;unique"`
	}

	UserModel struct {
		DB *gorm.DB
	}
)

// constructor func
func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{DB: db}
}

func (m *UserModel) Create(user *dto.User) error {
	newUser := &User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
	}
	result := m.DB.Create(newUser)
	return result.Error
}
