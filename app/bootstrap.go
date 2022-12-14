package app

import (
	"github.com/bonkzero404/gaskn/app/middleware"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/features/auth"
	cl "github.com/bonkzero404/gaskn/features/client"
	"github.com/bonkzero404/gaskn/features/role"
	"github.com/bonkzero404/gaskn/features/role_assignment"
	"github.com/bonkzero404/gaskn/features/user"
	"github.com/bonkzero404/gaskn/utils"

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
			return utils.ApiOk(c, utils.ExtractRouteAsFeatures(c, false))
		}).
		SetRouteName("FeatureLists").
		SetRouteDescriptionKeyLang("route:features").
		Execute()

	// Get feature per group
	route.Get(
		"/group",
		middleware.Authenticate(),
		middleware.Permission(),
		func(c *fiber.Ctx) error {
			return utils.ApiOk(c, utils.FeaturesGroupLists(c, false))
		}).
		SetRouteName("FeatureGroupLists").
		SetRouteDescriptionKeyLang("route:features:group").
		Execute()

	if config.Config("TENANCY") == "true" {
		route.Set(app).Group(utils.SetupSubApiGroup() + "/features").SetGroupName("Client/Features")

		// Get feature lists Tenant
		route.Get(
			"/",
			middleware.Authenticate(),
			middleware.Permission(),
			func(c *fiber.Ctx) error {
				return utils.ApiOk(c, utils.ExtractRouteAsFeatures(c, true))
			}).
			SetRouteName("FeatureLists").
			SetRouteDescriptionKeyLang("route:features").
			SetRouteTenant(true).
			Execute()

		// Get feature per group tenant
		route.Get(
			"/group",
			middleware.Authenticate(),
			middleware.Permission(),
			func(c *fiber.Ctx) error {
				return utils.ApiOk(c, utils.FeaturesGroupLists(c, true))
			}).
			SetRouteName("FeatureGroupLists").
			SetRouteDescriptionKeyLang("route:features:group").
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
