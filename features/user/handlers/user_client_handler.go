package handlers

import (
	responseDto "github.com/bonkzero404/gaskn/dto"
	"github.com/bonkzero404/gaskn/features/user/dto"
	"github.com/bonkzero404/gaskn/features/user/interactors"
	"github.com/bonkzero404/gaskn/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type UserClientHandler struct {
	UserClientService interactors.UserClient
}

func NewUserClientHandler(UserClientService interactors.UserClient) *UserClientHandler {
	return &UserClientHandler{
		UserClientService: UserClientService,
	}
}

func (service *UserClientHandler) CreateUserInvitation(c *fiber.Ctx) error {
	var request dto.UserInvitationRequest

	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userIdInvitationBy := claims["id"].(string)

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	response, err := service.UserClientService.CreateUserInvitation(c, &request, userIdInvitationBy)

	if err != nil {
		re := err.(*responseDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, responseDto.Errors{
			Message: utils.Lang(c, "user:err:create-invitation"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, &response)
}

func (service *UserClientHandler) UserInvitationAcceptance(c *fiber.Ctx) error {
	var request dto.UserInvitationApprovalRequest

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	var req = &request
	response, err := service.UserClientService.UserInviteAcceptance(c, req.Code, req.Status)

	if err != nil {
		re := err.(*responseDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, responseDto.Errors{
			Message: utils.Lang(c, "user:err:activation-failed"),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, &response)
}
