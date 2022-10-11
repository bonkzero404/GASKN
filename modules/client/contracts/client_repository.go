package contracts

import (
	"go-starterkit-project/database/stores"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ClientRepository interface {
	CreateClient(client *stores.Client) (*stores.Role, error)

	UpdateClientById(client *stores.Client) *gorm.DB

	DeleteClientById(client *stores.Client) *gorm.DB

	GetClientById(client *stores.Client, id string) *gorm.DB

	GetClientBySlug(client *stores.Client, slug string) *gorm.DB

	GetClientList(client *[]stores.Client, c *fiber.Ctx) (*utils.Pagination, error)

	GetClientListByUser(client *[]stores.ClientAssignment, c *fiber.Ctx, userId string) (*utils.Pagination, error)
}
