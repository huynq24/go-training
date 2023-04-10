package categorystorage

import (
	"context"
	"golang-training/internal/modules/category/model"
)

func (s *sqlStore) Delete(context context.Context, id int) error {
	return s.db.WithContext(context).Table(categorymodel.Category{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error
}
