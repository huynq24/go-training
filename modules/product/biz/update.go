package productbiz

import (
	"context"
	"errors"
	"golang-training/common"
	productmodel "golang-training/modules/product/model"
)

type UpdateProduct interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}) (*productmodel.Product, error)
	UpdateData(ctx context.Context, id int, data productmodel.Product) error
}

type updateProduct struct {
	Store UpdateProduct
}

func NewUpdateProductBiz(store UpdateProduct) *updateProduct {
	return &updateProduct{store}
}

func (biz *updateProduct) UpdateProductBiz(ctx context.Context, id int, data *productmodel.ProductUpdate) error {
	oldData, err := biz.Store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("Data deleted")
	}

	var updateData productmodel.Product
	updateData.Title = data.Title
	updateData.Image = data.Image
	updateData.Description = data.Description

	if data.CategoryId != "" {
		uid, err := common.FromBase58(data.CategoryId)
		if err != nil {
			return err
		}
		updateData.CategoryId = int(uid.GetLocalID())
	}

	if err := biz.Store.UpdateData(ctx, id, updateData); err != nil {
		return err
	}

	return nil
}
