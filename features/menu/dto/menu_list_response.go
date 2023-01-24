package dto

import (
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/google/uuid"
)

type MenuListResponse struct {
	ID              uuid.UUID           `json:"id"`
	MenuName        string              `json:"menu_name"`
	MenuDescription string              `json:"menu_description"`
	ParentId        uuid.UUID           `json:"parent_id"`
	MenuUrl         string              `json:"menu_url"`
	MenuIcon        string              `json:"menu_icon"`
	MenuType        stores.MenuType     `json:"menu_type"`
	Sort            int                 `json:"sort"`
	Children        *[]MenuListResponse `json:"children"`
}
