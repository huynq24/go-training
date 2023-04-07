package tagbiz

import (
	"context"
	tagmodel "golang-training/modules/tag/model"
)

type FindTagStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}) (*tagmodel.Tag, error)
	FindAllData(ctx context.Context) ([]tagmodel.Tag, error)
}

type findTagBiz struct {
	store FindTagStore
}

func NewFindTagBiz(store FindTagStore) *findTagBiz {
	return &findTagBiz{store: store}
}

func (biz *findTagBiz) FindTag(ctx context.Context, id int) (*tagmodel.Tag, error) {
	result, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (biz *findTagBiz) FindAllTags(ctx context.Context) ([]tagmodel.Tag, error) {
	result, err := biz.store.FindAllData(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}
