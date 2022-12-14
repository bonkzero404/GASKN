package implements

import (
	"errors"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/driver"
	responseDto "github.com/bonkzero404/gaskn/dto"
	"github.com/bonkzero404/gaskn/features/role/repositories"
	"github.com/bonkzero404/gaskn/features/role_assignment/dto"
	"github.com/bonkzero404/gaskn/features/role_assignment/interactors"
	"github.com/bonkzero404/gaskn/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleAssignment struct {
	RoleClientRepository repositories.RoleClientRepository
	RoleRepository       repositories.RoleRepository
}

func NewRoleAssignment(
	RoleClientRepository repositories.RoleClientRepository,
	RoleRepository repositories.RoleRepository,
) interactors.RoleAssignment {
	return &RoleAssignment{
		RoleClientRepository: RoleClientRepository,
		RoleRepository:       RoleRepository,
	}
}

func (interact RoleAssignment) CheckExistsRoleAssignment(c *fiber.Ctx, clientIdUuid uuid.UUID, roleIdUuid uuid.UUID) (*stores.RoleClient, error) {
	var clientRole = stores.RoleClient{}

	errRoleClient := interact.RoleClientRepository.GetRoleClientId(&clientRole, roleIdUuid.String(), clientIdUuid.String()).Error

	if errors.Is(errRoleClient, gorm.ErrRecordNotFound) {
		return &stores.RoleClient{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "role:err:read-exists"),
		}
	}

	return &clientRole, nil
}

func (interact RoleAssignment) CheckExistsRoleUserAssignment(c *fiber.Ctx, userId uuid.UUID, clientIdUuid uuid.UUID) (*stores.ClientAssignment, error) {
	var clientAssign = stores.ClientAssignment{}

	errRoleUserClient := interact.RoleClientRepository.GetUserHasClient(
		&clientAssign,
		userId.String(),
		clientIdUuid.String(),
	).Error

	if errors.Is(errRoleUserClient, gorm.ErrRecordNotFound) {
		return &stores.ClientAssignment{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "role:err:read-exists"),
		}
	}

	return &clientAssign, nil
}

// CreateRoleAssignment /**
func (interact RoleAssignment) CreateRoleAssignment(c *fiber.Ctx, req *dto.RoleAssignmentRequest) (*dto.RoleAssignmentResponse, error) {
	clientId := c.Params(config.Config("API_CLIENT_PARAM"))
	clientIdUuid, errClientIdUuid := uuid.Parse(clientId)
	roleIdUuid, errRoleUuid := uuid.Parse(req.RoleId)

	if errRoleUuid != nil || errClientIdUuid != nil {
		return &dto.RoleAssignmentResponse{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:invalid-format"),
		}
	}

	existsResp, errExists := interact.CheckExistsRoleAssignment(c, clientIdUuid, roleIdUuid)

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
		return &dto.RoleAssignmentResponse{}, &responseDto.ApiErrorResponse{
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

func (interact RoleAssignment) RemoveRoleAssignment(c *fiber.Ctx, req *dto.RoleAssignmentRequest) (*dto.RoleAssignmentResponse, error) {
	clientId := c.Params(config.Config("API_CLIENT_PARAM"))
	clientIdUuid, errClientIdUuid := uuid.Parse(clientId)
	roleIdUuid, errRoleUuid := uuid.Parse(req.RoleId)

	if errRoleUuid != nil || errClientIdUuid != nil {
		return &dto.RoleAssignmentResponse{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:invalid-format"),
		}
	}

	_, errExists := interact.CheckExistsRoleAssignment(c, clientIdUuid, roleIdUuid)

	if errExists != nil {
		return nil, errExists
	}

	if remove, _ := driver.RemovePolicy(
		roleIdUuid.String(),
		clientIdUuid.String(),
		req.RouteFeature,
		req.MethodFeature,
	); !remove {
		return &dto.RoleAssignmentResponse{}, &responseDto.ApiErrorResponse{
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

func (interact RoleAssignment) AssignUserPermitToRole(c *fiber.Ctx, req *dto.RoleUserAssignment) (*dto.RoleAssignmentResponse, error) {
	clientId := c.Params(config.Config("API_CLIENT_PARAM"))
	clientIdUuid, errClientIdUuid := uuid.Parse(clientId)
	userIdUuid, errUserUuid := uuid.Parse(req.UserId)
	roleIdUuid, errRoleUuid := uuid.Parse(req.RoleId)

	if errRoleUuid != nil || errClientIdUuid != nil || errUserUuid != nil {
		return &dto.RoleAssignmentResponse{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:invalid-format"),
		}
	}

	// Check if user has client
	existsResp, errExists := interact.CheckExistsRoleUserAssignment(c, userIdUuid, clientIdUuid)

	if errExists != nil {
		return nil, errExists
	}

	// Check if role has available
	var role = stores.Role{}
	errRole := interact.RoleRepository.GetRoleById(&role, roleIdUuid.String()).Error

	if errRole != nil {
		return &dto.RoleAssignmentResponse{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "role:err:read-exists"),
		}
	}

	saveUserRoleClient := interact.RoleClientRepository.CreateUserClientRole(userIdUuid, roleIdUuid, clientIdUuid)

	if !saveUserRoleClient {
		return &dto.RoleAssignmentResponse{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    "Gagal menetapkan pengguna ke peran",
		}
	}

	if save, _ := driver.AddGroupingPolicy(
		userIdUuid.String(),
		roleIdUuid.String(),
		clientIdUuid.String(),
		existsResp.User.FullName,
		role.RoleName,
		existsResp.Client.ClientName,
	); !save {
		return &dto.RoleAssignmentResponse{}, &responseDto.ApiErrorResponse{
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
