package driver

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var Enforcer *casbin.Enforcer

func InitCasbin() {
	adapter, _ := gormadapter.NewAdapterByDB(DB)
	Enforcer, _ = casbin.NewEnforcer("casbin_rbac_model.conf", adapter)
	Enforcer.EnableAutoSave(true)
}
