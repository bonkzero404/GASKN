package implements

import (
	"gaskn/driver"
	"gaskn/features/client/interactors"
	"gaskn/features/client/repositories"
	userContract "gaskn/features/user/repositories"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/gosimple/slug"

	"gaskn/config"
	"gaskn/database/stores"
	responseDto "gaskn/dto"
	"gaskn/features/client/dto"
	"gaskn/utils"
)

type Client struct {
	ClientRepository repositories.ClientRepository
	UserRepository   userContract.UserRepository
}

func NewClient(
	clientRepository repositories.ClientRepository,
	UserRepository userContract.UserRepository,
) interactors.Client {
	return &Client{
		ClientRepository: clientRepository,
		UserRepository:   UserRepository,
	}
}

func (interact Client) CreateClient(c *fiber.Ctx, client *dto.ClientRequest, userId string) (*dto.ClientResponse, error) {
	pUuid, _ := uuid.Parse(userId)
	clientRoute := config.Config("API_WRAP") + "/" + config.Config("API_VERSION") + "/" + config.Config("API_CLIENT") + "/:" + config.Config("API_CLIENT_PARAM")

	clientStore := stores.Client{
		ClientName:        client.ClientName,
		ClientDescription: client.ClientDescription,
		ClientSlug:        slug.Make(client.ClientName),
		UserId:            pUuid,
		IsActive:          true,
	}

	// Create client
	role, err := interact.ClientRepository.CreateClient(&clientStore)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return &dto.ClientResponse{}, &responseDto.ApiErrorResponse{
				StatusCode: fiber.StatusUnprocessableEntity,
				Message:    utils.Lang(c, "client:err:duplicate"),
			}
		}

		return &dto.ClientResponse{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, err.Error()),
		}
	}

	// Get User By Id
	var user = stores.User{}
	interact.UserRepository.FindUserById(&user, pUuid.String())

	// Crete group policy
	if g, err := driver.AddGroupingPolicy(
		pUuid.String(),
		role.ID.String(),
		clientStore.ID.String(),
		user.FullName,
		role.RoleName,
		clientStore.ClientName,
	); !g {
		return &dto.ClientResponse{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, err.Error()),
		}
	}

	// Create permission user to group
	if p, err := driver.AddPolicy(
		role.ID.String(),
		clientStore.ID.String(),
		"/"+clientRoute+"/*",
		"GET|POST|PUT|DELETE",
		"",
		role.RoleName,
		clientStore.ClientName,
	); !p {
		return &dto.ClientResponse{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, err.Error()),
		}
	}

	roleResponse := dto.ClientResponse{
		ID:                clientStore.ID.String(),
		ClientName:        clientStore.ClientName,
		ClientDescription: clientStore.ClientDescription,
		ClientSlug:        clientStore.ClientSlug,
		IsActive:          clientStore.IsActive,
	}

	return &roleResponse, nil
}

func (interact Client) UpdateClient(c *fiber.Ctx, client *dto.ClientRequest) (*dto.ClientResponse, error) {
	// Check role if exists
	var clientStore stores.Client

	clientId := c.Params(config.Config("API_CLIENT_PARAM"))

	errCheckClient := interact.ClientRepository.GetClientById(&clientStore, clientId).Error

	if errCheckClient != nil {
		return &dto.ClientResponse{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "client:err:read-exists"),
		}
	}

	clientStore.ClientName = client.ClientName
	clientStore.ClientDescription = client.ClientDescription
	clientStore.ClientSlug = slug.Make(client.ClientName)
	clientStore.IsActive = true

	err := interact.ClientRepository.UpdateClientById(&clientStore).Error

	if err != nil {
		return &dto.ClientResponse{}, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:failed-unknown"),
		}
	}

	clientResponse := dto.ClientResponse{
		ID:                clientStore.ID.String(),
		ClientName:        clientStore.ClientName,
		ClientDescription: clientStore.ClientDescription,
		ClientSlug:        clientStore.ClientSlug,
		IsActive:          clientStore.IsActive,
	}

	return &clientResponse, nil
}

func (interact Client) GetClientByUser(c *fiber.Ctx, userId string) (*utils.Pagination, error) {
	var clientAssignment []stores.ClientAssignment
	var resp []dto.ClientResponse

	res, err := interact.ClientRepository.GetClientListByUser(&clientAssignment, c, userId)

	if err != nil {
		return nil, &responseDto.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    err.Error(),
		}
	}

	for _, item := range clientAssignment {
		resp = append(resp, dto.ClientResponse{
			ID:                item.Client.ID.String(),
			ClientName:        item.Client.ClientName,
			ClientDescription: item.Client.ClientDescription,
			ClientSlug:        item.Client.ClientSlug,
			IsActive:          item.IsActive,
		})
	}

	res.Rows = resp

	return res, nil

}
