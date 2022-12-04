package utils

import (
	"fmt"
	"gaskn/config"
	"github.com/common-nighthawk/go-figure"
	"github.com/gofiber/fiber/v2"
)

func FiberConf() fiber.Config {
	cnf := fiber.Config{
		AppName:               config.Config("APP_NAME"),
		CaseSensitive:         true,
		DisableStartupMessage: true,
	}

	return cnf
}

func PrintBanner() {
	myFigure := figure.NewFigure(config.Config("APP_NAME"), "banner3-D", true)
	myFigure.Print()
	fmt.Println(config.Config("APP_BANER_DESC"))
	fmt.Println("")
}
