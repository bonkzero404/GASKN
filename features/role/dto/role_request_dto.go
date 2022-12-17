package dto

type RoleRequest struct {
	RoleName        string `json:"role_name" validate:"required,alphanum_extra"`
	RoleDescription string `json:"role_description" validate:"alphanum_extra"`
}
