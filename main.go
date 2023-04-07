package main

import (
	"github.com/gin-gonic/gin"
	"golang-training/app_context"
	"golang-training/common"
	"golang-training/modules/category/transport/gincategory"
	"golang-training/modules/product/transport/ginproduct"
	"golang-training/modules/tag/transport/gintag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	config, err := common.LoadConfig()
	if err != nil {
		panic(err)
	}

	user := config.Mysql.User
	password := config.Mysql.Password
	ip := config.Mysql.Ip
	port := config.Mysql.Port
	dbName := config.Mysql.DbName

	dsn := user + ":" + password + "@tcp(" + ip + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
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
		tags.GET("", gintag.FindAllTags(appCtx))
		tags.GET("/:id", gintag.FindTag(appCtx))
		tags.PATCH("/:id", gintag.UpdateTag(appCtx))
		tags.DELETE("/:id", gintag.DeleteTag(appCtx))
	}

	// category api
	categories := v1.Group("/categories")
	{
		categories.POST("", gincategory.CreateTag(appCtx))
		categories.GET("", gincategory.FindAllCategories(appCtx))
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
