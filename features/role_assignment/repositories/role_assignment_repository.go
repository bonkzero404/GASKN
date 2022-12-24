package repositories

import (
	"github.com/bonkzero404/gaskn/database/stores"
	"gorm.io/gorm"
)

type RoleAssignmentRepository interface {
	GetPermissionByRole(role *[]stores.PermissionRuleDetail, roleId string, clientId string) *gorm.DB
}
