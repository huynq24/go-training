package productmodel

import (
	common2 "golang-training/internal/common"
)

type Product struct {
	common2.SQLModel
	Title       string `json:"title" binding:"required" gorm:"column:title;"`
	Image       string `json:"image" gorm:"column:image"`
	Description string `json:"description" gorm:"column:description"`
	CategoryId  int    `json:"categoryId" gorm:"column:category_id"`
}

type ProductUpdate struct {
	Title       string `json:"title" binding:"required" gorm:"column:title;"`
	Image       string `json:"image" gorm:"column:image"`
	Description string `json:"description" gorm:"column:description"`
	CategoryId  string `json:"categoryId" gorm:"column:category_id"`
}

func (ProductUpdate) TableName() string {
	return Product{}.TableName()
}

func (Product) TableName() string {
	return "products"
}

func (p *Product) Mask() {
	p.GenUID(common2.DbTypeProduct)
}
