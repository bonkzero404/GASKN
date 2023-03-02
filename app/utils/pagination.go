package utils

import (
	"github.com/bonkzero404/gaskn/config"
	"math"
	"strconv"

	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int    `json:"limit,omitempty;query:limit"`
	Page       int    `json:"page,omitempty;query:page"`
	Sort       string `json:"sort,omitempty;query:sort"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
	Rows       any    `json:"rows"`
}

type PaginationRequest struct {
	Limit string
	Page  string
	Sort  string
}

func (p *Pagination) SetLimit(limit int) *Pagination {
	if config.Config("PAGE_LIMIT") == "" {
		p.Limit = 10
	}

	if p.Limit == 0 {
		p.Limit, _ = strconv.Atoi(config.Config("PAGE_LIMIT"))
	}

	if limit != 0 {
		p.Limit = limit
	} else {
		p.Limit, _ = strconv.Atoi(config.Config("PAGE_LIMIT"))
	}

	return p
}

func (p *Pagination) SetPage(page int) *Pagination {
	if p.Page == 0 {
		p.Page = 1
	}

	if page != 0 {
		p.Page = page
	} else {
		p.Page = 1
	}

	return p
}

func (p *Pagination) SetSort(sort string) *Pagination {
	if p.Sort == "" {
		p.Sort = "id desc"
	}

	if sort != "" {
		p.Sort = sort
	}

	return p
}

func (p *Pagination) SetOffset() int {
	return (p.Page - 1) * p.Limit
}

func (p *Pagination) Paginate(model any, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(model).Count(&totalRows)

	p.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(p.Limit)))
	p.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(p.SetOffset()).Limit(p.Limit).Order(p.Sort)
	}
}

func (pr *PaginationRequest) SetPagination() (int, int, string) {
	var page = 0
	var limit = 0
	var sort = "id desc"

	if pr.Page != "" {
		p, _ := strconv.Atoi(pr.Page)
		page = p
	}

	if pr.Limit != "" {
		l, _ := strconv.Atoi(pr.Limit)
		limit = l
	}

	if pr.Sort != "" {
		sort = pr.Sort
	}

	return page, limit, sort
}
