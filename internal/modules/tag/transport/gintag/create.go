package gintag

import (
	"github.com/gin-gonic/gin"
	"golang-training/internal/components/app_context"
	"golang-training/internal/modules/tag/biz"
	"golang-training/internal/modules/tag/model"
	"golang-training/internal/modules/tag/storage"
	"net/http"
)

func CreateTag(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data tagmodel.Tag

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := tagstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := tagbiz.NewCreateTagBiz(store)

		if err := biz.CreateTag(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, &data)
	}
}
