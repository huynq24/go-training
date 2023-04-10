package ginproduct

import (
	"github.com/gin-gonic/gin"
	"golang-training/internal/components/app_context"
	"golang-training/internal/modules/product/biz"
	"golang-training/internal/modules/product/model"
	"golang-training/internal/modules/product/storage"
	"net/http"
)

func CreateProduct(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data productmodel.Product

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := productstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := productbiz.NewCreateProductBiz(store)

		if err := biz.CreateProduct(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, &data)
	}
}
