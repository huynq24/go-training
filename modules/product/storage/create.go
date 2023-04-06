package productstorage

import (
	"context"
	productmodel "golang-training/modules/product/model"
)

func (s *sqlStore) Create(ctx context.Context, data *productmodel.Product) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
