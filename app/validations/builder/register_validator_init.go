package builder

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type ValidationItems struct {
	Tag               string
	ValidatorFunction validator.Func
}

type Validators struct {
	Validator []ValidationItems
}

func (v *Validators) RegisterValidatorInit(validate *validator.Validate, translator ut.Translator) {
	for _, validatorItem := range v.Validator {
		SetCustomTagLanguageValidator(validate, translator, validatorItem.Tag, validatorItem.ValidatorFunction)
	}
}
