package utils

import (
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/dto"
	"regexp"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	englishTranslation "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
)

func FilterParamContext(val string, locales ...string) string {
	for _, v := range locales {
		if v == val {
			return v
		}
	}

	if config.Config("LANG") != "" {
		return config.Config("LANG")
	}

	return "en"
}

func setLanguage(ctx *fiber.Ctx, validate *validator.Validate) ut.Translator {
	var trans ut.Translator

	enTrans := en.New()
	idTrans := id.New()

	uni := ut.New(enTrans, enTrans, idTrans)

	var lng = FilterParamContext(ctx.Query("lang"), "en", "id")
	trans, _ = uni.GetTranslator(lng)

	_ = englishTranslation.RegisterDefaultTranslations(validate, trans)

	return trans
}

func registerTagLanguage(
	validate *validator.Validate,
	translator ut.Translator,
	tag string,
	f validator.Func,
	message string,
) {
	_ = validate.RegisterValidation(tag, f)

	_ = validate.RegisterTranslation(tag, translator,
		func(ut ut.Translator) error {
			return ut.Add(tag, message, false)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			fld, _ := ut.T(fe.Field())
			t, err := ut.T(fe.Tag(), fld)
			if err != nil {
				return fe.(error).Error()
			}
			return t
		},
	)
}

func setCustomTagLanguage(
	ctx *fiber.Ctx,
	validate *validator.Validate,
	translator ut.Translator,
	tag string,
	f validator.Func,
) {
	var s string
	var lng = FilterParamContext(ctx.Query("lang"), "en", "id")

	var t ut.Translator
	t, _ = Trans.GetTranslator(lng)

	parseLang, _ := t.T(tag, s)

	registerTagLanguage(validate, translator, tag, f, parseLang)

}

func ValidateStruct(s any, ctx *fiber.Ctx) []*dto.ErrorResponse {
	var errors []*dto.ErrorResponse
	var validate = validator.New()

	var trans = setLanguage(ctx, validate)

	// Register custom tags
	setCustomTagLanguage(ctx, validate, trans, "alphanum_extra", ValidateAlphanumExtra)

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

func ValidateAlphanumExtra(val validator.FieldLevel) bool {
	isAlphaNumCustom := regexp.MustCompile(`^[-_' a-zA-Z0-9]+$`).MatchString(val.Field().String())

	return isAlphaNumCustom
}
