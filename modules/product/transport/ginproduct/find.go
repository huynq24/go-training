package ginproduct

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	"golang-training/common"
	productbiz "golang-training/modules/product/biz"
	productstorage "golang-training/modules/product/storage"
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
		context.JSON(http.StatusOK, result)
	}
}

func FindAllProducts(ctx app_context.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()

		store := productstorage.NewSQLStore(db)
		biz := productbiz.NewFindProductBiz(store)

		result, err := biz.FindAllProducts(context.Request.Context())
		for i := range result {
			result[i].Mask()
		}

		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, result)
	}
}
