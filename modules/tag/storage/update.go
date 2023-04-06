package tagstorage

import (
	"context"
	tagmodel "golang-training/modules/tag/model"
)

func (s *sqlStore) UpdateData(ctx context.Context, id int, data *tagmodel.Tag) error {
	db := s.db
	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
