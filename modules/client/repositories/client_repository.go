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

func (repository ClientRepository) CreateClient(client *stores.Client) (*stores.Role, error) {
	var role stores.Role

	if err := repository.DB.First(&role, "role_name = 'Owner' AND role_type = 'cl'").Error; err != nil {
		return &stores.Role{}, err
	}

	tx := repository.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return &stores.Role{}, err
	}

	// Create CLient
	if err := tx.Create(&client).Error; err != nil {
		tx.Rollback()
		return &stores.Role{}, err
	}

	// Set Role
	// Set role User
	roleUser := stores.RoleUser{
		RoleId:   role.ID,
		UserId:   client.UserId,
		IsActive: true,
	}

	// Create Role
	if err := tx.Create(&roleUser).Error; err != nil {
		tx.Rollback()
		return &stores.Role{}, err
	}

	// Set Role User with Client
	roleUserClient := stores.RoleUserClient{
		ClientId:   client.ID,
		RoleUserId: roleUser.ID,
		IsActive:   true,
	}

	// Create Role
	if err := tx.Create(&roleUserClient).Error; err != nil {
		tx.Rollback()
		return &stores.Role{}, err
	}

	// Set Role Client
	roleClient := stores.RoleClient{
		ClientId:  client.ID,
		RoleId:    role.ID,
		CanDelete: false,
		IsActive:  true,
	}

	// Create Role Client
	if err := tx.Create(&roleClient).Error; err != nil {
		tx.Rollback()
		return &stores.Role{}, err
	}

	// Set user to client assignment
	clientAssignment := stores.ClientAssignment{
		ClientId: client.ID,
		UserId:   client.UserId,
		IsActive: true,
	}

	if err := tx.Create(&clientAssignment).Error; err != nil {
		tx.Rollback()
		return &stores.Role{}, err
	}

	return &role, tx.Commit().Error
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

func (repository ClientRepository) GetClientBySlug(client *stores.Client, slug string) *gorm.DB {
	return repository.DB.First(&client, "client_slug = ?", slug)
}

func (repository ClientRepository) GetClientList(client *[]stores.Client, c *fiber.Ctx) (*utils.Pagination, error) {
	var pagination utils.Pagination

	err := repository.DB.Scopes(utils.Paginate(client, &pagination, repository.DB, c)).Find(&client).Error

	return &pagination, err
}

func (repository ClientRepository) GetClientListByUser(clientAssignment *[]stores.ClientAssignment, c *fiber.Ctx, userId string) (*utils.Pagination, error) {
	var pagination utils.Pagination

	err := repository.DB.Scopes(utils.Paginate(clientAssignment, &pagination, repository.DB, c)).Preload("Client").Find(&clientAssignment, "user_id = ?", userId).Error

	return &pagination, err
}
