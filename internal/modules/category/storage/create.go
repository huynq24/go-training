package categorystorage

import (
	"context"
	"golang-training/internal/modules/category/model"
)

func (s *sqlStore) Create(ctx context.Context, data *categorymodel.Category) error {
	return s.db.WithContext(ctx).Create(data).Error
}
