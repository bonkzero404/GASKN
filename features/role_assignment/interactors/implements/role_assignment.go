package implements

import (
	"errors"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/driver"
	globalDto "github.com/bonkzero404/gaskn/dto"
	roleRepo "github.com/bonkzero404/gaskn/features/role/repositories"
	"github.com/bonkzero404/gaskn/features/role_assignment/dto"
	"github.com/bonkzero404/gaskn/features/role_assignment/interactors"
	roleAssignmentRepo "github.com/bonkzero404/gaskn/features/role_assignment/repositories"
	"github.com/bonkzero404/gaskn/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleAssignment struct {
	RoleClientRepository     roleRepo.RoleClientRepository
	RoleRepository           roleRepo.RoleRepository
	RoleAssignmentRepository roleAssignmentRepo.RoleAssignmentRepository
}

func NewRoleAssignment(
	RoleClientRepository roleRepo.RoleClientRepository,
	RoleRepository roleRepo.RoleRepository,
	RoleAssignmentRepository roleAssignmentRepo.RoleAssignmentRepository,
) interactors.RoleAssignment {
	return &RoleAssignment{
		RoleClientRepository:     RoleClientRepository,
		RoleRepository:           RoleRepository,
		RoleAssignmentRepository: RoleAssignmentRepository,
	}
}

func (repository RoleAssignment) CheckExistsRoleAssignment(c *fiber.Ctx, clientIdUuid uuid.UUID, roleIdUuid uuid.UUID) (*stores.RoleClient, error) {
	var clientRole = stores.RoleClient{}

	errRoleClient := repository.RoleClientRepository.GetRoleClientId(&clientRole, roleIdUuid.String(), clientIdUuid.String()).Error

	if errors.Is(errRoleClient, gorm.ErrRecordNotFound) {
		return nil, &globalDto.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, config.RoleErrNotExists),
		}
	}

	return &clientRole, nil
}

func (repository RoleAssignment) CheckExistsRoleUserAssignment(c *fiber.Ctx, userId uuid.UUID, clientIdUuid uuid.UUID) (*stores.ClientAssignment, error) {
	var clientAssign = stores.ClientAssignment{}

	errRoleUserClient := repository.RoleClientRepository.GetUserHasClient(
		&clientAssign,
		userId.String(),
		clientIdUuid.String(),
	).Error

	if errors.Is(errRoleUserClient, gorm.ErrRecordNotFound) {
		return nil, &globalDto.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, config.RoleErrNotExists),
		}
	}

	return &clientAssign, nil
}

// CreateRoleAssignment /**
func (repository RoleAssignment) CreateRoleAssignment(c *fiber.Ctx, req *dto.RoleAssignmentRequest) (*dto.RoleAssignmentResponse, error) {
	clientId := c.Params(config.Config("API_CLIENT_PARAM"))
	clientIdUuid, errClientIdUuid := uuid.Parse(clientId)
	roleIdUuid, errRoleUuid := uuid.Parse(req.RoleId)

	if clientId != "" && !strings.Contains(req.RouteFeature, config.Config("API_CLIENT_PARAM")) {
		return nil, &globalDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.GlobalErrNotAllowed),
		}
	}

	if clientId != "" && (errRoleUuid != nil || errClientIdUuid != nil) {
		return nil, &globalDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.GlobalErrInvalidFormat),
		}
	}

	if errRoleUuid != nil {
		return nil, &globalDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.GlobalErrInvalidFormat),
		}
	}

	if clientId != "" {

		existsResp, errExists := repository.CheckExistsRoleAssignment(c, clientIdUuid, roleIdUuid)

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
			req.RouteGroup,
			req.RouteName,
			req.DescriptionKeyLang,
		); !save {
			return nil, &globalDto.ApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    utils.Lang(c, config.RoleAssignErrUnknown),
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

	errExistsRole := repository.RoleRepository.GetRoleById(&role, req.RoleId).Error

	if errExistsRole != nil {
		return nil, &globalDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.RoleErrNotExists),
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
		req.RouteGroup,
		req.RouteName,
		req.DescriptionKeyLang,
	); !save {
		return nil, &globalDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.RoleAssignErrUnknown),
		}
	}

	saveResponse := dto.RoleAssignmentResponse{
		RoleId:     roleIdUuid.String(),
		ClientName: config.Config("APP_NAME"),
		UserName:   "-",
	}

	return &saveResponse, nil
}

func (repository RoleAssignment) RemoveRoleAssignment(c *fiber.Ctx, req *dto.RoleAssignmentRequest) (*dto.RoleAssignmentResponse, error) {
	clientId := c.Params(config.Config("API_CLIENT_PARAM"))
	clientIdUuid, errClientIdUuid := uuid.Parse(clientId)
	roleIdUuid, errRoleUuid := uuid.Parse(req.RoleId)

	if clientId != "" && (errRoleUuid != nil || errClientIdUuid != nil) {
		return nil, &globalDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.GlobalErrInvalidFormat),
		}
	}

	if errRoleUuid != nil {
		return nil, &globalDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.GlobalErrInvalidFormat),
		}
	}

	if clientId != "" {
		res, errExists := repository.CheckExistsRoleAssignment(c, clientIdUuid, roleIdUuid)

		if errExists != nil {
			return nil, errExists
		}

		if remove, _ := driver.RemovePolicy(
			roleIdUuid.String(),
			clientIdUuid.String(),
			req.RouteFeature,
			req.MethodFeature,
		); !remove {
			return nil, &globalDto.ApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    utils.Lang(c, config.RoleAssignErrRemovePermit),
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

	errExistsRole := repository.RoleRepository.GetRoleById(&role, req.RoleId).Error

	if errExistsRole != nil {
		return nil, &globalDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.RoleErrNotExists),
		}
	}

	if remove, _ := driver.RemovePolicy(
		roleIdUuid.String(),
		"*",
		req.RouteFeature,
		req.MethodFeature,
	); !remove {
		return nil, &globalDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.RoleAssignErrRemovePermit),
		}
	}

	saveResponse := dto.RoleAssignmentResponse{
		RoleId:     roleIdUuid.String(),
		ClientName: config.Config("APP_NAME"),
		UserName:   "-",
	}

	return &saveResponse, nil
}

func (repository RoleAssignment) AssignUserPermission(c *fiber.Ctx, req *dto.RoleUserAssignment) (*dto.RoleAssignmentResponse, error) {
	clientId := c.Params(config.Config("API_CLIENT_PARAM"))
	clientIdUuid, errClientIdUuid := uuid.Parse(clientId)
	userIdUuid, errUserUuid := uuid.Parse(req.UserId)
	roleIdUuid, errRoleUuid := uuid.Parse(req.RoleId)

	if clientId != "" && (errRoleUuid != nil || errClientIdUuid != nil || errUserUuid != nil) {
		return nil, &globalDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.GlobalErrInvalidFormat),
		}
	}

	if errRoleUuid != nil || errUserUuid != nil {
		return nil, &globalDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.GlobalErrInvalidFormat),
		}
	}

	// Check if role has available
	var role = stores.Role{}
	errRole := repository.RoleRepository.GetRoleById(&role, roleIdUuid.String()).Error

	if errRole != nil {
		return nil, &globalDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.RoleErrNotExists),
		}
	}

	var roleUser = stores.RoleUser{}
	errRoleUser := repository.RoleClientRepository.GetRoleUser(&roleUser, req.UserId, req.RoleId).Error

	// Check if role user is already exists
	if errRoleUser == nil {
		return nil, &globalDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.RoleAssignErrAlreadyExists),
		}
	}

	// Check if user has client
	if clientId != "" {
		existsResp, errExists := repository.CheckExistsRoleUserAssignment(c, userIdUuid, clientIdUuid)

		if errExists != nil {
			return nil, errExists
		}

		saveUserRoleClient := repository.RoleClientRepository.CreateUserClientRole(userIdUuid, roleIdUuid, clientIdUuid)

		if !saveUserRoleClient {
			return nil, &globalDto.ApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    utils.Lang(c, config.RoleAssignErrFailed),
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
			return nil, &globalDto.ApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    utils.Lang(c, config.RoleAssignErrUnknown),
			}
		}

		saveResponse := dto.RoleAssignmentResponse{
			RoleId:     roleIdUuid.String(),
			ClientName: existsResp.Client.ClientName,
			UserName:   existsResp.User.FullName,
		}

		return &saveResponse, nil
	}

	saveUserRoleClient := repository.RoleClientRepository.CreateUserClientRole(userIdUuid, roleIdUuid, clientIdUuid)

	if !saveUserRoleClient {
		return nil, &globalDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.RoleAssignErrFailed),
		}
	}

	var roleUserGet = stores.RoleUser{}
	repository.RoleClientRepository.GetRoleUser(&roleUserGet, req.UserId, req.RoleId)

	if save, _ := driver.AddGroupingPolicy(
		userIdUuid.String(),
		roleIdUuid.String(),
		"*",
		roleUserGet.User.FullName,
		role.RoleName,
		"",
	); !save {
		return nil, &globalDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.RoleAssignErrUnknown),
		}
	}

	saveResponse := dto.RoleAssignmentResponse{
		RoleId:     roleIdUuid.String(),
		ClientName: config.Config("APP_NAME"),
		UserName:   roleUserGet.User.FullName,
	}

	return &saveResponse, nil
}

func (repository RoleAssignment) GetPermissionListByRole(c *fiber.Ctx) (*[]dto.RoleAssignmentListResponse, error) {
	var clientId = c.Params(config.Config("API_CLIENT_PARAM"))
	var roleId = c.Params("RoleId")
	var permissionRule []stores.PermissionRuleDetail
	var resp []dto.RoleAssignmentListResponse

	if roleId == "" {
		return nil, &globalDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.GlobalErrInvalidFormat),
		}
	}

	if clientId == "" {
		clientId = "*"
	}

	err := repository.RoleAssignmentRepository.GetPermissionByRole(&permissionRule, roleId, clientId).Error

	if err != nil {
		return nil, &globalDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, config.RoleAssignErrLoad),
		}
	}

	for _, item := range permissionRule {
		resp = append(resp, dto.RoleAssignmentListResponse{
			ID:           item.ID.String(),
			PermissionId: item.PermissionRuleId,
			RoleId:       item.PermissionRule.V0,
			ClientName:   item.ClientName,
			RoleName:     item.RoleName,
			GroupName:    item.GroupName,
			RouteName:    item.RouteName,
			Description:  utils.Lang(c, item.DescriptionKeyLang),
			Route:        item.PermissionRule.V2,
			Method:       item.PermissionRule.V3,
		})
	}

	return &resp, nil
}
