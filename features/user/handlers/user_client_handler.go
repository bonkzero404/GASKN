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

	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userIdInvitationBy := claims["id"].(string)

	if stat, errRequest := validations.ValidateRequest(c, &request); stat {
		return response.ApiUnprocessableEntity(c, errRequest)
	}

	responseInteract, err := interact.UserClientService.CreateUserInvitation(c, &request, userIdInvitationBy)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang(c, config.UserErrCreateActivation),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiCreated(c, responseInteract)
}

func (interact *UserClientHandler) UserInvitationAcceptance(c *fiber.Ctx) error {
	var request dto.UserInvitationApprovalRequest

	if stat, errRequest := validations.ValidateRequest(c, &request); stat {
		return response.ApiUnprocessableEntity(c, errRequest)
	}

	var req = &request
	responseInteract, err := interact.UserClientService.UserInviteAcceptance(c, req.Code, req.Status)

	if err != nil {
		re := err.(*http.SetApiErrorResponse)
		return response.ApiResponseError(c, re.StatusCode, http.SetErrors{
			Message: translation.Lang(c, config.UserErrActivationFailed),
			Cause:   err.Error(),
			Inputs:  nil,
		})
	}

	return response.ApiCreated(c, responseInteract)
}
