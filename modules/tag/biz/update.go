package tagbiz

import (
	"context"
	"errors"
	tagmodel "golang-training/modules/tag/model"
)

type UpdateTag interface {
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}) (*tagmodel.Tag, error)
	UpdateData(ctx context.Context, id int, data *tagmodel.Tag) error
}

type updateTag struct {
	Store UpdateTag
}

func NewUpdateTagBiz(store UpdateTag) *updateTag {
	return &updateTag{store}
}

func (biz *updateTag) UpdateTagBiz(ctx context.Context, id int, data *tagmodel.Tag) error {
	oldData, err := biz.Store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if oldData.Title == "" {
		return errors.New("Data deleted")
	}

	if err := biz.Store.UpdateData(ctx, id, data); err != nil {
		return err
	}

	return nil
}
