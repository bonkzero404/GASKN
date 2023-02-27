package validations

import (
	"github.com/bonkzero404/gaskn/app/http"
	"github.com/bonkzero404/gaskn/app/translation"
	"github.com/bonkzero404/gaskn/app/validations/builder"
	"github.com/bonkzero404/gaskn/app/validations/registry"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateRequest(ctx *fiber.Ctx, request any) (bool, http.SetErrors) {
	if err := ctx.BodyParser(&request); err != nil {
		return true, http.SetErrors{
			Message: translation.Lang(ctx, "global:err:body-parser"),
			Cause:   err.Error(),
			Inputs:  nil,
		}
	}

	errors := ValidateStruct(request, ctx)
	if errors != nil {
		return true, http.SetErrors{
			Message: translation.Lang(ctx, "global:err:validate"),
			Cause:   translation.Lang(ctx, "global:err:validate-cause"),
			Inputs:  errors,
		}
	}

	return false, http.SetErrors{}
}

func ValidateStruct(s any, ctx *fiber.Ctx) []*http.SetErrorResponse {
	var errors []*http.SetErrorResponse
	var validate = validator.New()

	var trans = builder.SetLanguageValidator(ctx, validate)

	// Register custom tags
	// SetCustomTagLanguageValidator(ctx, validate, trans, "alphanum_extra", validations.ValidateAlphanumExtra)
	registry.RegisterValidator(ctx, validate, trans)

	err := validate.Struct(s)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element http.SetErrorResponse

			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Message = err.Translate(trans)
			errors = append(errors, &element)
		}
	}
	return errors
}
