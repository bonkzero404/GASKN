package dto

type UserInvitationRequest struct {
	Email string `json:"email" validate:"required,email"`
}
