package categorystorage

import (
	"context"
	"golang-training/internal/modules/category/model"
)

func (s *sqlStore) UpdateData(ctx context.Context, id int, data *categorymodel.Category) error {
	db := s.db
	if err := db.WithContext(ctx).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
