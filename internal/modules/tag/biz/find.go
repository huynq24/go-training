package tagbiz

import (
	"context"
	"golang-training/internal/modules/tag/model"
)

type FindTagStore interface {
	FindTag(ctx context.Context, condition map[string]interface{}) (*tagmodel.Tag, error)
}

type findTagBiz struct {
	store FindTagStore
}

func NewFindTagBiz(store FindTagStore) *findTagBiz {
	return &findTagBiz{store: store}
}

func (biz *findTagBiz) FindTag(ctx context.Context, id int) (*tagmodel.Tag, error) {
	result, err := biz.store.FindTag(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return result, nil
}
