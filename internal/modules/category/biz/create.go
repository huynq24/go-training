package categorybiz

import (
	"context"
	"golang-training/internal/modules/category/model"
	"gorm.io/gorm"
)

type CreateCategoryStore interface {
	Create(ctx context.Context, data *categorymodel.Category) error
	FindCategoryExist(ctx context.Context, condition map[string]interface{}) (*categorymodel.Category, error)
}

type createCategoryBiz struct {
	store CreateCategoryStore
}

func NewCreateCategoryBiz(store CreateCategoryStore) *createCategoryBiz {
	return &createCategoryBiz{store: store}
}

func (biz *createCategoryBiz) CreateCategory(ctx context.Context, data *categorymodel.Category) error {
	_, err := biz.store.FindCategoryExist(ctx, map[string]interface{}{"title": data.Title})
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() && data.Title != "" {
			if err := biz.store.Create(ctx, data); err != nil {
				return err
			}
		}
		return nil
	}
	return nil
}
