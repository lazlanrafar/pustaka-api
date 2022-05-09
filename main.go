package main

import (
	"fmt"
	"pustaka-api/book"

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

	books, _ := bookService.FindAll()
	for index, book := range books {
		fmt.Println(index + 1,".", book.Title)
	}

	// book, _ := bookService.FindByID(1)
	// fmt.Println(book)

	// book := book.Book{
	// 	Title: "Atomic Habits",
	// 	Description: "A book about how to live a healthy life",
	// 	Price: 100,
	// 	Rating: 4,
	// }
	// err = bookService.Create(&book)
	// if err != nil {
	// 	panic(err)
	// }

	
	// TODO: ROUTER

	// router := gin.Default()

	// v1 := router.Group("/api/v1")

	// v1.GET("/", handler.RootHandler)
	// v1.GET("/path-variabel/:id", handler.PathVariabelHandler)
	// v1.GET("/query-string", handler.QueryStringHandler)
	// v1.POST("/books", handler.PostBooksHandler)

	// router.Run(":8080")
}