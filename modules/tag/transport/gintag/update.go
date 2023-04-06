package gintag

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	tagbiz "golang-training/modules/tag/biz"
	tagmodel "golang-training/modules/tag/model"
	tagstorage "golang-training/modules/tag/storage"
	"strconv"
)

func UpdateTag(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		var data tagmodel.Tag

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		store := tagstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := tagbiz.NewUpdateTagBiz(store)

		if err := biz.UpdateTagBiz(c.Request.Context(), uid, &data); err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(200, &data)
	}
}
