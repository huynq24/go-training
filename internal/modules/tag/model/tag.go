package tagmodel

import (
	common2 "golang-training/internal/common"
)

type Tag struct {
	common2.SQLModel
	Title string `json:"title" binding:"required" gorm:"column:title;"`
}

func (Tag) TableName() string {
	return "tags"
}

func (t *Tag) Mask() {
	t.GenUID(common2.DbTypeTag)
}
