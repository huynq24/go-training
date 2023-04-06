package ginproduct

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	productbiz "golang-training/modules/product/biz"
	productstorage "golang-training/modules/product/storage"
	"strconv"
)

func DeleteProduct(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			panic(err)
			return
		}

		store := productstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productbiz.NewDeleteProductBiz(store)

		if err := biz.DeleteProduct(c.Request.Context(), uid); err != nil {
			c.JSON(500, gin.H{
				"error": err,
			})
			panic(err)
			return
		}

		c.JSON(200, "Delete success")
	}
}
