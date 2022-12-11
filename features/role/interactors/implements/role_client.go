package implements

import (
	"errors"
	"gaskn/config"
	"gaskn/database/stores"
	respModel "gaskn/dto"
	"gaskn/features/role/dto"
	"gaskn/features/role/interactors"
	"gaskn/features/role/repositories"
	"gaskn/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RoleClient struct {
	RoleClientRepository repositories.RoleClientRepository
	RoleRepository       repositories.RoleRepository
}

func NewRoleClient(
	roleClientRepository repositories.RoleClientRepository,
	roleRepository repositories.RoleRepository,
) interactors.RoleClient {
	return &RoleClient{
		RoleClientRepository: roleClientRepository,
		RoleRepository:       roleRepository,
	}
}

func (interact RoleClient) CreateRoleClient(c *fiber.Ctx, roleDto *dto.RoleRequest) (*dto.RoleResponse, error) {
	var roleClient stores.RoleClient

	// Get client id from param url
	clientId := c.Params(config.Config("API_CLIENT_PARAM"))

	// Check role if inserted
	errRoleClient := interact.RoleClientRepository.GetRoleClientByName(&roleClient, roleDto.RoleName, clientId).Error

	if errors.Is(errRoleClient, gorm.ErrRecordNotFound) {
		roleData := stores.Role{
			RoleName:        roleDto.RoleName,
			RoleDescription: roleDto.RoleDescription,
			IsActive:        true,
			RoleType:        stores.CL,
		}

		r, err := interact.RoleClientRepository.CreateRoleClient(&roleData, clientId)

		if err != nil {
			return &dto.RoleResponse{}, &respModel.ApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    utils.Lang(c, "global:err:failed-unknown"),
			}
		}

		roleResponse := dto.RoleResponse{
			ID:              r.ID.String(),
			RoleName:        r.RoleName,
			RoleDescription: r.RoleDescription,
			IsActive:        r.IsActive,
		}

		return &roleResponse, nil
	}

	return &dto.RoleResponse{}, &respModel.ApiErrorResponse{
		StatusCode: fiber.StatusUnprocessableEntity,
		Message:    utils.Lang(c, "role:err:read-available"),
	}
}

func (interact RoleClient) GetRoleClientList(c *fiber.Ctx) (*utils.Pagination, error) {
	var roles []stores.Role
	var resp []dto.RoleResponse

	// Get client id from param url
	clientId := c.Params(config.Config("API_CLIENT_PARAM"))

	res, err := interact.RoleClientRepository.GetRoleClientList(&roles, c, clientId)

	if err != nil {
		return nil, &respModel.ApiErrorResponse{
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

func (interact RoleClient) UpdateRoleClient(c *fiber.Ctx, id string) (*dto.RoleResponse, error) {
	var roleStore stores.Role

	var roleClient stores.RoleClient

	// Get client id from param url
	clientId := c.Params(config.Config("API_CLIENT_PARAM"))

	// Check role if inserted
	errRoleClient := interact.RoleClientRepository.GetRoleClientById(&roleClient, id, clientId).Error

	if errRoleClient != nil {
		return &dto.RoleResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "role:err:read-exists"),
		}
	}

	roleStore.ID = roleClient.Role.ID
	roleStore.RoleName = roleClient.Role.RoleName
	roleStore.RoleDescription = roleClient.Role.RoleDescription
	roleStore.IsActive = true

	err := interact.RoleRepository.UpdateRoleById(&roleStore).Error

	if err != nil {
		return &dto.RoleResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:failed-unknown"),
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

func (interact RoleClient) DeleteRoleClientById(c *fiber.Ctx, id string) (*dto.RoleResponse, error) {
	var roleStore stores.Role

	var roleClient stores.RoleClient

	// Get client id from param url
	clientId := c.Params(config.Config("API_CLIENT_PARAM"))

	// Check role if inserted
	errRoleClient := interact.RoleClientRepository.GetRoleClientById(&roleClient, id, clientId).Error

	if errRoleClient != nil {
		return &dto.RoleResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "role:err:read-exists"),
		}
	}

	roleStore.ID = roleClient.Role.ID
	roleStore.RoleName = roleClient.Role.RoleName
	roleStore.RoleDescription = roleClient.Role.RoleDescription
	roleStore.IsActive = true

	err := interact.RoleRepository.DeleteRoleById(&roleStore).Error

	if err != nil {
		return &dto.RoleResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:failed-unknown"),
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
