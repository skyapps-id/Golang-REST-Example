package service

import (
	"fmt"
	"golang-rest-example/model"
	"golang-rest-example/repository"
	"golang-rest-example/shared"
	"golang-rest-example/shared/dto"
	"golang-rest-example/shared/mapper"
	"strconv"
)

type (
	BookService interface {
		Create(req dto.BookRequest) (*dto.BookResponse, error)
		Fatch() ([]dto.BookResponses, error)
		FatchByID(id string) (*dto.BookResponse, error)
		Update(ID string, req dto.BookRequest) (*dto.BookResponse, error)
		Delete(ID string) (*dto.BookResponse, error)
	}

	bookServiceImpl struct {
		repository repository.BookRepository
		deps       shared.Deps
	}
)

func NewBookService(repository repository.BookRepository, deps shared.Deps) BookService {
	return &bookServiceImpl{repository: repository, deps: deps}
}

func (i *bookServiceImpl) Create(req dto.BookRequest) (*dto.BookResponse, error) {
	book, err := i.repository.Create(model.Book{Title: req.Title})
	if err != nil {
		return nil, err
	}

	return mapper.MapBookToBookResponse(book), nil
}

func (i *bookServiceImpl) Fatch() ([]dto.BookResponses, error) {
	books, err := i.repository.Find()
	if err != nil {
		return nil, err
	}

	return mapper.MapBookToBookResponses(books), nil
}

func (i *bookServiceImpl) FatchByID(ID string) (*dto.BookResponse, error) {
	// Example read env value
	fmt.Println(i.deps.Config.DbName)

	id, _ := strconv.Atoi(ID)
	book, err := i.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return mapper.MapBookToBookResponse(book), nil
}

func (i *bookServiceImpl) Update(ID string, req dto.BookRequest) (*dto.BookResponse, error) {
	id, _ := strconv.Atoi(ID)

	err := i.repository.Update(model.Book{ID: id, Title: req.Title})
	if err != nil {
		return nil, err
	}

	resp, err := i.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return mapper.MapBookToBookResponse(resp), nil
}

func (i *bookServiceImpl) Delete(ID string) (*dto.BookResponse, error) {
	id, _ := strconv.Atoi(ID)

	resp, err := i.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	err = i.repository.Delete(resp)
	if err != nil {
		return nil, err
	}

	return mapper.MapBookToBookResponse(resp), nil
}
