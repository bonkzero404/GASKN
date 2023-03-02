package facades

import (
	"fmt"
	"github.com/bonkzero404/gaskn/app/http"
	"github.com/gofiber/fiber/v2"
)

const (
	AppErr              string = "APP_ERR"
	AppErrNotFound      string = "APP_ERR_NOTFOUND"
	AppErrUnauthorized  string = "APP_ERR_UNAUTHORIZED"
	AppErrForbidden     string = "APP_ERR_FORBIDDEN"
	AppErrUnprocessable string = "APP_ERR_UNPROCESSABLE"
	AppErrExpire        string = "APP_ERR_EXPIRE"
)

type ResponseError struct {
	StatusCode string
	Message    string
}

func (r *ResponseError) Error() string {
	// return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Message)
	return fmt.Sprintf("%v", r.Message)
}

func ConvertToHttpError(r *ResponseError) *http.SetApiErrorResponse {
	if r.StatusCode == AppErrNotFound {
		return &http.SetApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    r.Message,
		}
	}

	if r.StatusCode == AppErrUnauthorized {
		return &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnauthorized,
			Message:    r.Message,
		}
	}

	if r.StatusCode == AppErrForbidden {
		return &http.SetApiErrorResponse{
			StatusCode: fiber.StatusForbidden,
			Message:    r.Message,
		}
	}

	if r.StatusCode == AppErrUnprocessable {
		return &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    r.Message,
		}
	}

	if r.StatusCode == AppErrExpire {
		return &http.SetApiErrorResponse{
			StatusCode: fiber.StatusGone,
			Message:    r.Message,
		}
	}

	if r.StatusCode == AppErr {
		return &http.SetApiErrorResponse{
			StatusCode: fiber.StatusInternalServerError,
			Message:    r.Message,
		}
	}

	return &http.SetApiErrorResponse{
		StatusCode: fiber.StatusInternalServerError,
		Message:    r.Message,
	}
}
