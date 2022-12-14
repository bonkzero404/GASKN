package implements

import (
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/role/repositories"
	"github.com/bonkzero404/gaskn/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleClientRepository struct {
	DB *gorm.DB
}

func NewRoleClientRepository(db *gorm.DB) repositories.RoleClientRepository {
	return &RoleClientRepository{
		DB: db,
	}
}

func (repository RoleClientRepository) CreateRoleClient(role *stores.Role, clientId string) (*stores.Role, error) {
	uuidClientId, _ := uuid.Parse(clientId)

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
	if err := tx.Create(&role).Error; err != nil {
		tx.Rollback()
		return &stores.Role{}, err
	}

	roleClient := stores.RoleClient{
		RoleId:    role.ID,
		ClientId:  uuidClientId,
		IsActive:  true,
		CanDelete: true,
	}

	if err := tx.Create(&roleClient).Error; err != nil {
		tx.Rollback()
		return &stores.Role{}, err
	}

	return role, tx.Commit().Error
}

func (repository RoleClientRepository) CreateUserClientRole(userId uuid.UUID, roleId uuid.UUID, clientId uuid.UUID) bool {

	tx := repository.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return false
	}

	// Set Role
	// Set role User
	roleUser := stores.RoleUser{
		RoleId:   roleId,
		UserId:   userId,
		IsActive: true,
	}

	// Create Role
	if err := tx.Create(&roleUser).Error; err != nil {
		tx.Rollback()
		return false
	}

	// Set Role User with Client
	roleUserClient := stores.RoleUserClient{
		ClientId:   clientId,
		RoleUserId: roleUser.ID,
		IsActive:   true,
	}

	// Create Role
	if err := tx.Create(&roleUserClient).Error; err != nil {
		tx.Rollback()
		return false
	}

	if commitError := tx.Commit().Error; commitError != nil {
		return false
	}

	return true
}

func (repository RoleClientRepository) GetRoleClientById(roleClient *stores.RoleClient, id string, clientId string) *gorm.DB {
	return repository.DB.
		Preload("Client").
		Preload("Role").
		Take(&roleClient, "role_clients.id = ? and role_clients.client_id = ?", id, clientId)
}

func (repository RoleClientRepository) GetRoleClientByName(roleClient *stores.RoleClient, roleName string, clientId string) *gorm.DB {
	return repository.DB.
		Joins("left join roles on role_clients.role_id = roles.id").
		Preload("Client").
		First(&roleClient, "roles.role_name = ? and role_clients.client_id = ?", roleName, clientId)
}

func (repository RoleClientRepository) GetRoleClientList(role *[]stores.Role, c *fiber.Ctx, clientId string) (*utils.Pagination, error) {
	var pagination utils.Pagination

	err := repository.DB.Scopes(utils.Paginate(role, &pagination, repository.DB, c)).
		Joins("left join role_clients on role_clients.role_id = roles.id").
		Find(&role, "role_clients.client_id = ?", clientId).Error

	return &pagination, err
}

func (repository RoleClientRepository) GetRoleClientId(role *stores.RoleClient, roleId string, clientId string) *gorm.DB {
	return repository.DB.
		Preload("Client").
		Preload("Role").
		First(&role, "role_id = ? and client_id = ?", roleId, clientId)
}

func (repository RoleClientRepository) GetUserHasClient(clientAssignment *stores.ClientAssignment, userId string, clientId string) *gorm.DB {
	return repository.DB.
		Preload("Client").
		Preload("User").
		Take(&clientAssignment, "user_id = ? and client_id = ?", userId, clientId)
}
