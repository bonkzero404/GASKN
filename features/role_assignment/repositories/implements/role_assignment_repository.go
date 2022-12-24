package implements

import (
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/role_assignment/repositories"
	"gorm.io/gorm"
)

type RoleAssignmentRepository struct {
	DB *gorm.DB
}

func NewRoleAssignmentRepository(db *gorm.DB) repositories.RoleAssignmentRepository {
	return &RoleAssignmentRepository{
		DB: db,
	}
}

func (repository RoleAssignmentRepository) GetPermissionByRole(rolePermission *[]stores.PermissionRuleDetail, roleId string, clientId string) *gorm.DB {
	return repository.DB.Preload("PermissionRule").
		Joins("JOIN permission_rules ON permission_rule_details.permission_rule_id = permission_rules.id").
		Find(&rolePermission, "permission_rules.ptype = ? and permission_rules.v0 = ? and permission_rules.v1 = ?", "p", roleId, clientId)
}
