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

func (repository RoleClient) CreateRoleClient(clientId string, roleDto *dto.RoleRequest) (*dto.RoleResponse, error) {
	var roleClient stores.RoleClient

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
				Message:    translation.Lang(config.GlobalErrUnknown),
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
		Message:    translation.Lang(config.RoleErrAlreadyExists),
	}
}

func (repository RoleClient) GetRoleClientList(clientId string, page string, limit string, sort string) (*utils.Pagination, error) {
	var roles []stores.Role
	var resp []dto.RoleResponse

	paginateRequest := utils.PaginationRequest{
		Page:  page,
		Limit: limit,
		Sort:  sort,
	}

	pPage, pLimit, pSort := paginateRequest.SetPagination()

	res, err := repository.RoleClientRepository.GetRoleClientList(&roles, clientId, pPage, pLimit, pSort)

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

func (repository RoleClient) UpdateRoleClient(clientId string, id string) (*dto.RoleResponse, error) {
	var roleStore stores.Role

	var roleClient stores.RoleClient

	// Check role if inserted
	errRoleClient := repository.RoleClientRepository.GetRoleClientById(&roleClient, id, clientId).Error

	if errRoleClient != nil {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    translation.Lang(config.RoleErrNotExists),
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

func (repository RoleClient) DeleteRoleClientById(clientId string, id string) (*dto.RoleResponse, error) {
	var roleStore stores.Role

	var roleClient stores.RoleClient

	// Check role if inserted
	errRoleClient := repository.RoleClientRepository.GetRoleClientId(&roleClient, id, clientId).Error

	if errRoleClient != nil {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    translation.Lang(config.RoleErrNotExists),
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
