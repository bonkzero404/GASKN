package handlers

import (
	"github.com/bonkzero404/gaskn/config"
	globalDto "github.com/bonkzero404/gaskn/dto"
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

func (interact *UserClientHandler) CreateUserInvitation(c *fiber.Ctx) error {
	var request dto.UserInvitationRequest

	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userIdInvitationBy := claims["id"].(string)

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	response, err := interact.UserClientService.CreateUserInvitation(c, &request, userIdInvitationBy)

	if err != nil {
		re := err.(*globalDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, globalDto.Errors{
			Message: utils.Lang(c, config.UserErrCreateActivation),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, &response)
}

func (interact *UserClientHandler) UserInvitationAcceptance(c *fiber.Ctx) error {
	var request dto.UserInvitationApprovalRequest

	if stat, errRequest := utils.ValidateRequest(c, &request); stat {
		return utils.ApiUnprocessableEntity(c, errRequest)
	}

	var req = &request
	response, err := interact.UserClientService.UserInviteAcceptance(c, req.Code, req.Status)

	if err != nil {
		re := err.(*globalDto.ApiErrorResponse)
		return utils.ApiResponseError(c, re.StatusCode, globalDto.Errors{
			Message: utils.Lang(c, config.UserErrActivationFailed),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return utils.ApiCreated(c, &response)
}
