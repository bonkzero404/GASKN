package contracts

import (
	"go-starterkit-project/modules/client/dto"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
)

type ClientServiceInterface interface {
	CreateClient(c *fiber.Ctx, client *dto.ClientRequest, userId string) (*dto.ClientResponse, error)

	GetClientByUser(c *fiber.Ctx, userId string) (*utils.Pagination, error)

	UpdateClient(c *fiber.Ctx, role *dto.ClientRequest) (*dto.ClientResponse, error)

	// GetClientList(c *fiber.Ctx, userId string) (*utils.Pagination, error)

	// DeleteClientById(c *fiber.Ctx, id string, userId string) (*dto.ClientResponse, error)
}
