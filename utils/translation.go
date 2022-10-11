package utils

import (
	"gaskn/config"
	"log"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/gofiber/fiber/v2"
)

var Utrans *ut.UniversalTranslator

func SetupLang() {

	en := en.New()
	Utrans = ut.New(en, en, id.New())

	err := Utrans.Import(ut.FormatJSON, config.Config("DIR_LANG"))
	if err != nil {
		log.Fatal(err)
	}

	err = Utrans.VerifyTranslations()
	if err != nil {
		log.Fatal(err)
	}
}

func Lang(ctx *fiber.Ctx, key string, params ...string) string {
	var lng ut.Translator

	if ctx.Query("lang") != "" && ctx.Query("lang") == "en" {
		lng, _ = Utrans.GetTranslator("en")
	} else if ctx.Query("lang") != "" && ctx.Query("lang") == "id" {
		lng, _ = Utrans.GetTranslator("id")
	} else {
		lng, _ = Utrans.GetTranslator("en")
	}

	parseLang, _ := lng.T(key, params...)

	return parseLang
}
