package gincategory

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	categorybiz "golang-training/modules/category/biz"
	categorystorage "golang-training/modules/category/storage"
	"net/http"
	"strconv"
)

func FindCategory(ctx app_context.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			panic(err)
			return
		}

		store := categorystorage.NewSQLStore(db)
		biz := categorybiz.NewFindCategoryBiz(store)

		result, err := biz.FindCategory(context.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, result)
	}
}

func FindAllCategories(ctx app_context.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()

		store := categorystorage.NewSQLStore(db)
		biz := categorybiz.NewFindCategoryBiz(store)

		result, err := biz.FindAllCategories(context.Request.Context())

		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, result)
	}
}
