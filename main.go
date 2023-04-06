package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	"golang-training/modules/category/transport/gincategory"
	"golang-training/modules/product/transport/ginproduct"
	"golang-training/modules/tag/transport/gintag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//dsn := os.Getenv("MYSQL_CONN_STRING")
	dsn := "root:rootuser@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println("dsn:", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	appCtx := app_context.NewAppContext(db)
	r := gin.Default()
	v1 := r.Group("/v1")

	// tags api
	tags := v1.Group("/tags")
	{
		tags.POST("", gintag.CreateTag(appCtx))
		tags.GET("/:id", gintag.FindTag(appCtx))
		tags.PATCH("/:id", gintag.UpdateTag(appCtx))
		tags.DELETE("/:id", gintag.DeleteTag(appCtx))
	}

	// category api
	categories := v1.Group("/categories")
	{
		categories.POST("", gincategory.CreateTag(appCtx))
		categories.GET("/:id", gincategory.FindCategory(appCtx))
		categories.PATCH("/:id", gincategory.UpdateCategory(appCtx))
		categories.DELETE("/:id", gincategory.DeleteCategory(appCtx))
	}

	// product api
	products := v1.Group("/products")
	{
		products.POST("", ginproduct.CreateProduct(appCtx))
		products.GET("/:id", ginproduct.FindProduct(appCtx))
		products.PATCH("/:id", ginproduct.UpdateProduct(appCtx))
		products.DELETE("/:id", ginproduct.DeleteProduct(appCtx))
	}

	err = r.Run()
	if err != nil {
		return
	}
}
