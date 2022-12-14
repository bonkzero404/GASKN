package utils

import (
	"gaskn/dto"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	englishTranslation "github.com/go-playground/validator/v10/translations/en"
	indoTranslation "github.com/go-playground/validator/v10/translations/id"
	"github.com/gofiber/fiber/v2"

	"gaskn/config"
)

func ValidateStruct(s interface{}, ctx *fiber.Ctx) []*dto.ErrorResponse {
	var errors []*dto.ErrorResponse
	var validate *validator.Validate
	var trans ut.Translator

	enTrans := en.New()
	idTrans := id.New()

	if ctx.Query("lang") != "" && ctx.Query("lang") == "en" {
		uni := ut.New(enTrans, enTrans)
		trans, _ = uni.GetTranslator("en")
		validate = validator.New()

		err := englishTranslation.RegisterDefaultTranslations(validate, trans)
		if err != nil {
			return nil
		}
	} else if ctx.Query("lang") != "" && ctx.Query("lang") == "id" {
		uni := ut.New(idTrans, idTrans)
		trans, _ = uni.GetTranslator("id")
		validate = validator.New()

		err := indoTranslation.RegisterDefaultTranslations(validate, trans)
		if err != nil {
			return nil
		}
	} else {
		var defaultLang = "en"

		if config.Config("LANG") != "" {
			defaultLang = config.Config("LANG")
		}
		uni := ut.New(enTrans, enTrans)
		trans, _ = uni.GetTranslator(defaultLang)
		validate = validator.New()

		err := englishTranslation.RegisterDefaultTranslations(validate, trans)
		if err != nil {
			return nil
		}
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
