package dto

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
)

type ErrorMessages map[string]map[string]string

func loadErrorMessages(filename string) (ErrorMessages, error) {
	data, err := os.ReadFile(filename) // Use os.ReadFile instead of ioutil.ReadFile
	if err != nil {
		return nil, err
	}

	var messages ErrorMessages
	err = json.Unmarshal(data, &messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

type ValidationError struct {
	Field string
	Type  string
}

func Validate(dto interface{}) error {
	validate := validator.New()
	err := validate.Struct(dto)

	if err == nil {
		return nil
	}

	messages, errr := loadErrorMessages("configs/error_messages.json")
	if errr != nil {
		log.Fatalf("Could not load error messages: %v", errr)
	}

	errorMessage := ""
	for _, err := range err.(validator.ValidationErrors) {
		ve := &ValidationError{
			Field: err.Field(),
			Type:  err.Tag(),
		}
		if fieldMessages, ok := messages[ve.Field]; ok {
			if msg, ok := fieldMessages[ve.Type]; ok {
				errorMessage += fmt.Sprintf("%s\n", msg)
			}
		}
	}
	return errors.New(errorMessage)
}
