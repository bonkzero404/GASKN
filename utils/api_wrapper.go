package utils

import (
	"gaskn/dto"

	"github.com/gofiber/fiber/v2"
)

// ApiWrapper /*
func ApiWrapper(ctx *fiber.Ctx, code int, status string, data interface{}, errors interface{}) error {
	meta := dto.Meta{
		Route:  ctx.Route().Path,
		Method: ctx.Method(),
		Query:  string(ctx.Request().URI().QueryString()),
		Code:   code,
		Status: status,
	}

	if code >= 400 {
		responseJson := dto.Response{
			Valid: false,
			Meta:  meta,
			Error: errors,
			Data:  nil,
		}

		WriteRequestToLog(ctx, "[ACCESS][ERROR]", code, errors)
		return ctx.Status(code).JSON(responseJson)
	}

	responseJson := dto.Response{
		Valid: true,
		Meta:  meta,
		Error: nil,
		Data:  data,
	}

	WriteRequestToLog(ctx, "[ACCESS][SUCCESS]", code, data)
	return ctx.Status(code).JSON(responseJson)
}

func ApiErrorValidation(ctx *fiber.Ctx, errors dto.Errors) error {
	return ApiWrapper(ctx, fiber.StatusNotAcceptable, "error_validation", nil, errors)
}

func ApiUnprocessableEntity(ctx *fiber.Ctx, errors dto.Errors) error {
	return ApiWrapper(ctx, fiber.StatusUnprocessableEntity, "error_unprocessable_entity", nil, errors)
}

func ApiUnauthorized(ctx *fiber.Ctx, errors dto.Errors) error {
	return ApiWrapper(ctx, fiber.StatusUnauthorized, "error_unauthorized", nil, errors)
}

func ApiCreated(ctx *fiber.Ctx, data interface{}) error {
	return ApiWrapper(ctx, fiber.StatusCreated, "success_created", data, nil)
}

func ApiOk(ctx *fiber.Ctx, data interface{}) error {
	return ApiWrapper(ctx, fiber.StatusOK, "success_ok", data, nil)
}

func ApiResponseError(ctx *fiber.Ctx, code int, errors dto.Errors) error {
	return ApiWrapper(ctx, code, "error_api", nil, errors)
}

func ApiForbidden(ctx *fiber.Ctx, errors dto.Errors) error {
	return ApiWrapper(ctx, fiber.StatusForbidden, "error_forbidden", nil, errors)
}
