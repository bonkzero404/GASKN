package dto

type RoleAssignmentRequest struct {
	RoleId        string `json:"role_id" validate:"required"`
	RouteFeature  string `json:"route_feature" validate:"required"`
	MethodFeature string `json:"method_feature" validate:"required"`
}
