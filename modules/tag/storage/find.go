package tagstorage

import (
	"context"
	"errors"
	tagmodel "golang-training/modules/tag/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(context context.Context, condition map[string]interface{}) (*tagmodel.Tag, error) {
	var data tagmodel.Tag

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("record not found")
		}
		return nil, err
	}

	return &data, nil
}

func (s *sqlStore) FindAllData(context context.Context) ([]tagmodel.Tag, error) {
	var data []tagmodel.Tag

	if err := s.db.Find(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("record not found")
		}
		return nil, err
	}

	return data, nil
}
