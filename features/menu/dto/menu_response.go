package dto

import (
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/google/uuid"
)

type MenuResponse struct {
	ID              uuid.UUID            `json:"id"`
	MenuName        stores.LangAttribute `json:"menu_name"`
	MenuDescription stores.LangAttribute `json:"menu_description"`
	ParentId        string               `json:"parent_id"`
	MenuUrl         string               `json:"menu_url"`
	MenuIcon        string               `json:"menu_icon"`
	MenuType        stores.MenuType      `json:"menu_type"`
}
