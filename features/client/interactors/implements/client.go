package implements

import (
	"github.com/bonkzero404/gaskn/app/facades"
	"github.com/bonkzero404/gaskn/app/translation"
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/features/client/interactors"
	"github.com/bonkzero404/gaskn/features/client/repositories"
	userRepo "github.com/bonkzero404/gaskn/features/user/repositories"
	"github.com/bonkzero404/gaskn/infrastructures"
	"strings"

	"github.com/google/uuid"
	"github.com/gosimple/slug"

	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/client/dto"
)

type Client struct {
	ClientRepository repositories.ClientRepository
	UserRepository   userRepo.UserRepository
}

func NewClient(
	clientRepository repositories.ClientRepository,
	UserRepository userRepo.UserRepository,
) interactors.Client {
	return &Client{
		ClientRepository: clientRepository,
		UserRepository:   UserRepository,
	}
}

func (repository Client) CreateClient(client *dto.ClientRequest, userId string) (*dto.ClientResponse, error) {
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
	role, err := repository.ClientRepository.CreateClient(&clientStore)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, &facades.ResponseError{
				StatusCode: facades.AppErrUnprocessable,
				Message:    translation.Lang(config.ClientErrDuplicate),
			}
		}

		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrUnprocessable,
			Message:    translation.Lang(err.Error()),
		}
	}

	// Get User By Id
	var user = stores.User{}
	repository.UserRepository.FindUserById(&user, pUuid.String())

	// Crete group policy
	if g, err := infrastructures.AddGroupingPolicy(
		pUuid.String(),
		role.ID.String(),
		clientStore.ID.String(),
		user.FullName,
		role.RoleName,
		clientStore.ClientName,
	); !g {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrUnprocessable,
			Message:    translation.Lang(err.Error()),
		}
	}

	// Create permission user to group
	if p, err := infrastructures.AddPolicy(
		role.ID.String(),
		clientStore.ID.String(),
		"/"+clientRoute+"/*",
		"GET|POST|PUT|DELETE",
		"",
		role.RoleName,
		clientStore.ClientName,
		"",
		"",
		"",
	); !p {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrUnprocessable,
			Message:    translation.Lang(err.Error()),
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

func (repository Client) UpdateClient(clientId string, client *dto.ClientRequest) (*dto.ClientResponse, error) {
	// Check role if exists
	var clientStore stores.Client

	errCheckClient := repository.ClientRepository.GetClientById(&clientStore, clientId).Error

	if errCheckClient != nil {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrUnprocessable,
			Message:    translation.Lang(config.ClientErrAlreadyExists),
		}
	}

	clientStore.ClientName = client.ClientName
	clientStore.ClientDescription = client.ClientDescription
	clientStore.ClientSlug = slug.Make(client.ClientName)
	clientStore.IsActive = true

	err := repository.ClientRepository.UpdateClientById(&clientStore).Error

	if err != nil {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrUnprocessable,
			Message:    translation.Lang(config.GlobalErrUnknown),
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

func (repository Client) GetClientByUser(userId string, page string, limit string, sort string) (*utils.Pagination, error) {
	var clientAssignment []stores.ClientAssignment
	var resp []dto.ClientResponse
	paginateRequest := utils.PaginationRequest{
		Page:  page,
		Limit: limit,
		Sort:  sort,
	}

	pPage, pLimit, pSort := paginateRequest.SetPagination()

	res, err := repository.ClientRepository.GetClientListByUser(&clientAssignment, userId, pPage, pLimit, pSort)

	if err != nil {
		return nil, &facades.ResponseError{
			StatusCode: facades.AppErrUnprocessable,
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
