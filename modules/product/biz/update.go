package productbiz

import (
	"context"
	"errors"
	productmodel "golang-training/modules/product/model"
)

type UpdateProduct interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}) (*productmodel.Product, error)
	UpdateData(ctx context.Context, id int, data *productmodel.Product) error
}

type updateProduct struct {
	Store UpdateProduct
}

func NewUpdateProductBiz(store UpdateProduct) *updateProduct {
	return &updateProduct{store}
}

func (biz *updateProduct) UpdateProductBiz(ctx context.Context, id int, data *productmodel.Product) error {
	oldData, err := biz.Store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
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
