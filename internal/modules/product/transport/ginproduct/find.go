package ginproduct

import (
	"github.com/gin-gonic/gin"
	"golang-training/internal/common"
	"golang-training/internal/components/app_context"
	"golang-training/internal/modules/product/biz"
	"golang-training/internal/modules/product/storage"
	"net/http"
)

func FindProduct(ctx app_context.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		id, err := common.FromBase58(context.Param("id"))

		if err != nil {
			panic(err)
		}

		store := productstorage.NewSQLStore(db)
		biz := productbiz.NewFindProductBiz(store)

		result, err := biz.FindProduct(context.Request.Context(), int(id.GetLocalID()))
		if err != nil {
			panic(err)
		}
		result.Mask()
		result.Category.Mask()
		context.JSON(http.StatusOK, result)
	}
}
