package categorystorage

import (
	"context"
	"golang-training/common"
	categorymodel "golang-training/modules/category/model"
)

func (s *sqlStore) ListDataWithCondition(context context.Context, condition map[string]interface{}, filter *categorymodel.Filter, paging *common.Paging, moreKeys ...string) ([]categorymodel.Category, error) {
	var result []categorymodel.Category

	db := s.db
	db = db.WithContext(context).Table(categorymodel.Category{}.TableName()).Where(condition).Where("status in (1)")

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
		if uid, err := common.FromBase58(paging.FakeCursor); err == nil {
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
