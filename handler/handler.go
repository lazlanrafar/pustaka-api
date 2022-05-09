package handler

import (
	"net/http"
	"pustaka-api/model"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Azlan Rafar",
		"bio":  "Programmer",
	})
}

func PathVariabelHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func QueryStringHandler(c *gin.Context) {
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
	})
}

func PostBooksHandler(c *gin.Context){
	var book model.Book

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