package productbiz

import (
	"context"
	productmodel "golang-training/modules/product/model"
)

type CreateProductStore interface {
	Create(ctx context.Context, data *productmodel.Product) error
}

type createProductBiz struct {
	store CreateProductStore
}

func NewCreateProductBiz(store CreateProductStore) *createProductBiz {
	return &createProductBiz{store: store}
}

func (biz *createProductBiz) CreateProduct(ctx context.Context, data *productmodel.Product) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}
	return nil
}
