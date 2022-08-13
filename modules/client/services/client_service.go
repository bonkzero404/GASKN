package services

import (
	respModel "go-starterkit-project/domain/dto"
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/modules/client/domain/dto"
	"go-starterkit-project/modules/client/domain/interfaces"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
		UserId:            pUuid,
		IsActive:          true,
	}

	// Create client
	err := service.ClientRepository.CreateClient(&clientData)

	if err != nil {
		return &dto.ClientResponse{}, &respModel.ApiErrorResponse{
			StatusCode: fiber.StatusUnprocessableEntity,
			Message:    utils.Lang(c, "global:err:failed-unknown"),
		}
	}

	roleResponse := dto.ClientResponse{
		ID:                clientData.ID.String(),
		ClientName:        clientData.ClientName,
		ClientDescription: clientData.ClientDescription,
		IsActive:          clientData.IsActive,
	}

	return &roleResponse, nil
}
