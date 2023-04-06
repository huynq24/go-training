package categorybiz

import (
	"context"
	categorymodel "golang-training/modules/category/model"
)

type CreateCategoryStore interface {
	Create(ctx context.Context, data *categorymodel.Category) error
}

type createCategoryBiz struct {
	store CreateCategoryStore
}

func NewCreateCategoryBiz(store CreateCategoryStore) *createCategoryBiz {
	return &createCategoryBiz{store: store}
}

func (biz *createCategoryBiz) CreateCategory(ctx context.Context, data *categorymodel.Category) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}
	return nil
}
