package main

import (
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main(){
	// TODO: Connect to the database
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// TODO: migrate the schema
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)
	
	// TODO: ROUTER

	router := gin.Default()

	v1 := router.Group("/api/v1")

	v1.GET("/", bookHandler.GetAll)
	v1.GET("/book/:id", bookHandler.PathVariabelHandler)
	v1.GET("/book-query", bookHandler.QueryStringHandler)
	v1.POST("/book", bookHandler.PostBooksHandler)


	router.Run(":8080")
}