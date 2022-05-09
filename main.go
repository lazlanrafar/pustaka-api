package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()


	router.GET("/", rootHandler)
	router.GET("/path-variabel/:id", pathVariabelHandler)
	router.GET("/query-string", queryStringHandler)

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