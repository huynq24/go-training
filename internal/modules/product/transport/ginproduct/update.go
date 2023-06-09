package ginproduct

import (
	"github.com/gin-gonic/gin"
	"golang-training/internal/common"
	"golang-training/internal/components/app_context"
	"golang-training/internal/modules/product/biz"
	"golang-training/internal/modules/product/model"
	"golang-training/internal/modules/product/storage"
)

func UpdateProduct(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}

		var data productmodel.ProductUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := productstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productbiz.NewUpdateProductBiz(store)

		if err := biz.UpdateProductBiz(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(200, &data)
	}
}
