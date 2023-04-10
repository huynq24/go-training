package productbiz

import (
	"context"
	"golang-training/internal/common"
	productmodel2 "golang-training/internal/modules/product/model"
)

type ListProductStore interface {
	ListDataWithCondition(context context.Context, filter *productmodel2.Filter, paging *common.Paging, moreKeys ...string) ([]productmodel2.Product, error)
}

type listProductBiz struct {
	store ListProductStore
}

func NewListProductBiz(store ListProductStore) *listProductBiz {
	return &listProductBiz{store: store}
}

func (biz *listProductBiz) ListProduct(ctx context.Context, filter *productmodel2.Filter, paging *common.Paging) ([]productmodel2.Product, error) {
	result, err := biz.store.ListDataWithCondition(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
