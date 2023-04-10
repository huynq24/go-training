package productstorage

import (
	"context"
	common2 "golang-training/internal/common"
	productmodel2 "golang-training/internal/modules/product/model"
)

func (s *sqlStore) ListDataWithCondition(context context.Context, filter *productmodel2.Filter, paging *common2.Paging, moreKeys ...string) ([]productmodel2.Product, error) {
	var result []productmodel2.Product

	db := s.db
	db = db.WithContext(context).Table(productmodel2.Product{}.TableName()).Where("status in (1)")

	if v := filter; v != nil {
		if v.ProductTitle != "" {
			db = db.Where("title = ?", v.ProductTitle)
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
