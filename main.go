package main

import (
	"pustaka-api/book"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main(){
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)

	// books, _ := bookRepository.FindAll(1)
	// fmt.Println(books)

	// book, _ := bookRepository.FindByID(1)
	// fmt.Println(book)

	book := book.Book{
		Title: "Atomic Habits",
		Description: "A book about how to live a healthy life",
		Price: 100,
		Rating: 4,
	}
	err = bookRepository.Create(&book)
	if err != nil {
		panic(err)
	}

	
	// TODO: ROUTER

	// router := gin.Default()

	// v1 := router.Group("/api/v1")

	// v1.GET("/", handler.RootHandler)
	// v1.GET("/path-variabel/:id", handler.PathVariabelHandler)
	// v1.GET("/query-string", handler.QueryStringHandler)
	// v1.POST("/books", handler.PostBooksHandler)

	// router.Run(":8080")
}