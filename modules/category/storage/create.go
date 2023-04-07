package categorystorage

import (
	"context"
	categorymodel "golang-training/modules/category/model"
)

func (s *sqlStore) Create(ctx context.Context, data *categorymodel.Category) error {
	if err := s.db.WithContext(ctx).Create(data).Error; err != nil {
		return err
	}

	return nil
}
