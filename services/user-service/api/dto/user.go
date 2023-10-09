package dto

type (
	// create new user
	CreateUser struct {
		FirstName string `json:"firstname" validate:"required,min=2"`
		LastName  string `json:"lastname" validate:"required,min=2"`
		Username  string `json:"username" validate:"required,min=2"`
		Email     string `json:"email" validate:"required,min=2,email"`
		Password  string `json:"password" validate:"required,min=8"`
	}

	Users struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Username  string `json:"username"`
	}
)
