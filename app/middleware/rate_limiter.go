package middleware

import (
	"gaskn/utils"
	"time"

	respModel "gaskn/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

/*
*
This function is used as middleware for rate limit request
*/
func RateLimiter(max int, duration time.Duration) func(ctx *fiber.Ctx) error {
	return limiter.New(limiter.Config{
		LimitReached: func(ctx *fiber.Ctx) error {
			return utils.ApiResponseError(ctx, fiber.StatusRequestEntityTooLarge, respModel.Errors{
				Message: utils.Lang(ctx, "middleware:err:failed"),
				Cause:   utils.Lang(ctx, "middleware:err:rate-limiter"),
				Inputs:  nil,
			})
		},
		Max:        max,
		Expiration: duration * time.Second,
	})
}
