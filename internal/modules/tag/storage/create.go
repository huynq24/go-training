package tagstorage

import (
	"context"
	"golang-training/internal/modules/tag/model"
)

func (s *sqlStore) Create(ctx context.Context, data *tagmodel.Tag) error {
	if err := s.db.WithContext(ctx).Create(data).Error; err != nil {
		return err
	}

	return nil
}
