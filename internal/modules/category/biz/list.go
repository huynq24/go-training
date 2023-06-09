package categorybiz

import (
	"context"
	"golang-training/internal/common"
	categorymodel "golang-training/internal/modules/category/model"
)

type ListCategoryStore interface {
	ListDataWithCondition(context context.Context, filter *categorymodel.Filter, paging *common.Paging, moreKeys ...string) ([]categorymodel.Category, error)
}

type listCategoryBiz struct {
	store ListCategoryStore
}

func NewListCategoryBiz(store ListCategoryStore) *listCategoryBiz {
	return &listCategoryBiz{store: store}
}

func (biz *listCategoryBiz) ListCategory(ctx context.Context, filter *categorymodel.Filter, paging *common.Paging) ([]categorymodel.Category, error) {
	result, err := biz.store.ListDataWithCondition(ctx, filter, paging)
	if err != nil {
		return nil, err
	}

	return result, nil
}
