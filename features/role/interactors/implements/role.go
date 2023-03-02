package implements

import (
	"github.com/bonkzero404/gaskn/app/http"
	"github.com/bonkzero404/gaskn/app/translation"
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/role/dto"
	"github.com/bonkzero404/gaskn/features/role/interactors"
	"github.com/bonkzero404/gaskn/features/role/repositories"
	"github.com/gofiber/fiber/v2"
)

type Role struct {
	RoleRepository repositories.RoleRepository
}

func NewRole(
	roleRepository repositories.RoleRepository,
) interactors.Role {
	return &Role{
		RoleRepository: roleRepository,
	}
}

func (repository Role) CreateRole(role *dto.RoleRequest) (*dto.RoleResponse, error) {

	roleData := stores.Role{
		RoleName:        role.RoleName,
		RoleDescription: role.RoleDescription,
		IsActive:        true,
	}

	err := repository.RoleRepository.CreateRole(&roleData).Error

	if err != nil {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    translation.Lang(config.GlobalErrUnknown),
		}
	}

	roleResponse := dto.RoleResponse{
		ID:              roleData.ID.String(),
		RoleName:        roleData.RoleName,
		RoleDescription: roleData.RoleDescription,
		IsActive:        roleData.IsActive,
	}

	return &roleResponse, nil
}

func (repository Role) GetRoleList(page string, limit string, sort string) (*utils.Pagination, error) {
	var roles []stores.Role
	var resp []dto.RoleResponse

	paginateRequest := utils.PaginationRequest{
		Page:  page,
		Limit: limit,
		Sort:  sort,
	}

	pPage, pLimit, pSort := paginateRequest.SetPagination()

	res, err := repository.RoleRepository.GetRoleList(&roles, pPage, pLimit, pSort)

	if err != nil {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    err.Error(),
		}
	}

	for _, item := range roles {
		resp = append(resp, dto.RoleResponse{
			ID:              item.ID.String(),
			RoleName:        item.RoleName,
			RoleDescription: item.RoleDescription,
			IsActive:        item.IsActive,
		})
	}

	res.Rows = resp

	return res, nil
}

func (repository Role) UpdateRole(id string, role *dto.RoleRequest) (*dto.RoleResponse, error) {
	// Check role if exists
	var roleStore stores.Role

	errCheckRole := repository.RoleRepository.GetRoleById(&roleStore, id).Error

	if errCheckRole != nil {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    translation.Lang(config.RoleErrNotExists),
		}
	}

	roleStore.RoleName = role.RoleName
	roleStore.RoleDescription = role.RoleDescription
	roleStore.IsActive = true

	err := repository.RoleRepository.UpdateRoleById(&roleStore).Error

	if err != nil {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    translation.Lang(config.GlobalErrUnknown),
		}
	}

	roleResponse := dto.RoleResponse{
		ID:              roleStore.ID.String(),
		RoleName:        roleStore.RoleName,
		RoleDescription: roleStore.RoleDescription,
		IsActive:        roleStore.IsActive,
	}

	return &roleResponse, nil
}

func (repository Role) DeleteRoleById(id string) (*dto.RoleResponse, error) {
	// Check role if exists
	var roleStore stores.Role

	errCheckRole := repository.RoleRepository.GetRoleById(&roleStore, id).Error

	if errCheckRole != nil {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    translation.Lang(config.RoleErrNotExists),
		}
	}

	err := repository.RoleRepository.DeleteRoleById(&roleStore).Error

	if err != nil {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    translation.Lang(config.GlobalErrUnknown),
		}
	}

	roleResponse := dto.RoleResponse{
		ID:              roleStore.ID.String(),
		RoleName:        roleStore.RoleName,
		RoleDescription: roleStore.RoleDescription,
		IsActive:        roleStore.IsActive,
	}

	return &roleResponse, nil
}
