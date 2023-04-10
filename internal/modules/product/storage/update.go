package productstorage

import (
	"context"
	"golang-training/internal/modules/product/model"
)

func (s *sqlStore) UpdateData(ctx context.Context, id int, data productmodel.Product) error {
	return s.db.WithContext(ctx).Where("id = ?", id).Updates(&data).Error
}
