package book

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Book, error)
	FindByID(id int) (Book, error)
	Create(book *Book) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Book, error) {
	var books []Book
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *repository) FindByID(id int) (Book, error) {
	var book Book
	if err := r.db.First(&book, id).Error; err != nil {
		return Book{}, err
	}
	return book, nil
}

func (r *repository) Create(book *Book) error {
	return r.db.Create(book).Error
}