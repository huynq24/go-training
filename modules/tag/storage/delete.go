package tagstorage

import (
	"context"
	tagmodel "golang-training/modules/tag/model"
)

func (s *sqlStore) Delete(context context.Context, id int) error {
	if err := s.db.Table(tagmodel.Tag{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return err
	}

	return nil
}
