package tagstorage

import (
	"context"
	"errors"
	tagmodel "golang-training/modules/tag/model"
	"gorm.io/gorm"
)

func (s *sqlStore) Create(ctx context.Context, data *tagmodel.Tag) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}

func (s *sqlStore) FindTagExist(ctx context.Context, condition map[string]interface{}) (*tagmodel.Tag, error) {
	var data tagmodel.Tag

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("record not found")
		}
		return nil, err
	}

	return &data, nil
}
