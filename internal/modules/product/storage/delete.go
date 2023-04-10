package productstorage

import (
	"context"
	"golang-training/internal/modules/product/model"
)

func (s *sqlStore) Delete(context context.Context, id int) error {
	return s.db.WithContext(context).Table(productmodel.Product{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error
}
