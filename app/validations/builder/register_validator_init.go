package builder

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ValidationItems struct {
	Tag               string
	ValidatorFunction validator.Func
}

type Validators struct {
	Validator []ValidationItems
}

func (v *Validators) RegisterValidatorInit(ctx *fiber.Ctx, validate *validator.Validate, translator ut.Translator) {
	for _, validator := range v.Validator {
		SetCustomTagLanguageValidator(ctx, validate, translator, validator.Tag, validator.ValidatorFunction)
	}
}
