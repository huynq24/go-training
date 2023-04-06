package tagbiz

import (
	"context"
	tagmodel "golang-training/modules/tag/model"
)

type CreatTagStore interface {
	Create(context context.Context, data *tagmodel.Tag) error
}

type createTagBiz struct {
	store CreatTagStore
}

func NewCreateTagBiz(store CreatTagStore) *createTagBiz {
	return &createTagBiz{store: store}
}

func (biz *createTagBiz) CreateTag(ctx context.Context, data *tagmodel.Tag) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}
	return nil
}
