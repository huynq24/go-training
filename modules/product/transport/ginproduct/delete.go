package ginproduct

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	"golang-training/common"
	productbiz "golang-training/modules/product/biz"
	productstorage "golang-training/modules/product/storage"
)

func DeleteProduct(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}

		store := productstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productbiz.NewDeleteProductBiz(store)

		if err := biz.DeleteProduct(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(200, "Delete success")
	}
}
