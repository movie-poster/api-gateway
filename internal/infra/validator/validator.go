package validator

import (
	"strings"

	"api-gateway/internal/translation"

	"github.com/go-playground/validator/v10"
)

type Error struct {
	Key     string
	Message string
}

func GenerateMessage(v *validator.Validate, errs validator.ValidationErrors) []Error {
	trans := *translation.TranslationValidation(v)
	var errors []Error

	for _, e := range errs {
		error := Error{
			Key:     strings.Split(e.Translate(trans), " ")[0],
			Message: e.Translate((trans)),
		}
		errors = append(errors, error)
	}

	return errors
}
