package categorystorage

import (
	"context"
	"golang-training/internal/modules/category/model"
)

func (s *sqlStore) UpdateData(ctx context.Context, id int, data *categorymodel.Category) error {
	return s.db.WithContext(ctx).Where("id = ?", id).Updates(data).Error
}
