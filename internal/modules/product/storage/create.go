package productstorage

import (
	"context"
	"golang-training/internal/modules/product/model"
)

func (s *sqlStore) Create(ctx context.Context, data *productmodel.ProductCreate) error {
	return s.db.WithContext(ctx).Create(data).Error
}
