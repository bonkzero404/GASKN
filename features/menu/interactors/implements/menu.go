package implements

import (
	"github.com/bonkzero404/gaskn/database/stores"
	responseDto "github.com/bonkzero404/gaskn/dto"
	"github.com/bonkzero404/gaskn/features/menu/dto"
	"github.com/bonkzero404/gaskn/features/menu/interactors"
	"github.com/bonkzero404/gaskn/features/menu/repositories"
	"github.com/bonkzero404/gaskn/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Menu struct {
	MenuRepository repositories.MenuRepository
}

func NewMenu(
	MenuRepository repositories.MenuRepository,
) interactors.Menu {
	return &Menu{
		MenuRepository: MenuRepository,
	}
}

func (interact Menu) CreateMenu(c *fiber.Ctx, req *dto.MenuRequest) (*dto.MenuResponse, error) {
	var menu = stores.Menu{
		MenuName:        req.MenuName,
		MenuDescription: req.MenuDescription,
		MenuUrl:         req.MenuUrl,
		MenuType:        req.MenuType,
		IsActive:        true,
	}

	if req.ParentId != "" {
		var getMenu = stores.Menu{}

		// Check menu if exists
		errGetMenu := interact.MenuRepository.GetMenuById(&getMenu, req.ParentId)

		if errGetMenu != nil {
			return nil, &responseDto.ApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    utils.Lang(c, "blabla"),
			}
		}

		parentId, _ := uuid.Parse(req.ParentId)
		menu.ParentID = parentId
	}

	errSaveMenu := interact.MenuRepository.CreateMenu(&menu).Error

	if errSaveMenu != nil {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:failed-unknown"),
		}
	}

	resp := dto.MenuResponse{
		MenuName:        req.MenuName,
		MenuDescription: req.MenuDescription,
		ParentId:        menu.ParentID.String(),
		MenuUrl:         req.MenuUrl,
		MenuType:        req.MenuType,
	}

	return &resp, nil

}
