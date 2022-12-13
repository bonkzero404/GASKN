package implements

import (
	"errors"
	repoRole "gaskn/features/role/repositories"
	interactRoleUserAssignment "gaskn/features/role_assignment/interactors"
	"gaskn/features/user/factories/implements"
	"gaskn/features/user/interactors"
	"gaskn/features/user/repositories"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"gaskn/config"
	"gaskn/database/stores"
	responseDto "gaskn/dto"
	dtoAssignment "gaskn/features/role_assignment/dto"
	"gaskn/features/user/dto"
	"gaskn/utils"
)

type UserClient struct {
	UserRepository           repositories.UserRepository
	UserActionCodeRepository repositories.UserActionCodeRepository
	UserInvitationRepository repositories.UserInvitationRepository
	RepositoryAggregate      repositories.RepositoryAggregate
	RoleClientRepository     repoRole.RoleClientRepository
	ActionFactory            implements.ActionFactoryInterface
	RoleAssignment           interactRoleUserAssignment.RoleAssignment
}

func NewUserClient(
	UserRepository repositories.UserRepository,
	UserActionCodeRepository repositories.UserActionCodeRepository,
	UserInvitationRepository repositories.UserInvitationRepository,
	RepositoryAggregate repositories.RepositoryAggregate,
	Factory implements.ActionFactoryInterface,
	RoleClientRepository repoRole.RoleClientRepository,
	RoleAssignment interactRoleUserAssignment.RoleAssignment,
) interactors.UserClient {
	return &UserClient{
		UserRepository:           UserRepository,
		UserActionCodeRepository: UserActionCodeRepository,
		UserInvitationRepository: UserInvitationRepository,
		RepositoryAggregate:      RepositoryAggregate,
		ActionFactory:            Factory,
		RoleClientRepository:     RoleClientRepository,
		RoleAssignment:           RoleAssignment,
	}
}

func (interact UserClient) CreateUserInvitation(c *fiber.Ctx, req *dto.UserInvitationRequest, invitedByUser string) (map[string]interface{}, error) {
	var user stores.User
	var userInviteBy stores.User
	// var userInvitation stores.UserInvitation
	var userActionCode stores.UserActionCode

	var roleClient stores.RoleClient

	// Get client id from param url
	clientId := c.Params(config.Config("API_CLIENT_PARAM"))

	// Convert client id string to type UUID
	uuidClientId, _ := uuid.Parse(clientId)

	// Check user if exists
	errUser := interact.UserRepository.FindUserByEmail(&user, req.Email).Error

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "user:err:user-not-found"),
		}
	}

	// Check user if not active
	if !user.IsActive {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "user:err:user-not-active"),
		}
	}

	// Check user invited
	errUserInvitedBy := interact.UserRepository.FindUserById(&userInviteBy, invitedByUser).Error

	if errors.Is(errUserInvitedBy, gorm.ErrRecordNotFound) {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "user:err:user-not-found"),
		}
	}

	// Check role Client
	errRoleClient := interact.RoleClientRepository.GetRoleClientId(&roleClient, req.RoleId, clientId).Error

	if errors.Is(errRoleClient, gorm.ErrRecordNotFound) {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "role:err:read-exists"),
		}
	}

	// Check action code by user if exists
	errActionCode := interact.UserActionCodeRepository.FindExistsActionCode(&userActionCode, user.ID.String(), stores.INVITATION_CODE).Error

	if errActionCode != nil {
		actCode, errActFactory := interact.ActionFactory.CreateInvitation(&user, req.Url, userInviteBy.FullName, roleClient.Role.RoleName, clientId)

		if errActFactory != nil {
			return nil, errActFactory
		}

		userInvitationNew := stores.UserInvitation{
			UserId:           user.ID,
			ClientId:         uuidClientId,
			UserActionCodeId: actCode.ID,
			UrlFrontendMatch: req.Url,
			InvitedBy:        userInviteBy.FullName,
			RoleClientId:     roleClient.ID,
			Role:             roleClient.Role.RoleName,
			Status:           stores.PENDING,
		}

		errInvitation := interact.UserInvitationRepository.CreateUserInvitation(&userInvitationNew)

		if errInvitation != nil {
			return nil, errInvitation.Error
		}

		return nil, nil
	}

	t := time.Now()

	// Check if expired can re-create invitation
	if userActionCode.ExpiredAt.Before(t) {
		actCode, errActFactory := interact.ActionFactory.CreateInvitation(&user, req.Url, userInviteBy.FullName, roleClient.Role.RoleName, clientId)

		if errActFactory != nil {
			return nil, errActFactory
		}

		userInvitationNew := stores.UserInvitation{
			UserId:           user.ID,
			ClientId:         uuidClientId,
			UserActionCodeId: actCode.ID,
			UrlFrontendMatch: req.Url,
			InvitedBy:        userInviteBy.FullName,
			RoleClientId:     roleClient.ID,
			Role:             roleClient.Role.RoleName,
			Status:           stores.PENDING,
		}

		errInvitation := interact.UserInvitationRepository.CreateUserInvitation(&userInvitationNew)

		if errInvitation != nil {
			return nil, errInvitation.Error
		}

		return nil, nil
	}

	return nil, &responseDto.ApiErrorResponse{
		StatusCode: fiber.StatusUnprocessableEntity,
		Message:    utils.Lang(c, "user:err:user-invited"),
	}
}

func (interact UserClient) UserInviteAcceptance(c *fiber.Ctx, code string, accept stores.StatusInvitationType) (*stores.UserInvitation, error) {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	var user stores.User
	var userAct stores.UserActionCode
	var userInvitation stores.UserInvitation
	var roleClient stores.RoleClient

	clientId := c.Params(config.Config("API_CLIENT_PARAM"))

	errUser := interact.UserRepository.FindUserById(&user, userId).Error

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return &stores.UserInvitation{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "user:err:user-not-found"),
		}
	}

	if !user.IsActive {
		return &stores.UserInvitation{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "user:err:activate-already-active"),
		}
	}

	errAct := interact.UserActionCodeRepository.FindUserActionCode(&userAct, user.ID.String(), code).Error

	if errors.Is(errAct, gorm.ErrRecordNotFound) {
		return &stores.UserInvitation{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "user:err:activation-not-found"),
		}
	}

	t := time.Now()

	if userAct.ExpiredAt.Before(t) {
		return &stores.UserInvitation{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusGone,
			Message:    utils.Lang(c, "user:err:activation-expired"),
		}
	}

	errInvitation := interact.UserInvitationRepository.FindInvitationByActId(&userInvitation, userAct.ID.String()).Error

	if errors.Is(errInvitation, gorm.ErrRecordNotFound) {
		return &stores.UserInvitation{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "user:err:activation-not-found"),
		}
	}

	// Check role Client
	errRoleClient := interact.RoleClientRepository.GetRoleClientById(&roleClient, userInvitation.RoleClientId.String(), clientId).Error

	if errors.Is(errRoleClient, gorm.ErrRecordNotFound) {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "role:err:read-exists"),
		}
	}

	if accept == stores.APPROVED || accept == stores.REJECTED {

		userInvitationUpdate := stores.UserInvitation{
			ID:               userInvitation.ID,
			UserId:           user.ID,
			ClientId:         userInvitation.ClientId,
			UserActionCodeId: userAct.ID,
			UrlFrontendMatch: userInvitation.UrlFrontendMatch,
			InvitedBy:        userInvitation.InvitedBy,
			RoleClientId:     roleClient.ID,
			Role:             roleClient.Role.RoleName,
			Status:           accept,
		}

		errUserInvite := interact.UserInvitationRepository.UpdateUserInvitation(&userInvitationUpdate).Error

		if errUserInvite != nil {
			return &stores.UserInvitation{}, &responseDto.ApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    errUserInvite.Error(),
			}
		}

		_, err := interact.RepositoryAggregate.UpdateActionCodeUsed(user.ID.String(), code)
		if err != nil {
			return nil, err
		}

		if accept == stores.APPROVED {
			// Save to client assignment
			clientAssign := stores.ClientAssignment{
				ClientId: userInvitation.ClientId,
				UserId:   user.ID,
				IsActive: true,
			}

			interact.UserInvitationRepository.CreateClientAssignment(&clientAssign)

			var assignPermit = &dtoAssignment.RoleUserAssignment{
				UserId: user.ID.String(),
				RoleId: roleClient.RoleId.String(),
			}

			_, errAssignPermit := interact.RoleAssignment.AssignUserPermitToRole(c, assignPermit)

			if errAssignPermit != nil {
				return &stores.UserInvitation{}, &responseDto.ApiErrorResponse{
					StatusCode: fiber.StatusUnprocessableEntity,
					Message:    errAssignPermit.Error(),
				}
			}
		}

		return &userInvitationUpdate, nil
	} else {
		return &stores.UserInvitation{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "user:err:activation-not-found"),
		}
	}
}
