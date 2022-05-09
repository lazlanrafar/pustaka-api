package main

import (
	"pustaka-api/handler"
	"pustaka-api/model"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main(){
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Book{})

	router := gin.Default()

	v1 := router.Group("/api/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/path-variabel/:id", handler.PathVariabelHandler)
	v1.GET("/query-string", handler.QueryStringHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":8080")
}