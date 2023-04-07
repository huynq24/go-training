package productstorage

import (
	"context"
	"golang-training/common"
	productmodel "golang-training/modules/product/model"
)

func (s *sqlStore) ListDataWithCondition(context context.Context, condition map[string]interface{}, paging *common.Paging, moreKeys ...string) ([]productmodel.Product, error) {
	var result []productmodel.Product

	db := s.db
	db = db.WithContext(context).Table(productmodel.Product{}.TableName()).Where(condition).Where("status in (1)")

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
