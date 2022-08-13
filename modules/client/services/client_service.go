package services

import (
	respModel "go-starterkit-project/domain/dto"
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/modules/client/domain/dto"
	"go-starterkit-project/modules/client/domain/interfaces"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

type ClientService struct {
	ClientRepository interfaces.ClientRepositoryInterface
}

func NewClientService(
	clientRepository interfaces.ClientRepositoryInterface,
) interfaces.ClientServiceInterface {
	return &ClientService{
		ClientRepository: clientRepository,
	}
}

func (service ClientService) CreateClient(c *fiber.Ctx, client *dto.ClientRequest, userId string) (*dto.ClientResponse, error) {
	pUuid, _ := uuid.Parse(userId)

	clientData := stores.Client{
		ClientName:        client.ClientName,
		ClientDescription: client.ClientDescription,
		ClientSlug:        slug.Make(client.ClientName),
		UserId:            pUuid,
		IsActive:          true,
	}

	// Create client
	err := service.ClientRepository.CreateClient(&clientData)

	if err != nil {
		return &dto.ClientResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, err.Error()),
		}
	}

	roleResponse := dto.ClientResponse{
		ID:                clientData.ID.String(),
		ClientName:        clientData.ClientName,
		ClientDescription: clientData.ClientDescription,
		ClientSlug:        clientData.ClientSlug,
		IsActive:          clientData.IsActive,
	}

	return &roleResponse, nil
}

func (service ClientService) UpdateClient(c *fiber.Ctx, id string, client *dto.ClientRequest) (*dto.ClientResponse, error) {
	// Check role if exists
	var clientStore stores.Client

	errCheckRole := service.ClientRepository.GetClientById(&clientStore, id).Error

	if errCheckRole != nil {
		return &dto.ClientResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "role:err:read:exists"),
		}
	}

	clientStore.ClientName = client.ClientName
	clientStore.ClientDescription = client.ClientDescription
	clientStore.ClientSlug = slug.Make(client.ClientName)
	clientStore.IsActive = true

	err := service.ClientRepository.UpdateClientById(&clientStore).Error

	if err != nil {
		return &dto.ClientResponse{}, &respModel.ApiErrorResponse{
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
