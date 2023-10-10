package dto

type (
	// create new user
	CreateUser struct {
		FirstName string `json:"firstname" validate:"required"`
		LastName  string `json:"lastname" validate:"required"`
		Username  string `json:"username" validate:"required,min=3"`
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required,min=8"`
	}

	User struct {
		ID        string `json:"id"`
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		Username  string `json:"username"`
		Email     string `json:"email"`
	}
)
