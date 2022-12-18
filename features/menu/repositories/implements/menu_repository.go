package implements

import (
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/menu/repositories"
	"gorm.io/gorm"
)

type MenuRepository struct {
	DB *gorm.DB
}

func NewMenuRepository(db *gorm.DB) repositories.MenuRepository {
	return &MenuRepository{
		DB: db,
	}
}

func (repository MenuRepository) CreateMenu(menu *stores.Menu) *gorm.DB {
	return repository.DB.Create(&menu)
}

func (repository MenuRepository) GetMenuById(menu *stores.Menu, id string) *gorm.DB {
	return repository.DB.Take(&menu, "id = ?", id)
}

func (repository MenuRepository) GetMenuAllByType(menu *[]stores.Menu, menuType stores.MenuType) *gorm.DB {
	if menuType == "" {
		return repository.DB.Find(&menu)
	}

	return repository.DB.Find(&menu, "menu_type = ?", menuType)
}
