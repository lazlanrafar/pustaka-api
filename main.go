package main

import (
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	v1 := router.Group("/api/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/path-variabel/:id", handler.PathVariabelHandler)
	v1.GET("/query-string", handler.QueryStringHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run(":8080")
}