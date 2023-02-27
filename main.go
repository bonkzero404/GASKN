package main

import (
	"fmt"
	appRoute "github.com/bonkzero404/gaskn/app"
	"github.com/bonkzero404/gaskn/app/translation"
	utils2 "github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/database"
	"github.com/bonkzero404/gaskn/database/seeders"
	driver2 "github.com/bonkzero404/gaskn/infrastructures"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/helmet/v2"
)

func main() {
	// Fiber app
	app := fiber.New(utils2.FiberConf())

	translation.SetupLang()

	// Setup Logs
	appRoute.SetupLogs()

	// Call database connection
	driver2.ConnectDB()

	// Auto migration table
	database.MigrateDB()

	// Init Casbin
	driver2.InitCasbin()

	// Handling global cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: config.Config("ALLOW_ORIGINS"),
	}))

	// Securing with helmet
	app.Use(helmet.New())

	// Etag
	app.Use(etag.New(etag.Config{
		Weak: true,
	}))

	// Call all module
	appRoute.RouteInit(app)

	// Run Seeder
	for _, seed := range seeders.All() {
		if err := seed.Run(driver2.DB); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}

	// Set port
	appPort := fmt.Sprintf("%s:%s", config.Config("APP_HOST"), config.Config("APP_PORT"))

	// Print banner
	utils2.PrintBanner()

	// Listen app
	err := app.Listen(appPort)
	if err != nil {
		panic(err)
	}
}
