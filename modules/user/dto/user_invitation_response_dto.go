package dto

type UserInvitationResponse struct {
	Email string `json:"email" validate:"required,email"`
	Url   string `json:"url" validate:"required"`
}
