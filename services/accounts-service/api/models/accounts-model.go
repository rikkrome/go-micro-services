package models

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rikkrome/go-micro-services/services/accounts-service/api/dto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	AccountModel struct {
		DB *gorm.DB
	}
	Account struct {
		ID             uuid.UUID `gorm:"type:uuid;primary_key;"`
		FirstName      string    `gorm:"column:first_name"`
		LastName       string    `gorm:"column:last_name"`
		Username       string    `gorm:"column:username;unique"`
		Email          string    `gorm:"uniqueIndex,column:email;unique"`
		HashedPassword string    `gorm:"column:hashed_password"`
	}
	AuthTokens struct {
		AccessToken        string
		AccessTokenExpiry  int64
		RefreshToken       string
		RefreshTokenExpiry int64
	}
)

// constructor func
func NewAccountModel(db *gorm.DB) *AccountModel {
	return &AccountModel{DB: db}
}

func checkPassword(hashedPassword string, password string) error {
	// Compare the stored hashed password, with the hashed version of the password that was received
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		// If the two passwords don't match, return a 401 status
		return fmt.Errorf("incorrect password")
	}
	return nil
}

func (m *AccountModel) Create(accountData *dto.EmailSignUp) (account Account, err error) {
	err = dto.Validate(accountData)
	if err != nil {
		return Account{}, err
	}
	newUUID, err := uuid.NewRandom() // Generate a new UUID
	if err != nil {
		return Account{}, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(accountData.Password), bcrypt.DefaultCost)
	if err != nil {
		return Account{}, err
	}
	account = Account{
		ID:             newUUID,
		FirstName:      accountData.FirstName,
		LastName:       accountData.LastName,
		Username:       accountData.Username,
		Email:          accountData.Email,
		HashedPassword: string(hashedPassword),
	}
	var existingAccount Account
	if err := m.DB.Where("username = ? OR email = ?", account.Username, account.Email).First(&existingAccount).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Account doesn't exist, create new account
			if err := m.DB.Create(account).Error; err != nil {
				// Handle error creating account
				// fmt.Println("Error creating account:", err)
				return Account{}, err
			}
			// fmt.Println("Account created successfully")
			return account, nil
		}
	}
	err = errors.New("account already exists")
	return Account{}, err
}

func (m *AccountModel) GetAccountByEmail(accountData *dto.EmailLogin) (account Account, err error) {
	err = dto.Validate(accountData)
	if err != nil {
		return Account{}, err
	}
	result := m.DB.First(&account, "email = ?", &accountData.Email)
	err = checkPassword(account.HashedPassword, accountData.Password)
	if err != nil {
		return Account{}, err
	}
	return account, result.Error
}

func (m *AccountModel) GetMine(id *string) (dto.Account, error) {
	var account dto.Account
	result := m.DB.First(&account, "id = ?", &id)
	return account, result.Error
}

func (m *AccountModel) DeleteAccount(id *string) error {
	result := m.DB.Delete(&dto.Account{}, &id)
	return result.Error
}

func (m *AccountModel) GetAllAccounts() ([]dto.Account, error) {
	var accounts []dto.Account
	result := m.DB.Find(&accounts)
	return accounts, result.Error
}

func (m *AccountModel) DeleteAllAccounts() error {
	// result := m.DB.Delete(&User{})
	result := m.DB.Exec("DELETE FROM users")
	if result.Error != nil {
		log.Printf("Error deleting all users: %v", result.Error)
	}
	return result.Error
}

func (m *AccountModel) NewToken(id string) (tokens AuthTokens, err error) {
	// Create an array to hold the key
	base64Key := os.Getenv("SECRET_AUTH_KEY")
	// Decode the Base64 encoded key
	key, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		log.Fatal("Error decoding Base64 key: ", err)
		return AuthTokens{}, err
	}
	// Token
	tokenExp := time.Now().Add(time.Hour * 72).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": tokenExp,
	})
	tokenStr, err := token.SignedString(key)
	if err != nil {
		fmt.Println("Error SignedString:", err)
		return
	}
	// Refresh Token
	refreshTokenExp := time.Now().Add(time.Hour * 24 * 30).Unix()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": refreshTokenExp,
	})
	refreshTokenStr, err := refreshToken.SignedString(key)
	if err != nil {
		fmt.Println("Error SignedString:", err)
		return
	}
	tokens = AuthTokens{
		AccessToken:        tokenStr,
		AccessTokenExpiry:  tokenExp,
		RefreshToken:       refreshTokenStr,
		RefreshTokenExpiry: refreshTokenExp,
	}
	return tokens, nil
}
