package interfaces

import (
	"go-starterkit-project/modules/client/domain/dto"

	"github.com/gofiber/fiber/v2"
)

type ClientServiceInterface interface {
	CreateClient(c *fiber.Ctx, client *dto.ClientRequest, userId string) (*dto.ClientResponse, error)

	// GetClientList(c *fiber.Ctx, userId string) (*utils.Pagination, error)

	// UpdateClient(c *fiber.Ctx, id string, role *dto.ClientRequest, userId string) (*dto.ClientResponse, error)

	// DeleteClientById(c *fiber.Ctx, id string, userId string) (*dto.ClientResponse, error)
}
