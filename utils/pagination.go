package utils

import (
	"gaskn/config"
	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int         `json:"limit,omitempty;query:limit"`
	Page       int         `json:"page,omitempty;query:page"`
	Sort       string      `json:"sort,omitempty;query:sort"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

func (p *Pagination) GetOffset(c *fiber.Ctx) int {
	return (p.GetPage(c) - 1) * p.GetLimit(c)
}

func (p *Pagination) GetLimit(c *fiber.Ctx) int {
	if config.Config("PAGE_LIMIT") == "" {
		p.Limit = 10
	}

	if p.Limit == 0 {
		p.Limit, _ = strconv.Atoi(config.Config("PAGE_LIMIT"))
	}

	if c.Query("limit") != "" {
		if limit, err := strconv.Atoi(c.Query("limit")); err == nil {
			p.Limit = limit
		} else {
			p.Limit, _ = strconv.Atoi(config.Config("PAGE_LIMIT"))
		}
	}

	return p.Limit
}

func (p *Pagination) GetPage(c *fiber.Ctx) int {
	if p.Page == 0 {
		p.Page = 1
	}

	if c.Query("page") != "" {
		if page, err := strconv.Atoi(c.Query("page")); err == nil {
			p.Page = page
		} else {
			p.Page = 1
		}
	}
	return p.Page
}

func (p *Pagination) GetSort(c *fiber.Ctx) string {
	if p.Sort == "" {
		p.Sort = "id desc"
	}

	if c.Query("sort") != "" {
		p.Sort = c.Query("sort")
	}

	return p.Sort
}

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB, c *fiber.Ctx) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit(c))))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset(c)).Limit(pagination.GetLimit(c)).Order(pagination.GetSort(c))
	}
}
