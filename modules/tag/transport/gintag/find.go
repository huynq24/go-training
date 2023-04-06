package gintag

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	tagbiz "golang-training/modules/tag/biz"
	tagstorage "golang-training/modules/tag/storage"
	"net/http"
	"strconv"
)

func FindTag(ctx app_context.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		id, err := strconv.Atoi(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusBadRequest, err)
			panic(err)
			return
		}

		store := tagstorage.NewSQLStore(db)
		biz := tagbiz.NewFindTagBiz(store)

		result, err := biz.FindTag(context.Request.Context(), id)

		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, result)
	}
}

func FindAllTags(ctx app_context.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()

		store := tagstorage.NewSQLStore(db)
		biz := tagbiz.NewFindTagBiz(store)

		result, err := biz.FindAllTags(context.Request.Context())

		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, result)
	}
}
