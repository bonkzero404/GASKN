package dto

type RoleAssignmentListResponse struct {
	ID           string `json:"id"`
	PermissionId uint   `json:"permission_id"`
	RoleId       string `json:"role_id"`
	ClientName   string `json:"client_name"`
	RoleName     string `json:"role_name"`
	GroupName    string `json:"group_name"`
	RouteName    string `json:"route_name"`
	Description  string `json:"description"`
	Route        string `json:"route"`
	Method       string `json:"method"`
}
