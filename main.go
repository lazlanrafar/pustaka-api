package main

import (
	"pustaka-api/model"

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

	// ===================
	//        CRUD
	// ===================

	// TODO: CREATE

	// book := model.Book{
	// 	Title: "Atomic Habits",
	// 	Description: "A book about how to live a healthy life",
	// 	Price: 100,
	// 	Rating: 4,
	// }
	// db.Create(&book)

	// TODO: READ

	// var reports []model.Book
	// db.Find(&reports)
	// for _, report := range reports {
	// 	fmt.Println(report)
	// }

	// TODO: UPDATE

	// var report model.Book
	// db.Where("title = ?", "Atomic Habits").First(&report)
	// report.Title = "Atomic Habits (Updated)"
	// db.Save(&report)

	// TODO: DELETE

	// var report model.Book
	// db.Where("title = ?", "Atomic Habits (Updated)").First(&report)
	// db.Delete(&report)

	// TODO: ROUTER

	// router := gin.Default()

	// v1 := router.Group("/api/v1")

	// v1.GET("/", handler.RootHandler)
	// v1.GET("/path-variabel/:id", handler.PathVariabelHandler)
	// v1.GET("/query-string", handler.QueryStringHandler)
	// v1.POST("/books", handler.PostBooksHandler)

	// router.Run(":8080")
}