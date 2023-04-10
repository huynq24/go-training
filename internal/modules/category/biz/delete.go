package categorybiz

import (
	"context"
	"errors"
	"golang-training/internal/modules/category/model"
)

type DeleteCategoryStore interface {
	Delete(context context.Context, id int) error
	FindDataWithCondition(context context.Context, condition map[string]interface{}) (*categorymodel.Category, error)
}

type deleteCategoryBiz struct {
	store DeleteCategoryStore
}

func NewDeleteCategoryBiz(store DeleteCategoryStore) *deleteCategoryBiz {
	return &deleteCategoryBiz{store: store}
}

func (biz *deleteCategoryBiz) DeleteCategory(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("Data deleted")
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
