package dto

type UserForgotPassActRequest struct {
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,min=8"`
	RepeatPassword string `json:"repeat_password" validate:"required,min=8"`
	Code           string `json:"code" validate:"required"`
}
