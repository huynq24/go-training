package categorystorage

import (
	"context"
	categorymodel "golang-training/modules/category/model"
)

func (s *sqlStore) Delete(context context.Context, id int) error {
	if err := s.db.WithContext(context).Table(categorymodel.Category{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return err
	}

	return nil
}
