package dto

type RoleAssignmentRequest struct {
	RoleId             string `json:"role_id" validate:"required,uuid4"`
	RouteFeature       string `json:"route_feature" validate:"required,uri"`
	MethodFeature      string `json:"method_feature" validate:"required,oneof=POST GET PUT DELETE"`
	RouteGroup         string `json:"route_group" validate:"required"`
	RouteName          string `json:"route_name" validate:"required"`
	DescriptionKeyLang string `json:"description_key_lang" validate:"required"`
}
