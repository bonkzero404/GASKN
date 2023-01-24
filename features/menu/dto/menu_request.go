package dto

import "github.com/bonkzero404/gaskn/database/stores"

type MenuRequest struct {
	MenuName        stores.LangAttribute `json:"menu_name" validate:"required"`
	MenuDescription stores.LangAttribute `json:"menu_description"`
	ParentId        string               `json:"parent_id,omitempty" validate:"omitempty,uuid4"`
	MenuUrl         string               `json:"menu_url" validate:"omitempty,uri"`
	MenuIcon        string               `json:"menu_icon" validate:"omitempty"`
	Sort            int                  `json:"sort" validate:"number"`
	MenuType        stores.MenuType      `json:"menu_type" validate:"required,oneof=BO CL"`
}
