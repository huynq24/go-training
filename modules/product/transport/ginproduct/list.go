package ginproduct

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	"golang-training/common"
	productbiz "golang-training/modules/product/biz"
	productstorage "golang-training/modules/product/storage"
	"net/http"
)

func ListProducts(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			panic(err)
		}
		pagingData.Fulfill()
		store := productstorage.NewSQLStore(db)
		biz := productbiz.NewListProductBiz(store)

		result, err := biz.ListProduct(c.Request.Context(), &pagingData)

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
		})
	}
}
