package categorybiz

import (
	"context"
	"errors"
	"golang-training/internal/modules/category/model"
)

type UpdateCategory interface {
	FindCategory(ctx context.Context, condition map[string]interface{}) (*categorymodel.Category, error)
	UpdateData(ctx context.Context, id int, data *categorymodel.Category) error
}

type updateCategory struct {
	Store UpdateCategory
}

func NewUpdateCategoryBiz(store UpdateCategory) *updateCategory {
	return &updateCategory{store}
}

func (biz *updateCategory) UpdateCategoryBiz(ctx context.Context, id int, data *categorymodel.Category) error {
	oldData, err := biz.Store.FindCategory(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("Data deleted")
	}

	if err := biz.Store.UpdateData(ctx, id, data); err != nil {
		return err
	}

	return nil
}
