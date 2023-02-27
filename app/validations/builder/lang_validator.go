package builder

import (
	"github.com/bonkzero404/gaskn/app/translation"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	englishTranslation "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
)

func SetLanguageValidator(ctx *fiber.Ctx, validate *validator.Validate) ut.Translator {
	var trans ut.Translator

	enTrans := en.New()
	idTrans := id.New()

	uni := ut.New(enTrans, enTrans, idTrans)

	var lng = translation.FilterParamContext(ctx.Query("lang"), "en", "id")
	trans, _ = uni.GetTranslator(lng)

	_ = englishTranslation.RegisterDefaultTranslations(validate, trans)

	return trans
}

func RegisterTagLanguageValidator(validate *validator.Validate, translator ut.Translator, tag string, f validator.Func, message string) {
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

func SetCustomTagLanguageValidator(ctx *fiber.Ctx, validate *validator.Validate, translator ut.Translator, tag string, f validator.Func) {
	var s string
	var lng = translation.FilterParamContext(ctx.Query("lang"), "en", "id")

	var t ut.Translator
	t, _ = translation.Trans.GetTranslator(lng)

	parseLang, _ := t.T(tag, s)

	RegisterTagLanguageValidator(validate, translator, tag, f, parseLang)

}
