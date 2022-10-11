package utils

import (
	"go-starterkit-project/dto"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	id_translations "github.com/go-playground/validator/v10/translations/id"
	"github.com/gofiber/fiber/v2"
)

func ValidateStruct(s interface{}, ctx *fiber.Ctx) []*dto.ErrorResponse {
	var errors []*dto.ErrorResponse
	var validate *validator.Validate
	var trans ut.Translator

	en := en.New()
	id := id.New()

	if ctx.Query("lang") != "" && ctx.Query("lang") == "en" {
		uni := ut.New(en, en)
		trans, _ = uni.GetTranslator("en")
		validate = validator.New()

		en_translations.RegisterDefaultTranslations(validate, trans)
	} else if ctx.Query("lang") != "" && ctx.Query("lang") == "id" {
		uni := ut.New(id, id)
		trans, _ = uni.GetTranslator("id")
		validate = validator.New()

		id_translations.RegisterDefaultTranslations(validate, trans)
	} else {
		uni := ut.New(en, en)
		trans, _ = uni.GetTranslator("en")
		validate = validator.New()

		en_translations.RegisterDefaultTranslations(validate, trans)
	}

	err := validate.Struct(s)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element dto.ErrorResponse

			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Message = err.Translate(trans)
			errors = append(errors, &element)
		}
	}
	return errors
}
