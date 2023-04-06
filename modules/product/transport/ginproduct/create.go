package ginproduct

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	productbiz "golang-training/modules/product/biz"
	productmodel "golang-training/modules/product/model"
	productstorage "golang-training/modules/product/storage"
	"net/http"
)

func CreateProduct(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data productmodel.Product

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
			return
		}

		store := productstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productbiz.NewCreateProductBiz(store)

		if err := biz.CreateProduct(c.Request.Context(), &data); err != nil {
			panic(err)
			return
		}

		c.JSON(http.StatusOK, &data)
	}
}
