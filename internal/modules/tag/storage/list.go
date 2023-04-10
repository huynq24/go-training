package tagstorage

import (
	"context"
	"golang-training/internal/common"
	tagmodel "golang-training/internal/modules/tag/model"
)

func (s *sqlStore) ListTag(context context.Context, filter *tagmodel.Filter, paging *common.Paging, moreKeys ...string) ([]tagmodel.Tag, error) {
	var result []tagmodel.Tag

	db := s.db
	db = db.WithContext(context).Table(tagmodel.Tag{}.TableName()).Where("status in (1)")

	if v := filter; v != nil {
		if v.TagTitle != "" {
			db = db.Where("title = ?", v.TagTitle)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if paging.FakeCursor != "" {
		if uid, err := common.FromBase58(paging.FakeCursor); err == nil {
			db = db.Where("id < ?", uid.GetLocalID())
		} else {
			db = db.Offset((paging.Page - 1) * paging.Limit)
		}
	}

	if err := db.Limit(paging.Limit).Order("id desc").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
