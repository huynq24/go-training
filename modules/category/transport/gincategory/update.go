package gincategory

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	categorybiz "golang-training/modules/category/biz"
	categorymodel "golang-training/modules/category/model"
	categorystorage "golang-training/modules/category/storage"
	"strconv"
)

func UpdateCategory(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		var data categorymodel.Category

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		store := categorystorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := categorybiz.NewUpdateCategoryBiz(store)

		if err := biz.UpdateCategoryBiz(c.Request.Context(), uid, &data); err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(200, &data)
	}
}
