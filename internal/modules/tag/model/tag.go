package tagmodel

import (
	"golang-training/internal/common"
)

type Tag struct {
	common.SQLModel
	Title string `json:"title" binding:"required" gorm:"column:title;"`
}

func (Tag) TableName() string {
	return "tags"
}

func (t *Tag) Mask() {
	t.GenUID(common.DbTypeTag)
}
