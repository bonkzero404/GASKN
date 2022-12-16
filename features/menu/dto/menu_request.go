package dto

import "github.com/bonkzero404/gaskn/database/stores"

type MenuRequest struct {
	MenuName        string          `json:"menu_name" validate:"required,alpha"`
	MenuDescription string          `json:"menu_description,alphanum"`
	ParentId        string          `json:"parent_id,omitempty" validate:"omitempty,uuid4"`
	MenuUrl         string          `json:"menu_url" validate:"required,uri"`
	MenuType        stores.MenuType `json:"menu_type" validate:"required,oneof=BO CL"`
}
