package implements

import (
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/database/stores"
	globalDto "github.com/bonkzero404/gaskn/dto"
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

func (repository Menu) CreateMenu(c *fiber.Ctx, req *dto.MenuRequest) (*dto.MenuResponse, error) {
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
		errGetMenu := repository.MenuRepository.GetMenuById(&getMenu, req.ParentId).Error

		if errGetMenu != nil {
			return nil, &globalDto.ApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    utils.Lang(c, config.MenuErrNotFound),
			}
		}

		parentId, _ := uuid.Parse(req.ParentId)
		menu.ParentID = parentId
	}

	errSaveMenu := repository.MenuRepository.CreateMenu(&menu).Error

	if errSaveMenu != nil {
		return nil, &globalDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.GlobalErrUnknown),
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

func (repository Menu) GetTreeView(elements []dto.MenuListResponse, parentId uuid.UUID) []dto.MenuListResponse {
	var data []dto.MenuListResponse

	for _, element := range elements {
		if element.ParentId == parentId {
			children := repository.GetTreeView(elements, element.ID)

			if children != nil {
				element.Children = &children
			}

			data = append(data, element)
		}
	}

	return data
}

func (repository Menu) ValidationMenuMode(c *fiber.Ctx) string {
	mode := c.Query("mode")

	if mode == interactors.ModeList {
		return interactors.ModeList
	}

	if mode == interactors.ModeTree {
		return interactors.ModeTree
	}

	return interactors.ModeTree
}

func (repository Menu) ValidationMenuSort(c *fiber.Ctx) string {
	sort := c.Query("sort")

	if sort == interactors.SortAsc {
		return interactors.SortAsc
	}

	if sort == interactors.SortDesc {
		return interactors.SortDesc
	}

	return interactors.SortAsc
}

func (repository Menu) GetMenuAllByType(t stores.MenuType, mode string, sort string) ([]dto.MenuListResponse, error) {
	var menuLists []stores.Menu
	var resp []dto.MenuListResponse

	errResult := repository.MenuRepository.GetMenuAllByType(&menuLists, t, sort).Error

	if errResult != nil {
		return nil, &globalDto.ApiErrorResponse{
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
		list := repository.GetTreeView(resp, uuid.Nil)
		return list, nil
	}

	return resp, nil

}
