package productstorage

import (
	"context"
	"golang-training/internal/modules/product/model"
)

func (s *sqlStore) Create(ctx context.Context, data *productmodel.ProductCreate) error {
	if err := s.db.WithContext(ctx).Create(data).Error; err != nil {
		return err
	}

	return nil
}
