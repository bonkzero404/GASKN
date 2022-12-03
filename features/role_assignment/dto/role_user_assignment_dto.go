package dto

type RoleUserAssignment struct {
	UserId string `json:"user_id" validate:"required"`
	RoleId string `json:"role_id" validate:"required"`
}
