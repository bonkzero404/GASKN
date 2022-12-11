package interactors

import (
	"gaskn/features/client/dto"
	"gaskn/utils"

	"github.com/gofiber/fiber/v2"
)

type Client interface {
	CreateClient(c *fiber.Ctx, client *dto.ClientRequest, userId string) (*dto.ClientResponse, error)

	GetClientByUser(c *fiber.Ctx, userId string) (*utils.Pagination, error)

	UpdateClient(c *fiber.Ctx, role *dto.ClientRequest) (*dto.ClientResponse, error)
}
