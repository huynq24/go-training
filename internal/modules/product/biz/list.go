package productbiz

import (
	"context"
	"golang-training/internal/common"
	productmodel "golang-training/internal/modules/product/model"
)

type ListProductStore interface {
	ListProduct(context context.Context, filter *productmodel.Filter, paging *common.Paging, moreKeys ...string) ([]*productmodel.Product, error)
}

type listProductBiz struct {
	store ListProductStore
}

func NewListProductBiz(store ListProductStore) *listProductBiz {
	return &listProductBiz{store: store}
}

func (biz *listProductBiz) ListProduct(ctx context.Context, filter *productmodel.Filter, paging *common.Paging) ([]*productmodel.Product, error) {
	result, err := biz.store.ListProduct(ctx, filter, paging, "Category", "ProductTags.Tag")

	if err != nil {
		return nil, err
	}

	return result, nil
}
