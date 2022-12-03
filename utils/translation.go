package utils

import (
	"gaskn/config"
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

	if ctx.Query("lang") != "" && ctx.Query("lang") == "en" {
		lng, _ = Trans.GetTranslator("en")
	} else if ctx.Query("lang") != "" && ctx.Query("lang") == "id" {
		lng, _ = Trans.GetTranslator("id")
	} else {
		lng, _ = Trans.GetTranslator("en")
	}

	parseLang, _ := lng.T(key, params...)

	return parseLang
}
