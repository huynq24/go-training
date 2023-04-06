package tagmodel

import "golang-training/common"

type Tag struct {
	common.SQLModel `json:",inline"`
	Title           string `json:"title" gorm:"column:title;unique;not null;"`
}

func (Tag) TableName() string {
	return "tags"
}
