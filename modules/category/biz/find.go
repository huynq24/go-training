package categorybiz

import (
	"context"
	categorymodel "golang-training/modules/category/model"
)

type FindCategoryStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}) (*categorymodel.Category, error)
}

type findCategoryBiz struct {
	store FindCategoryStore
}

func NewFindCategoryBiz(store FindCategoryStore) *findCategoryBiz {
	return &findCategoryBiz{store: store}
}

func (biz *findCategoryBiz) FindCategory(ctx context.Context, id int) (*categorymodel.Category, error) {
	result, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return result, nil
}
