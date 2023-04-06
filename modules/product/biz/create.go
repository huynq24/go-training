package productbiz

import (
	"context"
	productmodel "golang-training/modules/product/model"
	"gorm.io/gorm"
)

type CreateProductStore interface {
	Create(ctx context.Context, data *productmodel.Product) error
	FindProductExist(ctx context.Context, condition map[string]interface{}) (*productmodel.Product, error)
}

type createProductBiz struct {
	store CreateProductStore
}

func NewCreateProductBiz(store CreateProductStore) *createProductBiz {
	return &createProductBiz{store: store}
}

func (biz *createProductBiz) CreateProduct(ctx context.Context, data *productmodel.Product) error {
	_, err := biz.store.FindProductExist(ctx, map[string]interface{}{"title": data.Title, "category_id": data.CategoryId})
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
