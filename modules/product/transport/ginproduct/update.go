package ginproduct

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	"golang-training/common"
	productbiz "golang-training/modules/product/biz"
	productmodel "golang-training/modules/product/model"
	productstorage "golang-training/modules/product/storage"
)

func UpdateProduct(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}

		var data productmodel.Product

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
