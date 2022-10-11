package dto

type UserReActivationRequest struct {
	Email string `json:"email" validate:"required,email"`
}
