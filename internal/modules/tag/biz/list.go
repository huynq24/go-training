package tagbiz

import (
	"context"
	"golang-training/internal/common"
	tagmodel2 "golang-training/internal/modules/tag/model"
)

type ListTagStore interface {
	ListDataWithCondition(context context.Context, filter *tagmodel2.Filter, paging *common.Paging, moreKeys ...string) ([]tagmodel2.Tag, error)
}

type listTagBiz struct {
	store ListTagStore
}

func NewListTagBiz(store ListTagStore) *listTagBiz {
	return &listTagBiz{store: store}
}

func (biz *listTagBiz) ListTag(ctx context.Context, filter *tagmodel2.Filter, paging *common.Paging) ([]tagmodel2.Tag, error) {
	result, err := biz.store.ListDataWithCondition(ctx, filter, paging)

	if err != nil {
		return nil, err
	}

	return result, nil
}
