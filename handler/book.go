package handler

import (
	"net/http"
	"pustaka-api/book"
	"strconv"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (handler *bookHandler) GetAll(c *gin.Context) {
	books, err := handler.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func (handler *bookHandler) PathVariabelHandler(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	report, _ := handler.bookService.FindByID(idInt)
	c.JSON(http.StatusOK, report)
}

func (handler *bookHandler) QueryStringHandler(c *gin.Context) {
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
	})
}

func (h *bookHandler) PostBooksHandler(c *gin.Context){
	var book book.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := h.bookService.Create(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Book created",
	})

}