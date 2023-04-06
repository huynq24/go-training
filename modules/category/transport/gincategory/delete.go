package gincategory

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	categorybiz "golang-training/modules/category/biz"
	categorystorage "golang-training/modules/category/storage"
	"strconv"
)

func DeleteCategory(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			panic(err)
			return
		}

		store := categorystorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := categorybiz.NewDeleteCategoryBiz(store)

		if err := biz.DeleteCategory(c.Request.Context(), uid); err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			panic(err)
			return
		}

		c.JSON(200, "Delete success")
	}
}
