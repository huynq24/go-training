package gincategory

import (
	"github.com/gin-gonic/gin"
	"golang-training/internal/components/app_context"
	"golang-training/internal/modules/category/biz"
	"golang-training/internal/modules/category/model"
	"golang-training/internal/modules/category/storage"
	"net/http"
)

func CreateTag(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data categorymodel.Category

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := categorystorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := categorybiz.NewCreateCategoryBiz(store)

		if err := biz.CreateCategory(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, &data)
	}
}
