package dto

type RoleAssignmentResponse struct {
	RoleId     string `json:"role_id"`
	ClientName string `json:"client_name"`
	UserName   string `json:"user"`
}
