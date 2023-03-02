package interactors

import (
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/features/client/dto"
)

type Client interface {
	CreateClient(client *dto.ClientRequest, userId string) (*dto.ClientResponse, error)

	GetClientByUser(userId string, page string, limit string, sort string) (*utils.Pagination, error)

	UpdateClient(clientId string, role *dto.ClientRequest) (*dto.ClientResponse, error)
}
