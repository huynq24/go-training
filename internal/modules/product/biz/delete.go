package productbiz

import (
	"context"
	"errors"
	"golang-training/internal/modules/product/model"
)

type DeleteProductStore interface {
	Delete(context context.Context, id int) error
	FindProduct(context context.Context, condition map[string]interface{}) (*productmodel.Product, error)
}

type deleteProductBiz struct {
	store DeleteProductStore
}

func NewDeleteProductBiz(store DeleteProductStore) *deleteProductBiz {
	return &deleteProductBiz{store: store}
}

func (biz *deleteProductBiz) DeleteProduct(ctx context.Context, id int) error {
	oldData, err := biz.store.FindProduct(ctx, map[string]interface{}{"id": id})
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
