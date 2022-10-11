package dto

type RoleResponse struct {
	ID              string `json:"id"`
	RoleName        string `json:"role_name"`
	RoleDescription string `json:"role_description"`
	IsActive        bool   `json:"is_active"`
}
