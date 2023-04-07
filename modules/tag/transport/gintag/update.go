package gintag

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	"golang-training/common"
	tagbiz "golang-training/modules/tag/biz"
	tagmodel "golang-training/modules/tag/model"
	tagstorage "golang-training/modules/tag/storage"
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
