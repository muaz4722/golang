package utils

import (
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ValidateStruct(s interface{}) map[string]string {
	err := Validate.Struct(s)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		errors[e.Field()] = e.Tag()
	}
	return errors
}
