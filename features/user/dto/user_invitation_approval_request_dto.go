package dto

import "gaskn/database/stores"

type UserInvitationApprovalRequest struct {
	Code   string                      `json:"code" validate:"required"`
	Status stores.StatusInvitationType `json:"status" validate:"required,oneof=approved rejected"`
}
