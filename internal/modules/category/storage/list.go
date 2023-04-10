package categorystorage

import (
	"context"
	common2 "golang-training/internal/common"
	categorymodel2 "golang-training/internal/modules/category/model"
)

func (s *sqlStore) ListDataWithCondition(context context.Context, filter *categorymodel2.Filter, paging *common2.Paging, moreKeys ...string) ([]categorymodel2.Category, error) {
	var result []categorymodel2.Category

	db := s.db
	db = db.WithContext(context).Table(categorymodel2.Category{}.TableName()).Where("status in (1)")

	if v := filter; v != nil {
		if v.CategoryTitle != "" {
			db = db.WithContext(context).Where("title = ?", v.CategoryTitle)
		}
	}

	if err := db.WithContext(context).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	for i := range moreKeys {
		db = db.WithContext(context).Preload(moreKeys[i])
	}

	if paging.FakeCursor != "" {
		if uid, err := common2.FromBase58(paging.FakeCursor); err == nil {
			db = db.WithContext(context).Where("id < ?", uid.GetLocalID())
		} else {
			db = db.WithContext(context).Offset((paging.Page - 1) * paging.Limit)
		}
	}

	if err := db.WithContext(context).Limit(paging.Limit).Order("id desc").Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
