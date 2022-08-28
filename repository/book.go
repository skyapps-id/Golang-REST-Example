package repository

import (
	"golang-rest-example/model"
	"golang-rest-example/shared/internal_const"
	"log"

	"gorm.io/gorm"
)

type (
	BookRepository interface {
		Create(book model.Book) (model.Book, error)
		Find() ([]model.Book, error)
		FindByID(id int) (model.Book, error)
		Update(book model.Book) error
		Delete(book model.Book) error
	}

	bookRepositoryImpl struct {
		orm *gorm.DB
	}
)

func NewBookRepository(orm *gorm.DB) BookRepository {
	return &bookRepositoryImpl{orm: orm}
}

func (r *bookRepositoryImpl) Create(book model.Book) (model.Book, error) {
	if result := r.orm.Create(&book); result.Error != nil {
		return book, result.Error
	}

	return book, nil
}

func (r *bookRepositoryImpl) Find() ([]model.Book, error) {
	var books []model.Book

	result := r.orm.Find(&books)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return books, nil
}

func (r *bookRepositoryImpl) FindByID(id int) (model.Book, error) {
	var (
		book = model.Book{ID: id}
		err  error
	)

	result := r.orm.First(&book)
	if result.Error == gorm.ErrRecordNotFound {
		err = internal_const.ErrRecordNotFound()
	}

	return book, err
}

func (r *bookRepositoryImpl) Update(book model.Book) error {
	return r.orm.Updates(&book).Error
}

func (r *bookRepositoryImpl) Delete(book model.Book) error {
	return r.orm.Delete(&book).Error
}
