package implements

import (
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/database/stores"
	responseDto "github.com/bonkzero404/gaskn/dto"
	"github.com/bonkzero404/gaskn/features/role/dto"
	"github.com/bonkzero404/gaskn/features/role/interactors"
	"github.com/bonkzero404/gaskn/features/role/repositories"
	"github.com/bonkzero404/gaskn/utils"

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

func (interact Role) CreateRole(c *fiber.Ctx, role *dto.RoleRequest) (*dto.RoleResponse, error) {

	roleData := stores.Role{
		RoleName:        role.RoleName,
		RoleDescription: role.RoleDescription,
		IsActive:        true,
	}

	err := interact.RoleRepository.CreateRole(&roleData).Error

	if err != nil {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.GlobalErrUnknown),
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

func (interact Role) GetRoleList(c *fiber.Ctx) (*utils.Pagination, error) {
	var roles []stores.Role
	var resp []dto.RoleResponse

	res, err := interact.RoleRepository.GetRoleList(&roles, c)

	if err != nil {
		return nil, &responseDto.ApiErrorResponse{
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

func (interact Role) UpdateRole(c *fiber.Ctx, id string, role *dto.RoleRequest) (*dto.RoleResponse, error) {
	// Check role if exists
	var roleStore stores.Role

	errCheckRole := interact.RoleRepository.GetRoleById(&roleStore, id).Error

	if errCheckRole != nil {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.RoleErrNotExists),
		}
	}

	roleStore.RoleName = role.RoleName
	roleStore.RoleDescription = role.RoleDescription
	roleStore.IsActive = true

	err := interact.RoleRepository.UpdateRoleById(&roleStore).Error

	if err != nil {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.GlobalErrUnknown),
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

func (interact Role) DeleteRoleById(c *fiber.Ctx, id string) (*dto.RoleResponse, error) {
	// Check role if exists
	var roleStore stores.Role

	errCheckRole := interact.RoleRepository.GetRoleById(&roleStore, id).Error

	if errCheckRole != nil {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.RoleErrNotExists),
		}
	}

	err := interact.RoleRepository.DeleteRoleById(&roleStore).Error

	if err != nil {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.GlobalErrUnknown),
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
