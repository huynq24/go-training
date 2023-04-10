package main

import (
	"github.com/gin-gonic/gin"
	"golang-training/internal/common"
	"golang-training/internal/components/app_context"
	"golang-training/internal/components/mysqldb"
	gincategory2 "golang-training/internal/modules/category/transport/gincategory"
	ginproduct2 "golang-training/internal/modules/product/transport/ginproduct"
	gintag2 "golang-training/internal/modules/tag/transport/gintag"
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
		tags.POST("", gintag2.CreateTag(appCtx))
		tags.GET("", gintag2.ListTags(appCtx))
		tags.GET("/:id", gintag2.FindTag(appCtx))
		tags.PATCH("/:id", gintag2.UpdateTag(appCtx))
		tags.DELETE("/:id", gintag2.DeleteTag(appCtx))
	}

	// category api
	categories := v1.Group("/categories")
	{
		categories.POST("", gincategory2.CreateTag(appCtx))
		categories.GET("", gincategory2.ListCategories(appCtx))
		categories.GET("/:id", gincategory2.FindCategory(appCtx))
		categories.PATCH("/:id", gincategory2.UpdateCategory(appCtx))
		categories.DELETE("/:id", gincategory2.DeleteCategory(appCtx))
	}

	// product api
	products := v1.Group("/products")
	{
		products.POST("", ginproduct2.CreateProduct(appCtx))
		products.GET("", ginproduct2.ListProducts(appCtx))
		products.GET("/:id", ginproduct2.FindProduct(appCtx))
		products.PATCH("/:id", ginproduct2.UpdateProduct(appCtx))
		products.DELETE("/:id", ginproduct2.DeleteProduct(appCtx))
	}

	err = r.Run()
	if err != nil {
		return
	}
}
