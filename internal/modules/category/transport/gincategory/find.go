package gincategory

import (
	"github.com/gin-gonic/gin"
	"golang-training/internal/common"
	"golang-training/internal/components/app_context"
	"golang-training/internal/modules/category/biz"
	"golang-training/internal/modules/category/storage"
	"net/http"
)

func FindCategory(ctx app_context.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		id, err := common.FromBase58(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			panic(err)
			return
		}

		store := categorystorage.NewSQLStore(db)
		biz := categorybiz.NewFindCategoryBiz(store)

		result, err := biz.FindCategory(context.Request.Context(), int(id.GetLocalID()))
		result.Mask()

		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, result)
	}
}
