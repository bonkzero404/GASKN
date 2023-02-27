package utils

import (
	"fmt"
	"github.com/bonkzero404/gaskn/config"
	"github.com/common-nighthawk/go-figure"
	"github.com/gofiber/fiber/v2"
)

func FiberConf() fiber.Config {
	cnf := fiber.Config{
		AppName:               config.Config("APP_NAME"),
		CaseSensitive:         true,
		DisableStartupMessage: true,
		// EnablePrintRoutes:     true,
	}

	return cnf
}

func PrintBanner() {
	myFigure := figure.NewFigure(config.Config("APP_NAME"), "banner3-D", true)
	myFigure.Print()
	fmt.Println(config.Config("APP_BANER_DESC"))
}
