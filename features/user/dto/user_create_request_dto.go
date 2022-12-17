package dto

type UserCreateRequest struct {
	FullName string `json:"full_name" validate:"required,alphanum_extra,min=3,max=100"`
	Password string `json:"password" validate:"required,min=8,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,e164,min=10"`
}
