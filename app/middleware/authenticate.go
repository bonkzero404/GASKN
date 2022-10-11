package middleware

import (
	"go-starterkit-project/config"
	"go-starterkit-project/dto"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

/*
*
This function is used as middleware for authentication
*/
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
