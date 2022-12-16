package dto

import "github.com/bonkzero404/gaskn/database/stores"

type UserInvitationApprovalRequest struct {
	Code   string                      `json:"code" validate:"required,alphanum,max=100"`
	Status stores.StatusInvitationType `json:"status" validate:"required,oneof=approved rejected"`
}
