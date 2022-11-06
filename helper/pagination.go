package helper

import (
	"github.com/krifik/test_fullstack_backend/entity"
	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int              `json:"limit,omitempty;query:limit"`
	Page       int              `json:"page,omitempty;query:page"`
	Sort       string           `json:"sort,omitempty;query:sort"`
	Order      string           `json:"order,omitempty;query:order"`
	TotalRows  int64            `json:"total"`
	TotalPages int              `json:"total_pages"`
	Search     string           `json:"search,omitempty;query:search"`
	Data       []entity.Student `json:"data"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}
func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 5
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}
func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "name"
	}
	return p.Sort
}
func (p *Pagination) GetOrder() string {
	if p.Order == "" {
		p.Order = "desc"
	}
	return p.Order
}
func (p *Pagination) SearchStudent(db *gorm.DB) *gorm.DB {
	return db.Where("name iLIKE ?", "%"+p.Search+"%")
}
