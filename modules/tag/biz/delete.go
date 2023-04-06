package tagbiz

import (
	"context"
	"errors"
	tagmodel "golang-training/modules/tag/model"
)

type DeleteTagStore interface {
	Delete(context context.Context, id int) error
	FindDataWithCondition(context context.Context, condition map[string]interface{}) (*tagmodel.Tag, error)
}

type deleteTagBiz struct {
	store DeleteTagStore
}

func NewDeleteTagBiz(store DeleteTagStore) *deleteTagBiz {
	return &deleteTagBiz{store: store}
}

func (biz *deleteTagBiz) DeleteTag(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if oldData.Title == "" {
		return errors.New("Data deleted")
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
