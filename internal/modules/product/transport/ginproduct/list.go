package ginproduct

import (
	"github.com/gin-gonic/gin"
	"golang-training/internal/common"
	"golang-training/internal/components/app_context"
	"golang-training/internal/modules/product/biz"
	"golang-training/internal/modules/product/model"
	"golang-training/internal/modules/product/storage"
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

		var filter productmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		store := productstorage.NewSQLStore(db)
		biz := productbiz.NewListProductBiz(store)

		result, err := biz.ListProduct(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask()
			result[i].Category.Mask()
			var productTag = result[i].ProductTags
			for j := range productTag {
				productTag[j].Tag.Mask()
			}
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
