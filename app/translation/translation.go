package translation

import (
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/database/stores"
	"gorm.io/datatypes"
	"log"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/gofiber/fiber/v2"
)

var LangContext = config.Config("LANG")

var Trans *ut.UniversalTranslator

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

func SetupLang() {

	translator := en.New()
	Trans = ut.New(translator, translator, id.New())

	err := Trans.Import(ut.FormatJSON, config.Config("DIR_LANG"))
	if err != nil {
		log.Fatal(err)
	}

	err = Trans.VerifyTranslations()
	if err != nil {
		log.Fatal(err)
	}
}

func Lang(key string, params ...string) string {
	var lng ut.Translator

	var locale = FilterParamContext(LangContext, "en", "id")

	lng, _ = Trans.GetTranslator(locale)

	parseLang, _ := lng.T(key, params...)

	return parseLang
}

func LangFromJsonParse(attribute datatypes.JSONType[stores.LangAttribute]) string {

	if LangContext == "en" {
		return attribute.Data.En
	}

	if LangContext == "id" {
		return attribute.Data.Id
	}

	return attribute.Data.En
}

func LangMiddleware(ctx *fiber.Ctx) error {
	var lng = ctx.Query("lang")

	if lng == "en" || lng == "id" {
		LangContext = lng
	} else {
		LangContext = config.Config("lang")
	}

	return ctx.Next()
}
