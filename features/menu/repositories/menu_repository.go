package repositories

import (
	"github.com/bonkzero404/gaskn/database/stores"
	"gorm.io/gorm"
)

type MenuRepository interface {
	CreateMenu(menu *stores.Menu) *gorm.DB

	GetMenuById(menu *stores.Menu, id string) *gorm.DB

	GetMenuAllByType(menu *[]stores.Menu, lang string, menuType stores.MenuType, sort string) *gorm.DB
}
