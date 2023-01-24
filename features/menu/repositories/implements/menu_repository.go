package implements

import (
	"github.com/bonkzero404/gaskn/config"
	"github.com/bonkzero404/gaskn/database/stores"
	"github.com/bonkzero404/gaskn/features/menu/repositories"
	"gorm.io/datatypes"
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

func (repository MenuRepository) GetMenuAllByType(menu *[]stores.Menu, lang string, menuType stores.MenuType, sort string) *gorm.DB {
	var lng = config.Config("LANG")

	if lang != "" {
		lng = lang
	}
	if menuType == "" {
		return repository.DB.Order("sort "+sort).Find(
			&menu,
			datatypes.JSONQuery("menu_name").HasKey(lng),
			datatypes.JSONQuery("menu_description").HasKey(lng))
	}

	return repository.DB.Order("sort "+sort).Find(
		&menu,
		"menu_type = ?", menuType,
		datatypes.JSONQuery("menu_name").HasKey(lng),
		datatypes.JSONQuery("menu_description").HasKey(lng))
}
