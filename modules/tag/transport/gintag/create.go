package gintag

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	tagbiz "golang-training/modules/tag/biz"
	tagmodel "golang-training/modules/tag/model"
	tagstorage "golang-training/modules/tag/storage"
	"net/http"
)

func CreateTag(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data tagmodel.Tag

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
			return
		}

		store := tagstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := tagbiz.NewCreateTagBiz(store)

		if err := biz.CreateTag(c.Request.Context(), &data); err != nil {
			panic(err)
			return
		}

		c.JSON(http.StatusOK, &data)
	}
}
