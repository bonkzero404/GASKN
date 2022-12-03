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

func (service RoleAssignmentService) CheckExistsRoleAssignment(clientIdUuid uuid.UUID, roleIdUuid uuid.UUID, req *dto.RoleAssignmentRequest) (*stores.RoleClient, error) {
	var clientRole = stores.RoleClient{}

	errRoleClient := service.RoleClientRepository.GetRoleClientId(&clientRole, req.RoleId, clientIdUuid.String()).Error

	if errors.Is(errRoleClient, gorm.ErrRecordNotFound) {
		return &stores.RoleClient{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    "Role client not found",
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
			Message:    "invalid format",
		}
	}

	existsResp, errExists := service.CheckExistsRoleAssignment(clientIdUuid, roleIdUuid, req)

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
			Message:    "Telah terjadi error, kemungkinan role sudah ditetapkan",
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
			Message:    "invalid format",
		}
	}

	_, errExists := service.CheckExistsRoleAssignment(clientIdUuid, roleIdUuid, req)

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
			Message:    "Gagal menhapus permission role",
		}
	}

	saveResponse := dto.RoleAssignmentResponse{
		RoleId:   roleIdUuid.String(),
		ClientId: clientIdUuid.String(),
	}

	return &saveResponse, nil
}
