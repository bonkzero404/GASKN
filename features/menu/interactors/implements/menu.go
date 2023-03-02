package implements

import (
	"github.com/bonkzero404/gaskn/app/http"
	"github.com/bonkzero404/gaskn/app/translation"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/menu/dto"
	"github.com/bonkzero404/gaskn/features/menu/interactors"
	"github.com/bonkzero404/gaskn/features/menu/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/datatypes"
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

func (repository Menu) CreateMenu(req *dto.MenuRequest) (*dto.MenuResponse, error) {
	var menu = stores.Menu{
		MenuName: datatypes.JSONType[stores.LangAttribute]{Data: stores.LangAttribute{
			En: req.MenuName.En,
			Id: req.MenuName.Id,
		}},
		MenuDescription: datatypes.JSONType[stores.LangAttribute]{Data: stores.LangAttribute{
			En: req.MenuDescription.En,
			Id: req.MenuDescription.Id,
		}},
		MenuUrl:  req.MenuUrl,
		MenuIcon: req.MenuIcon,
		MenuType: req.MenuType,
		Sort:     req.Sort,
		IsActive: true,
	}

	if req.ParentId != "" {
		var getMenu = stores.Menu{}

		// Check menu if exists
		errGetMenu := repository.MenuRepository.GetMenuById(&getMenu, req.ParentId).Error

		if errGetMenu != nil {
			return nil, &http.SetApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    translation.Lang(config.MenuErrNotFound),
			}
		}

		parentId, _ := uuid.Parse(req.ParentId)
		menu.ParentID = parentId
	}

	errSaveMenu := repository.MenuRepository.CreateMenu(&menu).Error

	if errSaveMenu != nil {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    translation.Lang(config.GlobalErrUnknown),
		}
	}

	resp := dto.MenuResponse{
		ID:              menu.ID,
		MenuName:        req.MenuName,
		MenuDescription: req.MenuDescription,
		ParentId:        menu.ParentID.String(),
		MenuUrl:         req.MenuUrl,
		MenuIcon:        req.MenuIcon,
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

func (repository Menu) ValidationMenuMode(mode string) string {

	if mode == interactors.ModeList {
		return interactors.ModeList
	}

	if mode == interactors.ModeTree {
		return interactors.ModeTree
	}

	return interactors.ModeTree
}

func (repository Menu) ValidationMenuSort(sort string) string {
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

	errResult := repository.MenuRepository.GetMenuAllByType(&menuLists, translation.LangContext, t, sort).Error

	if errResult != nil {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    translation.Lang(config.MenuErrNotFound),
		}
	}

	for _, item := range menuLists {
		resp = append(resp, dto.MenuListResponse{
			ID:              item.ID,
			MenuName:        translation.LangFromJsonParse(item.MenuName),
			MenuDescription: translation.LangFromJsonParse(item.MenuDescription),
			ParentId:        item.ParentID,
			MenuUrl:         item.MenuUrl,
			MenuIcon:        item.MenuIcon,
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
