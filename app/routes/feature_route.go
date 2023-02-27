package routes

import (
	"github.com/bonkzero404/gaskn/app/http/builder"
	middleware2 "github.com/bonkzero404/gaskn/app/http/middleware"
	"github.com/bonkzero404/gaskn/app/http/response"
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/config"
	"github.com/gofiber/fiber/v2"
)

func Feature(app *fiber.App) {
	var route = builder.RouteBuilder{}

	route.Set(app).Group(utils.ApiBasePath() + "/features").SetGroupName("Features")
	// Get feature lists
	route.Get(
		"/",
		middleware2.Authenticate(),
		middleware2.Permission(),
		func(c *fiber.Ctx) error {
			return response.ApiOk(c, builder.ExtractRouteAsFeatures(c, false))
		}).
		SetRouteName("FeatureLists").
		SetRouteDescriptionKeyLang("route:features").
		Execute()

	// Get feature per group
	route.Get(
		"/group",
		middleware2.Authenticate(),
		middleware2.Permission(),
		func(c *fiber.Ctx) error {
			return response.ApiOk(c, builder.FeaturesGroupLists(c, false))
		}).
		SetRouteName("FeatureGroupLists").
		SetRouteDescriptionKeyLang("route:features:group").
		Execute()

	if config.Config("TENANCY") == "true" {
		route.Set(app).Group(utils.ApiClientBasePath() + "/features").SetGroupName("Client/Features")

		// Get feature lists Tenant
		route.Get(
			"/",
			middleware2.Authenticate(),
			middleware2.Permission(),
			func(c *fiber.Ctx) error {
				return response.ApiOk(c, builder.ExtractRouteAsFeatures(c, true))
			}).
			SetRouteName("FeatureLists").
			SetRouteDescriptionKeyLang("route:features").
			SetRouteTenant(true).
			Execute()

		// Get feature per group tenant
		route.Get(
			"/group",
			middleware2.Authenticate(),
			middleware2.Permission(),
			func(c *fiber.Ctx) error {
				return response.ApiOk(c, builder.FeaturesGroupLists(c, true))
			}).
			SetRouteName("FeatureGroupLists").
			SetRouteDescriptionKeyLang("route:features:group").
			SetRouteTenant(true).
			Execute()
	}
}
