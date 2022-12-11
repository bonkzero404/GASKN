package app

import (
	"gaskn/app/middleware"
	"gaskn/config"
	"gaskn/features/auth"
	cl "gaskn/features/client"
	"gaskn/features/role"
	"gaskn/features/role_assignment"
	"gaskn/features/user"
	"gaskn/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func extrasFeature(app *fiber.App) {
	var route = utils.GasknRouter{}

	route.Set(app).Group(utils.SetupApiGroup() + "/features").SetGroupName("Features")
	// Get feature lists
	route.Get(
		"/",
		middleware.Authenticate(),
		middleware.Permission(),
		func(c *fiber.Ctx) error {
			return utils.ApiOk(c, utils.ExtractRouteAsFeatures(c.App(), false))
		}).
		SetRouteName("FeatureLists").
		SetRouteDescription("Admin get route lists").
		Execute()

	// Get feature per group
	route.Get(
		"/group",
		middleware.Authenticate(),
		middleware.Permission(),
		func(c *fiber.Ctx) error {
			return utils.ApiOk(c, utils.FeaturesGroupLists(c.App(), false))
		}).
		SetRouteName("FeatureGroupLists").
		SetRouteDescription("Admin get get group route lists").
		Execute()

	if config.Config("TENANCY") == "true" {
		route.Set(app).Group(utils.SetupSubApiGroup() + "/features").SetGroupName("Client/Features")

		// Get feature lists Tenant
		route.Get(
			"/",
			middleware.Authenticate(),
			middleware.Permission(),
			func(c *fiber.Ctx) error {
				return utils.ApiOk(c, utils.ExtractRouteAsFeatures(c.App(), true))
			}).
			SetRouteName("FeatureLists").
			SetRouteDescription("Admin get get route lists").
			SetRouteTenant(true).
			Execute()

		// Get feature per group tenant
		route.Get(
			"/group",
			middleware.Authenticate(),
			middleware.Permission(),
			func(c *fiber.Ctx) error {
				return utils.ApiOk(c, utils.FeaturesGroupLists(c.App(), true))
			}).
			SetRouteName("FeatureGroupLists").
			SetRouteDescription("Admin get get group route lists").
			SetRouteTenant(true).
			Execute()
	}
}

// Bootstrap /*
func Bootstrap(app *fiber.App) {
	// Monitor app
	app.Get("/monitor", monitor.New())

	// Register features
	extrasFeature(app)

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
}

func SetupLogs() {
	if config.Config("ENABLE_LOG") == "true" {
		utils.CraeteDirectory(config.Config("LOG_LOCATION"))
	}
}
