package productstorage

import (
	"context"
	"golang-training/common"
	productmodel "golang-training/modules/product/model"
)

func (s *sqlStore) ListDataWithCondition(context context.Context, condition map[string]interface{}, paging *common.Paging, moreKeys ...string) ([]productmodel.Product, error) {
	var result []productmodel.Product

	db := s.db
	db = db.Table(productmodel.Product{}.TableName()).Where(condition).Where("status in (1)")

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
