package dto

type UserInvitationResponse struct {
	InvitedTo     string `json:"invited_to"`
	InvitedBy     string `json:"invited_by"`
	InvitedToRole string `json:"invited_to_role"`
	ClientId      string `json:"client_id"`
}
