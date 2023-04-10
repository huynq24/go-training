package gintag

import (
	"github.com/gin-gonic/gin"
	"golang-training/internal/common"
	"golang-training/internal/components/app_context"
	"golang-training/internal/modules/tag/biz"
	"golang-training/internal/modules/tag/model"
	"golang-training/internal/modules/tag/storage"
)

func UpdateTag(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}

		var data tagmodel.Tag
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := tagstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := tagbiz.NewUpdateTagBiz(store)

		if err := biz.UpdateTagBiz(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(200, &data)
	}
}
