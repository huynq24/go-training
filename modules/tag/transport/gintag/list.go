package gintag

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	"golang-training/common"
	tagbiz "golang-training/modules/tag/biz"
	tagmodel "golang-training/modules/tag/model"
	tagstorage "golang-training/modules/tag/storage"
	"net/http"
)

func ListTags(appCtx app_context.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			panic(err)
		}
		pagingData.Fulfill()

		var filter tagmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		store := tagstorage.NewSQLStore(db)
		biz := tagbiz.NewListTagBiz(store)
		result, err := biz.ListTag(c.Request.Context(), &filter, &pagingData)

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
