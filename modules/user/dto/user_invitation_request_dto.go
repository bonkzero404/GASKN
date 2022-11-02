package dto

type UserInvitationRequest struct {
	Email string `json:"email" validate:"required,email"`
	Url   string `json:"url" validate:"required"`
}
