package productbiz

import (
	"context"
	"golang-training/common"
	productmodel "golang-training/modules/product/model"
)

type ListProductStore interface {
	ListDataWithCondition(context context.Context, condition map[string]interface{}, paging *common.Paging, moreKeys ...string) ([]productmodel.Product, error)
}

type listProductBiz struct {
	store ListProductStore
}

func NewListProductBiz(store ListProductStore) *listProductBiz {
	return &listProductBiz{store: store}
}

func (biz *listProductBiz) ListProduct(ctx context.Context, paging *common.Paging) ([]productmodel.Product, error) {
	result, err := biz.store.ListDataWithCondition(ctx, nil, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}