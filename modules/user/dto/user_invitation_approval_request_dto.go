package dto

import "gaskn/database/stores"

type UserInvitationApprovalRequest struct {
	Email  string                      `json:"email" validate:"required,email"`
	Code   string                      `json:"code" validate:"required"`
	Status stores.StatusInvitationType `json:"status" validate:"required,oneof=approved rejected"`
}
