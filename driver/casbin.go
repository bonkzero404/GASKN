package driver

import (
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var Enforcer *casbin.Enforcer

func InitCasbin() {
	adapter, _ := gormadapter.NewAdapterByDBUseTableName(DB, "permission", "rules")
	Enforcer, _ = casbin.NewEnforcer("casbin_rbac_model.conf", adapter)
	Enforcer.EnableAutoSave(true)
}

func checkPolicy(casbinRule *stores.PermissionRule) (*stores.PermissionRule, error) {
	db := DB
	err := db.Take(
		&casbinRule,
		"ptype = ? and v0 = ? and v1 = ? and v2 = ? and v3 = ?",
		casbinRule.Ptype,
		casbinRule.V0,
		casbinRule.V1,
		casbinRule.V2,
		casbinRule.V3,
	).Error

	if err != nil {
		return &stores.PermissionRule{}, nil
	}

	return casbinRule, nil
}

func addDetailRule(casbinRuleId uint, userName string, roleName string, clientName string) (bool, error) {
	db := DB
	casbinRuleDetail := stores.PermissionRuleDetail{
		PermissionRuleId: casbinRuleId,
		UserName:         userName,
		RoleName:         roleName,
		ClientName:       clientName,
	}

	errCreateDetail := db.Create(&casbinRuleDetail).Error

	if errCreateDetail != nil {
		return false, errCreateDetail
	}

	return true, nil
}

func AddGroupingPolicy(userId string, roleId string, clientId string, userName string, roleName string, clientName string) (bool, error) {
	if g, err := Enforcer.AddGroupingPolicy(userId, roleId, clientId); !g {
		return false, err
	}

	casbinRule := stores.PermissionRule{
		Ptype: "g",
		V0:    userId,
		V1:    roleId,
		V2:    clientId,
	}

	dataRule, err := checkPolicy(&casbinRule)

	if err == nil {
		_, errCreateDetail := addDetailRule(dataRule.ID, userName, roleName, clientName)

		if errCreateDetail != nil {
			return false, errCreateDetail
		}
	}

	return true, nil
}

func AddPolicy(roleId string, clientId string, routeEndpoint string, httpMethod string, userName string, roleName string, clientName string) (bool, error) {
	if g, err := Enforcer.AddPolicy(roleId, clientId, routeEndpoint, httpMethod); !g {
		return false, err
	}

	casbinRule := stores.PermissionRule{
		Ptype: "p",
		V0:    roleId,
		V1:    clientId,
		V2:    routeEndpoint,
		V3:    httpMethod,
	}

	dataRule, err := checkPolicy(&casbinRule)

	if err == nil {
		_, errCreateDetail := addDetailRule(dataRule.ID, userName, roleName, clientName)

		if errCreateDetail != nil {
			return false, errCreateDetail
		}
	}

	return true, nil
}

func RemovePolicy(roleId string, clientId string, routeEndpoint string, httpMethod string) (bool, error) {
	if p, err := Enforcer.RemovePolicy(roleId, clientId, routeEndpoint, httpMethod); !p {
		return false, err
	}

	return true, nil
}

//goland:noinspection GoUnusedExportedFunction
func RemoveGroupingPolicy(userId string, roleId string, clientId string) (bool, error) {
	if g, err := Enforcer.RemoveGroupingPolicy(userId, roleId, clientId); !g {
		return false, err
	}

	return true, nil
}
