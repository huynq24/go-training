package productbiz

import (
	"context"
	"golang-training/internal/modules/product/model"
)

type FindProductStore interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}) (*productmodel.Product, error)
}

type findProductBiz struct {
	store FindProductStore
}

func NewFindProductBiz(store FindProductStore) *findProductBiz {
	return &findProductBiz{store: store}
}

func (biz *findProductBiz) FindProduct(ctx context.Context, id int) (*productmodel.Product, error) {
	result, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return result, nil
}
