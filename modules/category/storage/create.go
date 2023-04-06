package categorystorage

import (
	"context"
	"errors"
	categorymodel "golang-training/modules/category/model"
	"gorm.io/gorm"
)

func (s *sqlStore) Create(ctx context.Context, data *categorymodel.Category) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *sqlStore) FindCategoryExist(ctx context.Context, condition map[string]interface{}) (*categorymodel.Category, error) {
	var data categorymodel.Category

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("record not found")
		}
		return nil, err
	}

	return &data, nil
}
