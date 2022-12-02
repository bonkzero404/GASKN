package dto

type UserCreateRequest struct {
	FullName string `json:"full_name" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,numeric,min=10"`
}
