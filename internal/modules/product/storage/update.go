package productstorage

import (
	"context"
	"golang-training/internal/modules/product/model"
)

func (s *sqlStore) UpdateData(ctx context.Context, id int, data productmodel.Product) error {
	db := s.db
	if err := db.WithContext(ctx).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
