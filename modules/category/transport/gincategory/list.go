package gincategory

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	"golang-training/common"
	categorybiz "golang-training/modules/category/biz"
	categorymodel "golang-training/modules/category/model"
	categorystorage "golang-training/modules/category/storage"
	"net/http"
)

func ListCategories(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			panic(err)
		}
		pagingData.Fulfill()

		var filter categorymodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		store := categorystorage.NewSQLStore(db)
		biz := categorybiz.NewListCategoryBiz(store)
		result, err := biz.ListCategory(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask()
			if i == len(result)-1 {
				pagingData.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"result":     result,
			"pagingData": pagingData,
			"filter":     filter,
		})
	}
}
