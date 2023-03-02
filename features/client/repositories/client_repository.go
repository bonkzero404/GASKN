package repositories

import (
	"github.com/bonkzero404/gaskn/app/utils"
	"gorm.io/gorm"

	"github.com/bonkzero404/gaskn/database/stores"
)

type ClientRepository interface {
	CreateClient(client *stores.Client) (*stores.Role, error)

	CreateClientAssignment(client *stores.ClientAssignment) *gorm.DB

	UpdateClientById(client *stores.Client) *gorm.DB

	DeleteClientById(client *stores.Client) *gorm.DB

	GetClientById(client *stores.Client, id string) *gorm.DB

	GetClientBySlug(client *stores.Client, slug string) *gorm.DB

	GetClientList(client *[]stores.Client, page int, limit int, sort string) (*utils.Pagination, error)

	GetClientListByUser(client *[]stores.ClientAssignment, userId string, page int, limit int, sort string) (*utils.Pagination, error)
}
