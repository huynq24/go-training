package productmodel

import (
	"golang-training/common"
)

type Product struct {
	common.SQLModel `json:",inline"`
	Title           string `json:"title" gorm:"column:title;unique;not null;"`
	Image           string `json:"image" gorm:"column:image"`
	Description     string `json:"description" gorm:"column:description"`
	CategoryId      int    `json:"categoryId" gorm:"column:category_id"`
}

func (Product) TableName() string {
	return "products"
}

func (p *Product) Mask() {
	p.GenUID(common.DbTypeProduct)
}
