package implements

import (
	"errors"
	"github.com/bonkzero404/gaskn/app/http"
	"github.com/bonkzero404/gaskn/app/translation"
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/role/dto"
	"github.com/bonkzero404/gaskn/features/role/interactors"
	"github.com/bonkzero404/gaskn/features/role/repositories"
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

func (repository RoleClient) CreateRoleClient(c *fiber.Ctx, roleDto *dto.RoleRequest) (*dto.RoleResponse, error) {
	var roleClient stores.RoleClient

	// Get client id from param url
	clientId := c.Params(config.Config("API_CLIENT_PARAM"))

	// Check role if inserted
	errRoleClient := repository.RoleClientRepository.GetRoleClientByName(&roleClient, roleDto.RoleName, clientId).Error

	if errors.Is(errRoleClient, gorm.ErrRecordNotFound) {
		roleData := stores.Role{
			RoleName:        roleDto.RoleName,
			RoleDescription: roleDto.RoleDescription,
			IsActive:        true,
			RoleType:        stores.CL,
		}

		r, err := repository.RoleClientRepository.CreateRoleClient(&roleData, clientId)

		if err != nil {
			return nil, &http.SetApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    translation.Lang(c, config.GlobalErrUnknown),
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

	return nil, &http.SetApiErrorResponse{
		StatusCode: fiber.StatusUnprocessableEntity,
		Message:    translation.Lang(c, config.RoleErrAlreadyExists),
	}
}

func (repository RoleClient) GetRoleClientList(c *fiber.Ctx) (*utils.Pagination, error) {
	var roles []stores.Role
	var resp []dto.RoleResponse

	// Get client id from param url
	clientId := c.Params(config.Config("API_CLIENT_PARAM"))

	res, err := repository.RoleClientRepository.GetRoleClientList(&roles, c, clientId)

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

func (repository RoleClient) UpdateRoleClient(c *fiber.Ctx, id string) (*dto.RoleResponse, error) {
	var roleStore stores.Role

	var roleClient stores.RoleClient

	// Get client id from param url
	clientId := c.Params(config.Config("API_CLIENT_PARAM"))

	// Check role if inserted
	errRoleClient := repository.RoleClientRepository.GetRoleClientById(&roleClient, id, clientId).Error

	if errRoleClient != nil {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    translation.Lang(c, config.RoleErrNotExists),
		}
	}

	roleStore.ID = roleClient.Role.ID
	roleStore.RoleName = roleClient.Role.RoleName
	roleStore.RoleDescription = roleClient.Role.RoleDescription
	roleStore.IsActive = true

	err := repository.RoleRepository.UpdateRoleById(&roleStore).Error

	if err != nil {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    translation.Lang(c, config.GlobalErrUnknown),
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

func (repository RoleClient) DeleteRoleClientById(c *fiber.Ctx, id string) (*dto.RoleResponse, error) {
	var roleStore stores.Role

	var roleClient stores.RoleClient

	// Get client id from param url
	clientId := c.Params(config.Config("API_CLIENT_PARAM"))

	// Check role if inserted
	errRoleClient := repository.RoleClientRepository.GetRoleClientId(&roleClient, id, clientId).Error

	if errRoleClient != nil {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    translation.Lang(c, config.RoleErrNotExists),
		}
	}

	roleStore.ID = roleClient.Role.ID
	roleStore.RoleName = roleClient.Role.RoleName
	roleStore.RoleDescription = roleClient.Role.RoleDescription
	roleStore.IsActive = true

	err := repository.RoleRepository.DeleteRoleById(&roleStore).Error

	if err != nil {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    translation.Lang(c, config.GlobalErrUnknown),
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
