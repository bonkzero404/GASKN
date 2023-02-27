package interactors

import (
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/features/client/dto"
	"github.com/gofiber/fiber/v2"
)

type Client interface {
	CreateClient(c *fiber.Ctx, client *dto.ClientRequest, userId string) (*dto.ClientResponse, error)

	GetClientByUser(c *fiber.Ctx, userId string) (*utils.Pagination, error)

	UpdateClient(c *fiber.Ctx, role *dto.ClientRequest) (*dto.ClientResponse, error)
}
