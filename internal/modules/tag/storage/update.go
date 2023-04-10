package tagstorage

import (
	"context"
	"golang-training/internal/modules/tag/model"
)

func (s *sqlStore) UpdateData(ctx context.Context, id int, data *tagmodel.Tag) error {
	return s.db.WithContext(ctx).Where("id = ?", id).Updates(data).Error
}
