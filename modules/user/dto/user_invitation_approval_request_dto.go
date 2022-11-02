package dto

type UserInvitationApprovalRequest struct {
	Code   string `json:"code" validate:"required"`
	Status string `json:"status" validate:"required,in:pending,approved,rejected"`
}
