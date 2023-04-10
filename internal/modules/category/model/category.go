package categorymodel

import (
	"golang-training/internal/common"
)

type Category struct {
	common.SQLModel
	Title string `json:"title" binding:"required" gorm:"column:title;"`
}

func (Category) TableName() string {
	return "categories"
}

func (c *Category) Mask() {
	c.GenUID(common.DbTypeCategory)
}
