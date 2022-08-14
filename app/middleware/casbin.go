package middleware

import (
	"go-starterkit-project/database/driver"
	"go-starterkit-project/domain/dto"
	"go-starterkit-project/utils"

	fibercasbin "github.com/arsmn/fiber-casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gofiber/fiber/v2"
)

func Casbin() *fibercasbin.CasbinMiddleware {
	db := driver.DB
	adapter, _ := gormadapter.NewAdapterByDB(db)
	authz := fibercasbin.New(fibercasbin.Config{
		Enforcer:      driver.Enforcer,
		ModelFilePath: "casbin_rbac_model.conf",
		PolicyAdapter: adapter,
		Unauthorized: func(ctx *fiber.Ctx) error {
			return utils.ApiUnauthorized(ctx, dto.Errors{
				Message: utils.Lang(ctx, "middleware:err:unauthorized"),
				Cause:   ctx.Next().Error(),
				Inputs:  nil,
			})
		},
		Forbidden: func(ctx *fiber.Ctx) error {
			return utils.ApiUnauthorized(ctx, dto.Errors{
				Message: utils.Lang(ctx, "Forbidden"),
				Cause:   ctx.Next().Error(),
				Inputs:  nil,
			})
		},
		Lookup: func(ctx *fiber.Ctx) string {
			// get subject from BasicAuth, JWT, Cookie etc in real world
			return "ok"
		},
	})
	return authz
}
