package categorystorage

import (
	"context"
	categorymodel "golang-training/modules/category/model"
)

func (s *sqlStore) UpdateData(ctx context.Context, id int, data *categorymodel.Category) error {
	db := s.db
	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
