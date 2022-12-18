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
		Sort:            req.Sort,
		IsActive:        true,
	}

	if req.ParentId != "" {
		var getMenu = stores.Menu{}

		// Check menu if exists
		errGetMenu := interact.MenuRepository.GetMenuById(&getMenu, req.ParentId).Error

		if errGetMenu != nil {
			return nil, &responseDto.ApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    utils.Lang(c, "menu:err:menu-not-found"),
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
		ID:              menu.ID,
		MenuName:        req.MenuName,
		MenuDescription: req.MenuDescription,
		ParentId:        menu.ParentID.String(),
		MenuUrl:         req.MenuUrl,
		MenuType:        req.MenuType,
	}

	return &resp, nil
}

func (interact Menu) GetTreeView(elements []dto.MenuListResponse, parentId uuid.UUID) []dto.MenuListResponse {
	var data []dto.MenuListResponse

	for _, element := range elements {
		if element.ParentId == parentId {
			children := interact.GetTreeView(elements, element.ID)

			if children != nil {
				element.Children = &children
			}

			data = append(data, element)
		}
	}

	return data
}

func (interact Menu) ValidationMenuMode(c *fiber.Ctx) string {
	mode := c.Query("mode")

	if mode == interactors.ModeList {
		return interactors.ModeList
	}

	if mode == interactors.ModeTree {
		return interactors.ModeTree
	}

	return interactors.ModeTree
}

func (interact Menu) GetMenuAllByType(t stores.MenuType, mode string) ([]dto.MenuListResponse, error) {
	var menuLists []stores.Menu
	var resp []dto.MenuListResponse

	errResult := interact.MenuRepository.GetMenuAllByType(&menuLists, t).Error

	if errResult != nil {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "blabla",
		}
	}

	for _, item := range menuLists {
		resp = append(resp, dto.MenuListResponse{
			ID:              item.ID,
			MenuName:        item.MenuName,
			MenuDescription: item.MenuDescription,
			ParentId:        item.ParentID,
			MenuUrl:         item.MenuUrl,
			Sort:            item.Sort,
			MenuType:        item.MenuType,
		})
	}

	if mode == interactors.ModeTree {
		list := interact.GetTreeView(resp, uuid.Nil)
		return list, nil
	}

	return resp, nil

}
