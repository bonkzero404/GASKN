package middleware

import (
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/dto"
	"github.com/bonkzero404/gaskn/utils"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

// Authenticate /*
func Authenticate() func(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return utils.ApiUnauthorized(ctx, dto.Errors{
				Message: utils.Lang(ctx, "middleware:err:unauthorized"),
				Cause:   err.Error(),
				Inputs:  nil,
			})
		},
		SigningKey: []byte(config.Config("JWT_SECRET")),
	})
}
