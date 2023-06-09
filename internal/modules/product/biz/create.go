package productbiz

import (
	"context"
	"golang-training/internal/modules/product/model"
	"gorm.io/gorm"
)

type CreateProductStore interface {
	Create(ctx context.Context, data *productmodel.ProductCreate) error
	FindProductExist(ctx context.Context, condition map[string]interface{}) (*productmodel.Product, error)
}

type createProductBiz struct {
	store CreateProductStore
}

func NewCreateProductBiz(store CreateProductStore) *createProductBiz {
	return &createProductBiz{store: store}
}

func (biz *createProductBiz) CreateProduct(ctx context.Context, data *productmodel.ProductCreate) error {
	_, err := biz.store.FindProductExist(ctx, map[string]interface{}{"title": data.Title, "category_id": data.CategoryId})
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		return err
	}
	return biz.store.Create(ctx, data)
}
