package repository

import (
	"golang-rest-example/model"
	"golang-rest-example/shared/internal_const"
	"golang-rest-example/utils"
	"log"

	"gorm.io/gorm"
)

type (
	BookRepository interface {
		Create(ctx *utils.MsContext, book model.Book) (model.Book, error)
		Find(ctx *utils.MsContext) ([]model.Book, error)
		FindByID(ctx *utils.MsContext, id int) (model.Book, error)
		Update(ctx *utils.MsContext, book model.Book) error
		Delete(ctx *utils.MsContext, book model.Book) error
	}

	bookRepositoryImpl struct {
		orm *gorm.DB
	}
)

func NewBookRepository(orm *gorm.DB) BookRepository {
	return &bookRepositoryImpl{orm: orm}
}

func (r *bookRepositoryImpl) Create(ctx *utils.MsContext, book model.Book) (model.Book, error) {
	if result := r.orm.WithContext(ctx).Create(&book); result.Error != nil {
		return book, result.Error
	}

	return book, nil
}

func (r *bookRepositoryImpl) Find(ctx *utils.MsContext) ([]model.Book, error) {
	var books []model.Book

	result := r.orm.WithContext(ctx).Find(&books)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return books, nil
}

func (r *bookRepositoryImpl) FindByID(ctx *utils.MsContext, id int) (model.Book, error) {
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

func (r *bookRepositoryImpl) Update(ctx *utils.MsContext, book model.Book) error {
	return r.orm.WithContext(ctx).Updates(&book).Error
}

func (r *bookRepositoryImpl) Delete(ctx *utils.MsContext, book model.Book) error {
	return r.orm.WithContext(ctx).Delete(&book).Error
}
