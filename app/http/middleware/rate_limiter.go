package middleware

import (
	"github.com/bonkzero404/gaskn/app/http"
	"github.com/bonkzero404/gaskn/app/http/response"
	"github.com/bonkzero404/gaskn/app/translation"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// RateLimiter /*
func RateLimiter(max int, duration time.Duration) func(ctx *fiber.Ctx) error {
	return limiter.New(limiter.Config{
		LimitReached: func(ctx *fiber.Ctx) error {
			return response.ApiResponseError(ctx, fiber.StatusRequestEntityTooLarge, http.SetErrors{
				Message: translation.Lang("middleware:err:failed"),
				Cause:   translation.Lang("middleware:err:rate-limiter"),
				Inputs:  nil,
			})
		},
		Max:        max,
		Expiration: duration * time.Second,
	})
}
