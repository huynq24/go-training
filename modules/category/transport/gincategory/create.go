package gincategory

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	categorybiz "golang-training/modules/category/biz"
	categorymodel "golang-training/modules/category/model"
	categorystorage "golang-training/modules/category/storage"
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
