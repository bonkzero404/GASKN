package response

import (
	"github.com/bonkzero404/gaskn/app/http"
	"github.com/bonkzero404/gaskn/app/logger"
	"github.com/gofiber/fiber/v2"
)

// ApiResponse /*
func ApiResponse(ctx *fiber.Ctx, code int, status string, data any, errors any) error {
	meta := http.SetMeta{
		Route:  ctx.Route().Path,
		Method: ctx.Method(),
		Query:  string(ctx.Request().URI().QueryString()),
		Code:   code,
		Status: status,
	}

	if code >= 400 {
		responseJson := http.SetResponse{
			Valid: false,
			Meta:  meta,
			Error: errors,
			Data:  nil,
		}

		logger.CreateAccessLog(ctx, "[ACCESS][ERROR]", code, errors)
		return ctx.Status(code).JSON(responseJson)
	}

	responseJson := http.SetResponse{
		Valid: true,
		Meta:  meta,
		Error: nil,
		Data:  data,
	}

	logger.CreateAccessLog(ctx, "[ACCESS][SUCCESS]", code, data)
	return ctx.Status(code).JSON(responseJson)
}

func ApiErrorValidation(ctx *fiber.Ctx, errors http.SetErrors) error {
	return ApiResponse(ctx, fiber.StatusNotAcceptable, "error_validation", nil, errors)
}

func ApiUnprocessableEntity(ctx *fiber.Ctx, errors http.SetErrors) error {
	return ApiResponse(ctx, fiber.StatusUnprocessableEntity, "error_unprocessable_entity", nil, errors)
}

func ApiUnauthorized(ctx *fiber.Ctx, errors http.SetErrors) error {
	return ApiResponse(ctx, fiber.StatusUnauthorized, "error_unauthorized", nil, errors)
}

func ApiCreated(ctx *fiber.Ctx, data any) error {
	return ApiResponse(ctx, fiber.StatusCreated, "success_created", data, nil)
}

func ApiOk(ctx *fiber.Ctx, data any) error {
	return ApiResponse(ctx, fiber.StatusOK, "success_ok", data, nil)
}

func ApiResponseError(ctx *fiber.Ctx, code int, errors http.SetErrors) error {
	return ApiResponse(ctx, code, "error_api", nil, errors)
}

func ApiForbidden(ctx *fiber.Ctx, errors http.SetErrors) error {
	return ApiResponse(ctx, fiber.StatusForbidden, "error_forbidden", nil, errors)
}
