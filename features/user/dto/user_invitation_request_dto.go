package dto

type UserInvitationRequest struct {
	Email  string `json:"email" validate:"required,email"`
	Url    string `json:"url" validate:"required,url"`
	RoleId string `json:"role_id" validate:"required,uuid4"`
}
