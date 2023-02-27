package implements

import (
	"errors"
	"github.com/bonkzero404/gaskn/app/http"
	"github.com/bonkzero404/gaskn/app/translation"
	roleRepo "github.com/bonkzero404/gaskn/features/role/repositories"
	roleAssignInteract "github.com/bonkzero404/gaskn/features/role_assignment/interactors"
	"github.com/bonkzero404/gaskn/features/user/factories"
	"github.com/bonkzero404/gaskn/features/user/interactors"
	"github.com/bonkzero404/gaskn/features/user/repositories"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/database/stores"
	roleAssignDto "github.com/bonkzero404/gaskn/features/role_assignment/dto"
	"github.com/bonkzero404/gaskn/features/user/dto"
)

type UserClient struct {
	UserRepository           repositories.UserRepository
	UserActionCodeRepository repositories.UserActionCodeRepository
	UserInvitationRepository repositories.UserInvitationRepository
	RepositoryAggregate      repositories.RepositoryAggregate
	RoleClientRepository     roleRepo.RoleClientRepository
	ActionFactory            factories.ActionFactory
	RoleAssignment           roleAssignInteract.RoleAssignment
}

func NewUserClient(
	UserRepository repositories.UserRepository,
	UserActionCodeRepository repositories.UserActionCodeRepository,
	UserInvitationRepository repositories.UserInvitationRepository,
	RepositoryAggregate repositories.RepositoryAggregate,
	Factory factories.ActionFactory,
	RoleClientRepository roleRepo.RoleClientRepository,
	RoleAssignment roleAssignInteract.RoleAssignment,
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

func (repository UserClient) CreateUserInvitation(c *fiber.Ctx, req *dto.UserInvitationRequest, invitedByUser string) (*dto.UserInvitationResponse, error) {
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
	errUser := repository.UserRepository.FindUserByEmail(&user, req.Email).Error

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    translation.Lang(c, config.UserErrNotFound),
		}
	}

	// Check user if not active
	if !user.IsActive {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    translation.Lang(c, config.UserErrNotActive),
		}
	}

	// Check user invited
	errUserInvitedBy := repository.UserRepository.FindUserById(&userInviteBy, invitedByUser).Error

	if errors.Is(errUserInvitedBy, gorm.ErrRecordNotFound) {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    translation.Lang(c, config.UserErrNotFound),
		}
	}

	// Check role Client
	errRoleClient := repository.RoleClientRepository.GetRoleClientId(&roleClient, req.RoleId, clientId).Error

	if errors.Is(errRoleClient, gorm.ErrRecordNotFound) {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    translation.Lang(c, config.RoleErrNotExists),
		}
	}

	// Check action code by user if exists
	errActionCode := repository.UserActionCodeRepository.FindExistsActionCode(&userActionCode, user.ID.String(), stores.INVITATION_CODE).Error

	if errActionCode != nil {
		actCode, errActFactory := repository.ActionFactory.CreateInvitation(&user, req.Url, userInviteBy.FullName, roleClient.Role.RoleName, clientId)

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

		errInvitation := repository.UserInvitationRepository.CreateUserInvitation(&userInvitationNew)

		if errInvitation.Error != nil {
			return nil, errInvitation.Error
		}

		var resp = dto.UserInvitationResponse{
			InvitedBy:     userInviteBy.FullName,
			InvitedTo:     user.FullName,
			InvitedToRole: roleClient.Role.RoleName,
			ClientId:      clientId,
		}

		return &resp, nil
	}

	t := time.Now()

	// Check if expired can re-create invitation
	if userActionCode.ExpiredAt.Before(t) {
		actCode, errActFactory := repository.ActionFactory.CreateInvitation(&user, req.Url, userInviteBy.FullName, roleClient.Role.RoleName, clientId)

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

		errInvitation := repository.UserInvitationRepository.CreateUserInvitation(&userInvitationNew)

		if errInvitation.Error != nil {
			return nil, errInvitation.Error
		}

		var resp = dto.UserInvitationResponse{
			InvitedBy:     userInviteBy.FullName,
			InvitedTo:     user.FullName,
			InvitedToRole: roleClient.Role.RoleName,
			ClientId:      clientId,
		}

		return &resp, nil
	}

	return nil, &http.SetApiErrorResponse{
		StatusCode: fiber.StatusUnprocessableEntity,
		Message:    translation.Lang(c, config.UserErrInvited),
	}
}

func (repository UserClient) UserInviteAcceptance(c *fiber.Ctx, code string, accept stores.StatusInvitationType) (*dto.UserInvitationResponse, error) {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	var user stores.User
	var userAct stores.UserActionCode
	var userInvitation stores.UserInvitation
	var roleClient stores.RoleClient

	clientId := c.Params(config.Config("API_CLIENT_PARAM"))

	errUser := repository.UserRepository.FindUserById(&user, userId).Error

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    translation.Lang(c, config.UserErrNotFound),
		}
	}

	if !user.IsActive {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    translation.Lang(c, config.UserErrAlreadyActive),
		}
	}

	errAct := repository.UserActionCodeRepository.FindUserActionCode(&userAct, user.ID.String(), code).Error

	if errors.Is(errAct, gorm.ErrRecordNotFound) {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    translation.Lang(c, config.UserErrActivationNotFound),
		}
	}

	t := time.Now()

	if userAct.ExpiredAt.Before(t) {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusGone,
			Message:    translation.Lang(c, config.UserErrActivationExpired),
		}
	}

	errInvitation := repository.UserInvitationRepository.FindInvitationByActId(&userInvitation, userAct.ID.String()).Error

	if errors.Is(errInvitation, gorm.ErrRecordNotFound) {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    translation.Lang(c, config.UserErrActivationNotFound),
		}
	}

	// Check role Client
	errRoleClient := repository.RoleClientRepository.GetRoleClientById(&roleClient, userInvitation.RoleClientId.String(), clientId).Error

	if errors.Is(errRoleClient, gorm.ErrRecordNotFound) {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    translation.Lang(c, config.RoleErrNotExists),
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

		errUserInvite := repository.UserInvitationRepository.UpdateUserInvitation(&userInvitationUpdate).Error

		if errUserInvite != nil {
			return nil, &http.SetApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    errUserInvite.Error(),
			}
		}

		_, err := repository.RepositoryAggregate.UpdateActionCodeUsed(user.ID.String(), code)
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

			repository.UserInvitationRepository.CreateClientAssignment(&clientAssign)

			var assignPermit = &roleAssignDto.RoleUserAssignment{
				UserId: user.ID.String(),
				RoleId: roleClient.RoleId.String(),
			}

			_, errAssignPermit := repository.RoleAssignment.AssignUserPermission(c, assignPermit)

			if errAssignPermit != nil {
				return nil, &http.SetApiErrorResponse{
					StatusCode: fiber.StatusUnprocessableEntity,
					Message:    errAssignPermit.Error(),
				}
			}
		}

		var resp = dto.UserInvitationResponse{
			InvitedBy:     userInvitation.InvitedBy,
			InvitedTo:     user.FullName,
			InvitedToRole: roleClient.Role.RoleName,
			ClientId:      clientId,
		}

		return &resp, nil

	} else {
		return nil, &http.SetApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    translation.Lang(c, config.UserErrActivationNotFound),
		}
	}
}
