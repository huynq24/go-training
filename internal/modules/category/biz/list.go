package categorybiz

import (
	"context"
	"golang-training/internal/common"
	categorymodel2 "golang-training/internal/modules/category/model"
)

type ListCategoryStore interface {
	ListDataWithCondition(context context.Context, filter *categorymodel2.Filter, paging *common.Paging, moreKeys ...string) ([]categorymodel2.Category, error)
}

type listCategoryBiz struct {
	store ListCategoryStore
}

func NewListCategoryBiz(store ListCategoryStore) *listCategoryBiz {
	return &listCategoryBiz{store: store}
}

func (biz *listCategoryBiz) ListCategory(ctx context.Context, filter *categorymodel2.Filter, paging *common.Paging) ([]categorymodel2.Category, error) {
	result, err := biz.store.ListDataWithCondition(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
