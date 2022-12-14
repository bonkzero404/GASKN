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
	"strings"

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
		return nil, &responseDto.ApiErrorResponse{
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
		return nil, &responseDto.ApiErrorResponse{
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

	if clientId != "" && !strings.Contains(req.RouteFeature, config.Config("API_CLIENT_PARAM")) {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:not-allowed"),
		}
	}

	if clientId != "" && (errRoleUuid != nil || errClientIdUuid != nil) {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:invalid-format"),
		}
	} else if errRoleUuid != nil {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:invalid-format"),
		}
	}

	if clientId != "" {

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
			return nil, &responseDto.ApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    utils.Lang(c, "role-assign:err:failed-unknown"),
			}
		}

		saveResponse := dto.RoleAssignmentResponse{
			RoleId:     roleIdUuid.String(),
			ClientName: existsResp.Client.ClientName,
			UserName:   "-",
		}

		return &saveResponse, nil
	}

	// Check Role if exists
	var role = stores.Role{}

	errExistsRole := interact.RoleRepository.GetRoleById(&role, req.RoleId).Error

	if errExistsRole != nil {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "role:err:read-exists"),
		}
	}

	if save, _ := driver.AddPolicy(
		roleIdUuid.String(),
		"*",
		req.RouteFeature,
		req.MethodFeature,
		"",
		role.RoleName,
		"",
	); !save {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "role-assign:err:failed-unknown"),
		}
	}

	saveResponse := dto.RoleAssignmentResponse{
		RoleId:     roleIdUuid.String(),
		ClientName: config.Config("APP_NAME"),
		UserName:   "-",
	}

	return &saveResponse, nil
}

func (interact RoleAssignment) RemoveRoleAssignment(c *fiber.Ctx, req *dto.RoleAssignmentRequest) (*dto.RoleAssignmentResponse, error) {
	clientId := c.Params(config.Config("API_CLIENT_PARAM"))
	clientIdUuid, errClientIdUuid := uuid.Parse(clientId)
	roleIdUuid, errRoleUuid := uuid.Parse(req.RoleId)

	if clientId != "" && (errRoleUuid != nil || errClientIdUuid != nil) {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:invalid-format"),
		}
	} else if errRoleUuid != nil {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:invalid-format"),
		}
	}

	if clientId != "" {
		res, errExists := interact.CheckExistsRoleAssignment(c, clientIdUuid, roleIdUuid)

		if errExists != nil {
			return nil, errExists
		}

		if remove, _ := driver.RemovePolicy(
			roleIdUuid.String(),
			clientIdUuid.String(),
			req.RouteFeature,
			req.MethodFeature,
		); !remove {
			return nil, &responseDto.ApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    utils.Lang(c, "role-assign:err:failed-remove-permit"),
			}
		}

		saveResponse := dto.RoleAssignmentResponse{
			RoleId:     roleIdUuid.String(),
			ClientName: res.Client.ClientName,
			UserName:   "-",
		}

		return &saveResponse, nil
	}

	// Check Role if exists
	var role = stores.Role{}

	errExistsRole := interact.RoleRepository.GetRoleById(&role, req.RoleId).Error

	if errExistsRole != nil {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "role:err:read-exists"),
		}
	}

	if remove, _ := driver.RemovePolicy(
		roleIdUuid.String(),
		"*",
		req.RouteFeature,
		req.MethodFeature,
	); !remove {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "role-assign:err:failed-remove-permit"),
		}
	}

	saveResponse := dto.RoleAssignmentResponse{
		RoleId:     roleIdUuid.String(),
		ClientName: config.Config("APP_NAME"),
		UserName:   "-",
	}

	return &saveResponse, nil
}

func (interact RoleAssignment) AssignUserPermitToRole(c *fiber.Ctx, req *dto.RoleUserAssignment) (*dto.RoleAssignmentResponse, error) {
	clientId := c.Params(config.Config("API_CLIENT_PARAM"))
	clientIdUuid, errClientIdUuid := uuid.Parse(clientId)
	userIdUuid, errUserUuid := uuid.Parse(req.UserId)
	roleIdUuid, errRoleUuid := uuid.Parse(req.RoleId)

	if clientId != "" && (errRoleUuid != nil || errClientIdUuid != nil || errUserUuid != nil) {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:invalid-format"),
		}
	} else if errRoleUuid != nil || errUserUuid != nil {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:invalid-format"),
		}
	}

	// Check if role has available
	var role = stores.Role{}
	errRole := interact.RoleRepository.GetRoleById(&role, roleIdUuid.String()).Error

	if errRole != nil {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "role:err:read-exists"),
		}
	}

	var roleUser = stores.RoleUser{}
	errRoleUser := interact.RoleClientRepository.GetRoleUser(&roleUser, req.UserId, req.RoleId).Error

	// Check if role user is already exists
	if errRoleUser == nil {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "role-assign:err:exists"),
		}
	}

	// Check if user has client
	if clientId != "" {
		existsResp, errExists := interact.CheckExistsRoleUserAssignment(c, userIdUuid, clientIdUuid)

		if errExists != nil {
			return nil, errExists
		}

		saveUserRoleClient := interact.RoleClientRepository.CreateUserClientRole(userIdUuid, roleIdUuid, clientIdUuid)

		if !saveUserRoleClient {
			return nil, &responseDto.ApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    utils.Lang(c, "role-assign:err:failed"),
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
			return nil, &responseDto.ApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    utils.Lang(c, "role-assign:err:failed-unknown"),
			}
		}

		saveResponse := dto.RoleAssignmentResponse{
			RoleId:     roleIdUuid.String(),
			ClientName: existsResp.Client.ClientName,
			UserName:   existsResp.User.FullName,
		}

		return &saveResponse, nil
	}

	saveUserRoleClient := interact.RoleClientRepository.CreateUserClientRole(userIdUuid, roleIdUuid, clientIdUuid)

	if !saveUserRoleClient {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "role-assign:err:failed"),
		}
	}

	var roleUserGet = stores.RoleUser{}
	interact.RoleClientRepository.GetRoleUser(&roleUserGet, req.UserId, req.RoleId)

	if save, _ := driver.AddGroupingPolicy(
		userIdUuid.String(),
		roleIdUuid.String(),
		"*",
		roleUserGet.User.FullName,
		role.RoleName,
		"",
	); !save {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "role-assign:err:failed-unknown"),
		}
	}

	saveResponse := dto.RoleAssignmentResponse{
		RoleId:     roleIdUuid.String(),
		ClientName: config.Config("APP_NAME"),
		UserName:   roleUserGet.User.FullName,
	}

	return &saveResponse, nil
}
