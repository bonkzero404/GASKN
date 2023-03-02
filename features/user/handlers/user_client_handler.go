package handlers

import (
	"github.com/bonkzero404/gaskn/app/http"
	"github.com/bonkzero404/gaskn/app/http/response"
	"github.com/bonkzero404/gaskn/app/translation"
	"github.com/bonkzero404/gaskn/app/validations"
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/features/user/dto"
	"github.com/bonkzero404/gaskn/features/user/interactors"
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
	var clientId = c.Params(config.Config("API_CLIENT_PARAM"))

	var token = c.Locals("user").(*jwt.Token)
	var claims = token.Claims.(jwt.MapClaims)
	var userIdInvitationBy = claims["id"].(string)

	if stat, errRequest := validations.ValidateRequest(c, &request); stat {
		return response.ApiUnprocessableEntity(c, errRequest)
	}

	responseInteract, err := interact.UserClientService.CreateUserInvitation(clientId, &request, userIdInvitationBy)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang(config.UserErrCreateActivation),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiCreated(c, responseInteract)
}

func (interact *UserClientHandler) UserInvitationAcceptance(c *fiber.Ctx) error {
	var request dto.UserInvitationApprovalRequest
	var clientId = c.Params(config.Config("API_CLIENT_PARAM"))

	var token = c.Locals("user").(*jwt.Token)
	var claims = token.Claims.(jwt.MapClaims)
	var userId = claims["id"].(string)

	if stat, errRequest := validations.ValidateRequest(c, &request); stat {
		return response.ApiUnprocessableEntity(c, errRequest)
	}

	var req = &request
	responseInteract, err := interact.UserClientService.UserInviteAcceptance(clientId, userId, req.Code, req.Status)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang(config.UserErrActivationFailed),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiCreated(c, responseInteract)
}
