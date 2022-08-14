package interfaces

import (
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ClientRepositoryInterface interface {
	CreateClient(client *stores.Client) (*stores.Role, error)

	UpdateClientById(client *stores.Client) *gorm.DB

	DeleteClientById(client *stores.Client) *gorm.DB

	GetClientById(client *stores.Client, id string) *gorm.DB

	GetClientBySlug(client *stores.Client, slug string) *gorm.DB

	GetClientList(client *[]stores.Client, c *fiber.Ctx) (*utils.Pagination, error)
}
