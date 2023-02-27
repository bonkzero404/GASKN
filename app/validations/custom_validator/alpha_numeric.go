package custom_validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidateAlphanumExtra(val validator.FieldLevel) bool {
	isAlphaNumCustom := regexp.MustCompile(`^[-_' a-zA-Z0-9]+$`).MatchString(val.Field().String())

	return isAlphaNumCustom
}
