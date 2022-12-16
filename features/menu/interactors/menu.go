package interactors

import (
	"github.com/bonkzero404/gaskn/features/menu/dto"
	"github.com/gofiber/fiber/v2"
)

type Menu interface {
	CreateMenu(c *fiber.Ctx, menuReq *dto.MenuRequest) (*dto.MenuResponse, error)
}
