package tagstorage

import (
	"context"
	"golang-training/internal/modules/tag/model"
)

func (s *sqlStore) UpdateData(ctx context.Context, id int, data *tagmodel.Tag) error {
	db := s.db
	if err := db.WithContext(ctx).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
