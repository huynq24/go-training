package main

import (
	"github.com/gin-gonic/gin"
	"golang-training/internal/common"
	"golang-training/internal/components/app_context"
	"golang-training/internal/components/mysqldb"
	"golang-training/internal/modules/category/transport/gincategory"
	"golang-training/internal/modules/product/transport/ginproduct"
	"golang-training/internal/modules/tag/transport/gintag"
)

func main() {
	config, err := common.LoadConfig()
	if err != nil {
		panic(err)
	}
	db := mysqldb.Connect(config)
	appCtx := app_context.NewAppContext(db, config)

	r := gin.Default()
	v1 := r.Group("/v1")

	// tags api
	tags := v1.Group("/tags")
	{
		tags.POST("", gintag.CreateTag(appCtx))
		tags.GET("", gintag.ListTags(appCtx))
		tags.GET("/:id", gintag.FindTag(appCtx))
		tags.PATCH("/:id", gintag.UpdateTag(appCtx))
		tags.DELETE("/:id", gintag.DeleteTag(appCtx))
	}

	// category api
	categories := v1.Group("/categories")
	{
		categories.POST("", gincategory.CreateTag(appCtx))
		categories.GET("", gincategory.ListCategories(appCtx))
		categories.GET("/:id", gincategory.FindCategory(appCtx))
		categories.PATCH("/:id", gincategory.UpdateCategory(appCtx))
		categories.DELETE("/:id", gincategory.DeleteCategory(appCtx))
	}

	// product api
	products := v1.Group("/products")
	{
		products.POST("", ginproduct.CreateProduct(appCtx))
		products.GET("", ginproduct.ListProducts(appCtx))
		products.GET("/:id", ginproduct.FindProduct(appCtx))
		products.PATCH("/:id", ginproduct.UpdateProduct(appCtx))
		products.DELETE("/:id", ginproduct.DeleteProduct(appCtx))
	}

	err = r.Run()
	if err != nil {
		return
	}
}
