package productstorage

import (
	"context"
	"errors"
	"golang-training/internal/modules/product/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindProduct(context context.Context, condition map[string]interface{}) (*productmodel.Product, error) {
	var data productmodel.Product

	if err := s.db.WithContext(context).Preload("Category").Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("record not found")
		}
		return nil, err
	}

	return &data, nil
}

func (s *sqlStore) FindProductExist(ctx context.Context, condition map[string]interface{}) (*productmodel.Product, error) {
	var data productmodel.Product

	if err := s.db.WithContext(ctx).Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("record not found")
		}
		return nil, err
	}

	return &data, nil
}
