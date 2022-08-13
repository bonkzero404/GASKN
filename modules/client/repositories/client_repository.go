package repositories

import (
	"go-starterkit-project/domain/stores"
	"go-starterkit-project/modules/client/domain/interfaces"
	"go-starterkit-project/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ClientRepository struct {
	DB *gorm.DB
}

func NewClientRepository(db *gorm.DB) interfaces.ClientRepositoryInterface {
	return &ClientRepository{
		DB: db,
	}
}

func (repository ClientRepository) CreateClient(client *stores.Client) error {

	tx := repository.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// Create CLient
	if err := tx.Create(&client).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Set Role
	// Set role data
	role := stores.Role{
		RoleName:        "Owner",
		RoleDescription: "Role owner tenant",
		IsActive:        true,
		CanDelete:       true,
	}

	// Create Role
	if err := tx.Create(&role).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Set role assignments
	roleAssignment := stores.RoleUser{
		ClientId: client.ID,
		UserId:   client.UserId,
		RoleId:   role.ID,
		IsActive: true,
	}

	if err := tx.Create(&roleAssignment).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (repository ClientRepository) UpdateClientById(client *stores.Client) *gorm.DB {
	return repository.DB.Save(&client)
}

func (repository ClientRepository) DeleteClientById(client *stores.Client) *gorm.DB {
	return repository.DB.Delete(&client)
}

func (repository ClientRepository) GetClientById(client *stores.Client, id string) *gorm.DB {
	return repository.DB.First(&client, "id = ?", id)
}

func (repository ClientRepository) GetClientList(client *[]stores.Client, c *fiber.Ctx) (*utils.Pagination, error) {
	var pagination utils.Pagination

	err := repository.DB.Scopes(utils.Paginate(client, &pagination, repository.DB, c)).Find(&client).Error

	return &pagination, err
}
