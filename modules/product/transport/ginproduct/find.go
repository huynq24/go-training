package ginproduct

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	productbiz "golang-training/modules/product/biz"
	productstorage "golang-training/modules/product/storage"
	"net/http"
	"strconv"
)

func FindProduct(ctx app_context.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			panic(err)
			return
		}

		store := productstorage.NewSQLStore(db)
		biz := productbiz.NewFindProductBiz(store)

		result, err := biz.FindProduct(context.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, result)
	}
}

func FindAllProducts(ctx app_context.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()

		store := productstorage.NewSQLStore(db)
		biz := productbiz.NewFindProductBiz(store)

		result, err := biz.FindAllProducts(context.Request.Context())

		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, result)
	}
}
