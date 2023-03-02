package registry

import (
	"github.com/bonkzero404/gaskn/app/validations/builder"
	"github.com/bonkzero404/gaskn/app/validations/custom_validator"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func RegisterValidator(validate *validator.Validate, translator ut.Translator) {
	// Add your custom validation rule
	validators := builder.Validators{
		Validator: []builder.ValidationItems{
			{
				Tag:               "alphanum_extra",
				ValidatorFunction: custom_validator.ValidateAlphanumExtra,
			},
		},
	}

	validators.RegisterValidatorInit(validate, translator)
}
