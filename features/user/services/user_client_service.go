package services

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"gaskn/config"
	"gaskn/database/stores"
	respModel "gaskn/dto"
	"gaskn/features/user/contracts"
	"gaskn/features/user/dto"
	"gaskn/features/user/services/factories"
	"gaskn/utils"
)

type UserClientService struct {
	UserRepository           contracts.UserRepository
	UserActionCodeRepository contracts.UserActionCodeRepository
	UserInvitationRepository contracts.UserInvitationRepository
	RepositoryAggregate      contracts.RepositoryAggregate
	ActionFactory            factories.ActionFactoryInterface
}

func NewUserClientService(
	UserRepository contracts.UserRepository,
	UserActtionCodeRepository contracts.UserActionCodeRepository,
	UserInvitationRepository contracts.UserInvitationRepository,
	RepositoryAggregate contracts.RepositoryAggregate,
	Factory factories.ActionFactoryInterface,
) contracts.UserClientService {
	return &UserClientService{
		UserRepository:           UserRepository,
		UserActionCodeRepository: UserActtionCodeRepository,
		UserInvitationRepository: UserInvitationRepository,
		RepositoryAggregate:      RepositoryAggregate,
		ActionFactory:            Factory,
	}
}

func (service UserClientService) CreateUserInvitation(c *fiber.Ctx, req *dto.UserInvitationRequest, invitedByUser string, actType stores.ActCodeType) (map[string]interface{}, error) {
	var user stores.User
	var userInviteBy stores.User
	var userInvitation stores.UserInvitation

	// Get client id from param url
	clientId := c.Params(config.Config("API_CLIENT_PARAM"))

	// Convert client id string to type UUID
	uuidClientId, _ := uuid.Parse(clientId)

	// Check user if exists
	errUser := service.UserRepository.FindUserByEmail(&user, req.Email).Error

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "user:err:user-not-found"),
		}
	}

	// Check user if not active
	if !user.IsActive {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "user:err:user-not-active"),
		}
	}

	// Check user invited
	errUserInvitedBy := service.UserRepository.FindUserById(&userInviteBy, invitedByUser).Error

	if errors.Is(errUserInvitedBy, gorm.ErrRecordNotFound) {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "user:err:user-not-found"),
		}
	}

	// Check invitation data is available
	checkInvitation := service.UserInvitationRepository.FindUserInvitation(&userInvitation, user.ID.String(), clientId)

	if checkInvitation.RowsAffected > 0 {
		return nil, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "user:err:user-invited"),
		}
	}

	actCode, errActFactory := service.ActionFactory.CreateInvitation(&user, req.Url, userInviteBy.FullName)

	if errActFactory != nil {
		return nil, errActFactory
	}

	userInvitationNew := stores.UserInvitation{
		UserId:           user.ID,
		ClientId:         uuidClientId,
		UserActionCodeId: actCode.ID,
		UrlFrontendMatch: req.Url,
		InvitedBy:        userInviteBy.FullName,
		Status:           stores.PENDING,
	}

	errInvitation := service.UserInvitationRepository.CreateUserInvitation(&userInvitationNew)

	if errInvitation != nil {
		return nil, errInvitation.Error
	}

	return nil, nil
}

func (service UserClientService) UserInviteAcceptance(c *fiber.Ctx, email string, code string, accept stores.StatusInvitationType) (*stores.UserInvitation, error) {
	var user stores.User
	var userAct stores.UserActionCode
	var userInvitation stores.UserInvitation

	clientId := c.Params(config.Config("API_CLIENT_PARAM"))

	// Convert client id string to type UUID
	uuidClientId, _ := uuid.Parse(clientId)

	errUser := service.UserRepository.FindUserByEmail(&user, email).Error

	if errors.Is(errUser, gorm.ErrRecordNotFound) {
		return &stores.UserInvitation{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "user:err:user-not-found"),
		}
	}

	if !user.IsActive {
		return &stores.UserInvitation{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "user:err:activate-already-active"),
		}
	}

	errAct := service.UserActionCodeRepository.FindUserActionCode(&userAct, user.ID.String(), code).Error

	if errors.Is(errAct, gorm.ErrRecordNotFound) {
		return &stores.UserInvitation{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "user:err:activation-not-found"),
		}
	}

	t := time.Now()

	if userAct.ExpiredAt.Before(t) {
		return &stores.UserInvitation{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusGone,
			Message:    utils.Lang(c, "user:err:activation-expired"),
		}
	}

	errInvitation := service.UserInvitationRepository.FindUserInvitation(&userInvitation, user.ID.String(), clientId).Error

	if errors.Is(errInvitation, gorm.ErrRecordNotFound) {
		return &stores.UserInvitation{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusNotFound,
			Message:    utils.Lang(c, "user:err:activation-not-found"),
		}
	}

	if accept == stores.APPROVED || accept == stores.REJECTED {

		userInvitationUpdate := stores.UserInvitation{
			ID:               userInvitation.ID,
			UserId:           user.ID,
			ClientId:         uuidClientId,
			UserActionCodeId: userAct.ID,
			UrlFrontendMatch: userInvitation.UrlFrontendMatch,
			InvitedBy:        userInvitation.InvitedBy,
			Status:           accept,
		}

		errUserInvite := service.UserInvitationRepository.UpdateUserInvitation(&userInvitationUpdate).Error

		if errUserInvite != nil {
			return &stores.UserInvitation{}, &respModel.ApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    errUserInvite.Error(),
			}
		}

		service.RepositoryAggregate.UpdateActionCodeUsed(user.ID.String(), code)

		if accept == stores.APPROVED {
			// Save to client assignment
			clientAssign := stores.ClientAssignment{
				ClientId: uuidClientId,
				UserId:   user.ID,
				IsActive: true,
			}

			service.UserInvitationRepository.CreateClientAssignment(&clientAssign)
		}

		return &userInvitationUpdate, nil
	} else {
		return &stores.UserInvitation{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "user:err:activation-not-found"),
		}
	}
}
