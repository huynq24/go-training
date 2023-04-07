package tagbiz

import (
	"context"
	"golang-training/common"
	tagmodel "golang-training/modules/tag/model"
)

type ListTagStore interface {
	ListDataWithCondition(context context.Context, filter *tagmodel.Filter, paging *common.Paging, moreKeys ...string) ([]tagmodel.Tag, error)
}

type listTagBiz struct {
	store ListTagStore
}

func NewListTagBiz(store ListTagStore) *listTagBiz {
	return &listTagBiz{store: store}
}

func (biz *listTagBiz) ListTag(ctx context.Context, filter *tagmodel.Filter, paging *common.Paging) ([]tagmodel.Tag, error) {
	result, err := biz.store.ListDataWithCondition(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
