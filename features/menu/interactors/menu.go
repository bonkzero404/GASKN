package interactors

import (
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/menu/dto"
	"github.com/google/uuid"
)

//goland:noinspection GoUnusedConst,GoUnusedConst
const (
	ModeList string = "list"
	ModeTree string = "tree"
	SortAsc  string = "asc"
	SortDesc string = "desc"
)

type Menu interface {
	CreateMenu(menuReq *dto.MenuRequest) (*dto.MenuResponse, error)

	GetTreeView(elements []dto.MenuListResponse, parentId uuid.UUID) []dto.MenuListResponse

	ValidationMenuMode(mode string) string

	ValidationMenuSort(sort string) string

	GetMenuAllByType(t stores.MenuType, mode string, sort string) ([]dto.MenuListResponse, error)
}
