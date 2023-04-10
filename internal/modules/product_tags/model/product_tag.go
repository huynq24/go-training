package producttagmodel

import (
	"golang-training/internal/common"
	tagmodel "golang-training/internal/modules/tag/model"
)

type ProductTag struct {
	common.SQLModel
	ProductId int           `json:"-" gorm:"column:product_id"`
	TagId     int           `json:"-" gorm:"column:tag_id"`
	Tag       *tagmodel.Tag `json:"tag" gorm:"TagId"`
}

func (ProductTag) TableName() string {
	return "product_tags"
}
