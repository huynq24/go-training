package productstorage

import (
	"context"
	"errors"
	productmodel "golang-training/modules/product/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(context context.Context, condition map[string]interface{}) (*productmodel.Product, error) {
	var data productmodel.Product

	if err := s.db.WithContext(context).Where(condition).First(&data).Error; err != nil {
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
