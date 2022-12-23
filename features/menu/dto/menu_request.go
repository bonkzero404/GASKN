package dto

import "github.com/bonkzero404/gaskn/database/stores"

type MenuRequest struct {
	MenuName        string          `json:"menu_name" validate:"required,alphanum_extra"`
	MenuDescription string          `json:"menu_description,alphanum_extra"`
	ParentId        string          `json:"parent_id,omitempty" validate:"omitempty,uuid4"`
	MenuUrl         string          `json:"menu_url" validate:"omitempty,uri"`
	Sort            int             `json:"sort" validate:"number"`
	MenuType        stores.MenuType `json:"menu_type" validate:"required,oneof=BO CL"`
}
