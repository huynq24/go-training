package gincategory

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	"golang-training/common"
	categorybiz "golang-training/modules/category/biz"
	categorymodel "golang-training/modules/category/model"
	categorystorage "golang-training/modules/category/storage"
)

func UpdateCategory(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}

		var data categorymodel.Category
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := categorystorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := categorybiz.NewUpdateCategoryBiz(store)

		if err := biz.UpdateCategoryBiz(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(200, &data)
	}
}
