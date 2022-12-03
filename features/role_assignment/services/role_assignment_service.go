package services

import (
	"errors"
	"gaskn/config"
	"gaskn/database/stores"
	"gaskn/driver"
	respModel "gaskn/dto"
	fclient "gaskn/features/role/contracts"
	"gaskn/features/role_assignment/contracts"
	"gaskn/features/role_assignment/dto"
	"gaskn/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleAssignmentService struct {
	RoleClientRepository fclient.RoleClientRepository
}

func NewRoleAssignmentService(
	RoleClientRepository fclient.RoleClientRepository,
) contracts.RoleAssignmentService {
	return &RoleAssignmentService{
		RoleClientRepository: RoleClientRepository,
	}
}

func (service RoleAssignmentService) CheckExistsRoleAssignment(c *fiber.Ctx, clientIdUuid uuid.UUID, roleIdUuid uuid.UUID) (*stores.RoleClient, error) {
	var clientRole = stores.RoleClient{}

	errRoleClient := service.RoleClientRepository.GetRoleClientId(&clientRole, roleIdUuid.String(), clientIdUuid.String()).Error

	if errors.Is(errRoleClient, gorm.ErrRecordNotFound) {
		return &stores.RoleClient{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "role:err:read-exists"),
		}
	}

	return &clientRole, nil
}

// CreateRoleAssignment /**
func (service RoleAssignmentService) CreateRoleAssignment(c *fiber.Ctx, req *dto.RoleAssignmentRequest) (*dto.RoleAssignmentResponse, error) {
	clientId := c.Params(config.Config("API_CLIENT_PARAM"))
	clientIdUuid, errClientIdUuid := uuid.Parse(clientId)
	roleIdUuid, errRoleUuid := uuid.Parse(req.RoleId)

	if errRoleUuid != nil || errClientIdUuid != nil {
		return &dto.RoleAssignmentResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:invalid-format"),
		}
	}

	existsResp, errExists := service.CheckExistsRoleAssignment(c, clientIdUuid, roleIdUuid)

	if errExists != nil {
		return nil, errExists
	}

	if save, _ := driver.AddPolicy(
		roleIdUuid.String(),
		clientIdUuid.String(),
		req.RouteFeature,
		req.MethodFeature,
		"",
		existsResp.Role.RoleName,
		existsResp.Client.ClientName,
	); !save {
		return &dto.RoleAssignmentResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "role-assign:err:failed-unknown"),
		}
	}

	saveResponse := dto.RoleAssignmentResponse{
		RoleId:   roleIdUuid.String(),
		ClientId: clientIdUuid.String(),
	}

	return &saveResponse, nil
}

func (service RoleAssignmentService) RemoveRoleAssignment(c *fiber.Ctx, req *dto.RoleAssignmentRequest) (*dto.RoleAssignmentResponse, error) {
	clientId := c.Params(config.Config("API_CLIENT_PARAM"))
	clientIdUuid, errClientIdUuid := uuid.Parse(clientId)
	roleIdUuid, errRoleUuid := uuid.Parse(req.RoleId)

	if errRoleUuid != nil || errClientIdUuid != nil {
		return &dto.RoleAssignmentResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:invalid-format"),
		}
	}

	_, errExists := service.CheckExistsRoleAssignment(c, clientIdUuid, roleIdUuid)

	if errExists != nil {
		return nil, errExists
	}

	if remove, _ := driver.RemovePolicy(
		roleIdUuid.String(),
		clientIdUuid.String(),
		req.RouteFeature,
		req.MethodFeature,
	); !remove {
		return &dto.RoleAssignmentResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "role-assign:err:failed-remove-permit"),
		}
	}

	saveResponse := dto.RoleAssignmentResponse{
		RoleId:   roleIdUuid.String(),
		ClientId: clientIdUuid.String(),
	}

	return &saveResponse, nil
}
