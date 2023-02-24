package utils

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

var Trans *ut.UniversalTranslator

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

func Lang(ctx *fiber.Ctx, key string, params ...string) string {
	var lng ut.Translator

	var locale = FilterParamContext(ctx.Query("lang"), "en", "id")

	lng, _ = Trans.GetTranslator(locale)

	parseLang, _ := lng.T(key, params...)

	return parseLang
}

func LangFromJsonParse(ctx *fiber.Ctx, attribute datatypes.JSONType[stores.LangAttribute]) string {
	var lng = ctx.Query("lang")

	if lng == "en" {
		return attribute.Data.En
	}

	if lng == "id" {
		return attribute.Data.Id
	}

	return attribute.Data.En
}
