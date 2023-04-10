package gincategory

import (
	"github.com/gin-gonic/gin"
	"golang-training/internal/common"
	"golang-training/internal/components/app_context"
	"golang-training/internal/modules/category/biz"
	"golang-training/internal/modules/category/storage"
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
