package tagstorage

import (
	"context"
	common2 "golang-training/internal/common"
	tagmodel2 "golang-training/internal/modules/tag/model"
)

func (s *sqlStore) ListDataWithCondition(context context.Context, filter *tagmodel2.Filter, paging *common2.Paging, moreKeys ...string) ([]tagmodel2.Tag, error) {
	var result []tagmodel2.Tag

	db := s.db
	db = db.WithContext(context).Table(tagmodel2.Tag{}.TableName()).Where("status in (1)")

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
		if uid, err := common2.FromBase58(paging.FakeCursor); err == nil {
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
