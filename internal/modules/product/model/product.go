package productmodel

import (
	"golang-training/internal/common"
	categorymodel "golang-training/internal/modules/category/model"
	producttagmodel "golang-training/internal/modules/product_tags/model"
)

type Product struct {
	common.SQLModel
	Title       string                        `json:"title" binding:"required" gorm:"column:title;"`
	Image       string                        `json:"image" gorm:"column:image"`
	Description string                        `json:"description" gorm:"column:description"`
	CategoryId  int                           `json:"-" gorm:"column:category_id"`
	Category    categorymodel.Category        `json:"category" gorm:"foreignKey:Id"`
	ProductTags []*producttagmodel.ProductTag `json:"productTags" gorm:"foreignKey:ProductId"`
}

type ProductCreate struct {
	common.SQLModel
	Title       string `json:"title" binding:"required" gorm:"column:title;"`
	Image       string `json:"image" gorm:"column:image"`
	Description string `json:"description" gorm:"column:description"`
	CategoryId  int    `json:"categoryId" gorm:"column:category_id"`
}

type ProductUpdate struct {
	Title       *string `json:"title" binding:"required" gorm:"column:title;"`
	Image       *string `json:"image" gorm:"column:image"`
	Description *string `json:"description" gorm:"column:description"`
	CategoryId  *string `json:"categoryId" gorm:"column:category_id"`
}

func (ProductCreate) TableName() string {
	return Product{}.TableName()
}

func (ProductUpdate) TableName() string {
	return Product{}.TableName()
}

func (Product) TableName() string {
	return "products"
}

func (p *Product) Mask() {
	p.GenUID(common.DbTypeProduct)
}
