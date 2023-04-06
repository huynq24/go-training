package productstorage

import (
	"context"
	"errors"
	productmodel "golang-training/modules/product/model"
	"gorm.io/gorm"
)

func (s *sqlStore) Create(ctx context.Context, data *productmodel.Product) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *sqlStore) FindProductExist(ctx context.Context, condition map[string]interface{}) (*productmodel.Product, error) {
	var data productmodel.Product

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("record not found")
		}
		return nil, err
	}

	return &data, nil
}
