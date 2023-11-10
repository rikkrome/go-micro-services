package dto

import "github.com/google/uuid"

type (
	SuccessResponse struct {
		Success bool `json:"success"`
	}
	EmailSignUp struct {
		FirstName string `json:"firstname" validate:"required"`
		LastName  string `json:"lastname" validate:"required"`
		Username  string `json:"username" validate:"required,min=3"`
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required,min=8"`
	}
	EmailLogin struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}
	Account struct {
		ID        uuid.UUID `json:"id"`
		FirstName string    `json:"firstname"`
		LastName  string    `json:"lastname"`
		Username  string    `json:"username"`
		Email     string    `json:"email"`
	}
	AuthTokens struct {
		ID                 string `json:"id" validate:"required"`
		AccessToken        string `json:"accessToken" validate:"required"`
		AccessTokenExpiry  int64  `json:"accessTokenExpiry" validate:"required"`
		RefreshToken       string `json:"refreshToken" validate:"required"`
		RefreshTokenExpiry int64  `json:"refreshTokenExpiry" validate:"required"`
	}
)
