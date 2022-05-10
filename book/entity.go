package book

import "time"

type Book struct {
	ID          int 	 `gorm:"primary_key"`
	Title       string   `gorm:"type:varchar(255)" json:"title" binding:"required"`
	Description string   `gorm:"type:varchar(255)" json:"description" binding:"required"`
	Price       int      `gorm:"type:int" json:"price" binding:"required"`
	Rating      int      `gorm:"type:int" json:"rating" binding:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}