package gintag

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	tagbiz "golang-training/modules/tag/biz"
	tagstorage "golang-training/modules/tag/storage"
	"strconv"
)

func DeleteTag(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		store := tagstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := tagbiz.NewDeleteTagBiz(store)

		if err := biz.DeleteTag(c.Request.Context(), uid); err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(200, "Delete success")
	}
}
