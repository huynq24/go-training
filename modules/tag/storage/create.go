package tagstorage

import (
	"context"
	tagmodel "golang-training/modules/tag/model"
)

func (s *sqlStore) Create(ctx context.Context, data *tagmodel.Tag) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
