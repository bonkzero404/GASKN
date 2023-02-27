package implements

import (
	"github.com/bonkzero404/gaskn/app/utils"
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/role/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) repositories.RoleRepository {
	return &RoleRepository{
		DB: db,
	}
}

func (repository RoleRepository) CreateRole(role *stores.Role) *gorm.DB {
	return repository.DB.Create(&role)
}

func (repository RoleRepository) UpdateRoleById(role *stores.Role) *gorm.DB {
	return repository.DB.Save(&role)
}

func (repository RoleRepository) DeleteRoleById(role *stores.Role) *gorm.DB {
	return repository.DB.Delete(&role)
}

func (repository RoleRepository) GetRoleById(role *stores.Role, id string) *gorm.DB {
	return repository.DB.Take(&role, "id = ?", id)
}

func (repository RoleRepository) GetRoleList(role *[]stores.Role, c *fiber.Ctx) (*utils.Pagination, error) {
	var pagination utils.Pagination

	err := repository.DB.Scopes(utils.Paginate(role, &pagination, repository.DB, c)).Find(&role).Error

	return &pagination, err
}
