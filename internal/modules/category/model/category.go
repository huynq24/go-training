package categorymodel

import (
	common2 "golang-training/internal/common"
)

type Category struct {
	common2.SQLModel
	Title string `json:"title" binding:"required" gorm:"column:title;"`
}

func (Category) TableName() string {
	return "categories"
}

func (c *Category) Mask() {
	c.GenUID(common2.DbTypeCategory)
}
