package tagbiz

import (
	"context"
	"golang-training/internal/modules/tag/model"
	"gorm.io/gorm"
)

type CreatTagStore interface {
	Create(context context.Context, data *tagmodel.Tag) error
	FindTagExist(ctx context.Context, condition map[string]interface{}) (*tagmodel.Tag, error)
}

type createTagBiz struct {
	store CreatTagStore
}

func NewCreateTagBiz(store CreatTagStore) *createTagBiz {
	return &createTagBiz{store: store}
}

func (biz *createTagBiz) CreateTag(ctx context.Context, data *tagmodel.Tag) error {
	_, err := biz.store.FindTagExist(ctx, map[string]interface{}{"title": data.Title})
	if err != nil && err.Error() == gorm.ErrRecordNotFound.Error() {
		return biz.store.Create(ctx, data)
	}
	return nil
}
