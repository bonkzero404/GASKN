package dto

type UserActivationRequest struct {
	Code string `json:"code" validate:"required"`
}
