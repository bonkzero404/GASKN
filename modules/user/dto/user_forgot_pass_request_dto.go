package dto

type UserForgotPassRequest struct {
	Email string `json:"email" validate:"required,email"`
}
