package utils

import (
	"github.com/go-playground/validator"
)

func ValidateInput(data interface{}) error {
	validate := validator.New()

	if err := validate.Struct(data); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}
	return nil
}
