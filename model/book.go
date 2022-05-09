package model

type Book struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required"`
}