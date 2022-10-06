package app

import (
	"go-starterkit-project/app/middleware"
	"go-starterkit-project/config"
	"go-starterkit-project/modules/auth"
	cl "go-starterkit-project/modules/client"
	"go-starterkit-project/modules/role"
	"go-starterkit-project/modules/user"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

/*
*
This function is used to register all modules,
this registration is the last process to register
all modules
*/
func Bootstrap(app *fiber.App) {
	// Monitor app
	app.Get("/monitor", monitor.New())

	// Get feature lists
	app.Get(
		utils.SetupApiGroup()+"/features",
		middleware.Authenticate(),
		func(c *fiber.Ctx) error {
			return utils.ApiOk(c, utils.ExtractRouteAsFeatures(c.App()))
		})

	// Get feature per group
	app.Get(
		utils.SetupApiGroup()+"/features/group",
		middleware.Authenticate(),
		func(c *fiber.Ctx) error {
			return utils.ApiOk(c, utils.FeaturesGroupLists(c.App()))
		})

	// Register module user
	user.RegisterModule(app)

	// Register module auth
	auth.RegisterModule(app)

	// Register module role
	role.RegisterModule(app)

	// Register Client
	cl.RegisterModule(app)
}

func SetupLogs() {
	if config.Config("ENABLE_LOG") == "true" {
		utils.CraeteDirectory(config.Config("LOG_LOCATION"))
	}
}
