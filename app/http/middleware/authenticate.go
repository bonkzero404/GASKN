package middleware

import (
	"github.com/bonkzero404/gaskn/app/http"
	"github.com/bonkzero404/gaskn/app/http/response"
	"github.com/bonkzero404/gaskn/app/translation"
	"github.com/bonkzero404/gaskn/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v3"
)

// Authenticate /*
func Authenticate() func(ctx *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return response.ApiUnauthorized(ctx, http.SetErrors{
				Message: translation.Lang("middleware:err:unauthorized"),
				Cause:   err.Error(),
				Inputs:  nil,
			})
		},
		SigningKey: []byte(config.Config("JWT_SECRET")),
	})
}
