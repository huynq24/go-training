package gincategory

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	"golang-training/common"
	categorybiz "golang-training/modules/category/biz"
	categorystorage "golang-training/modules/category/storage"
)

func DeleteCategory(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}

		store := categorystorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := categorybiz.NewDeleteCategoryBiz(store)

		if err := biz.DeleteCategory(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(200, "Delete success")
	}
}
