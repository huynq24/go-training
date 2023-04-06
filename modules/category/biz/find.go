package categorybiz

import (
	"context"
	categorymodel "golang-training/modules/category/model"
)

type FindCategoryStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}) (*categorymodel.Category, error)
	FindAllData(ctx context.Context) (*[]categorymodel.Category, error)
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

func (biz *findCategoryBiz) FindAllCategories(ctx context.Context) (*[]categorymodel.Category, error) {
	result, err := biz.store.FindAllData(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}
