package interactors

import (
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/menu/dto"
	"github.com/gofiber/fiber/v2"
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
	CreateMenu(c *fiber.Ctx, menuReq *dto.MenuRequest) (*dto.MenuResponse, error)

	GetTreeView(elements []dto.MenuListResponse, parentId uuid.UUID) []dto.MenuListResponse

	ValidationMenuMode(c *fiber.Ctx) string

	ValidationMenuSort(c *fiber.Ctx) string

	GetMenuAllByType(c *fiber.Ctx, t stores.MenuType, mode string, sort string) ([]dto.MenuListResponse, error)
}
