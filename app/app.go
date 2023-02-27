package app

import (
	"github.com/bonkzero404/gaskn/app/routes"
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/features/auth"
	cl "github.com/bonkzero404/gaskn/features/client"
	"github.com/bonkzero404/gaskn/features/menu"
	"github.com/bonkzero404/gaskn/features/role"
	"github.com/bonkzero404/gaskn/features/role_assignment"
	"github.com/bonkzero404/gaskn/features/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

// RouteInit /*
func RouteInit(app *fiber.App) {

	// Monitor app
	app.Get("/monitor", monitor.New())

	// Register features
	routes.Feature(app)

	// Register feature user
	user.RegisterFeature(app)

	// Register feature auth
	auth.RegisterFeature(app)

	// Register feature role
	role.RegisterFeature(app)

	// Register Client
	cl.RegisterFeature(app)

	// Register feature Role Assignment
	role_assignment.RegisterFeature(app)

	// Register feature menu
	menu.RegisterFeature(app)
}

func SetupLogs() {
	if config.Config("ENABLE_LOG") == "true" {
		utils.CraeteDirectory(config.Config("LOG_LOCATION"))
	}
}
