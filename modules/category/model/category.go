package categorymodel

import "golang-training/common"

type Category struct {
	common.SQLModel `json:",inline"`
	Title           string `json:"title" gorm:"column:title;unique;not null;"`
}

func (Category) TableName() string {
	return "categories"
}
