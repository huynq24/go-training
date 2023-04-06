package ginproduct

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	productbiz "golang-training/modules/product/biz"
	productmodel "golang-training/modules/product/model"
	productstorage "golang-training/modules/product/storage"
	"strconv"
)

func UpdateProduct(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		var data productmodel.Product

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		store := productstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productbiz.NewUpdateProductBiz(store)

		if err := biz.UpdateProductBiz(c.Request.Context(), uid, &data); err != nil {
			c.JSON(401, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(200, &data)
	}
}
