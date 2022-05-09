package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	v1 := router.Group("/api/v1")

	v1.GET("/", rootHandler)
	v1.GET("/path-variabel/:id", pathVariabelHandler)
	v1.GET("/query-string", queryStringHandler)
	v1.POST("/books", postBooksHandler)

	router.Run(":8080")
}	

func rootHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"name": "Azlan Rafar",
		"bio": "Programmer",
	})
}

func pathVariabelHandler(c *gin.Context){
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func queryStringHandler(c *gin.Context){
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
	})
}

type Book struct {
	Title string `json:"title" binding:"required"`
	Price int `json:"price" binding:"required"`
}

func postBooksHandler(c *gin.Context){
	var book Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"title": book.Title,
		"price": book.Price,
	})
}