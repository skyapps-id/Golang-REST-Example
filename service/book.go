package service

import (
	"fmt"
	"golang-rest-example/model"
	"golang-rest-example/repository"
	"golang-rest-example/shared"
	"golang-rest-example/shared/dto"
	"golang-rest-example/shared/mapper"
	"golang-rest-example/utils"
	"strconv"
)

type (
	BookService interface {
		Create(ctx *utils.MsContext, req dto.BookRequest) (*dto.BookResponse, error)
		Fatch(ctx *utils.MsContext) ([]dto.BookResponses, error)
		FatchByID(ctx *utils.MsContext, ID string) (*dto.BookResponse, error)
		Update(ctx *utils.MsContext, ID string, req dto.BookRequest) (*dto.BookResponse, error)
		Delete(ctx *utils.MsContext, ID string) (*dto.BookResponse, error)
	}

	bookServiceImpl struct {
		repository repository.BookRepository
		deps       shared.Deps
	}
)

func NewBookService(repository repository.BookRepository, deps shared.Deps) BookService {
	return &bookServiceImpl{repository: repository, deps: deps}
}

func (i *bookServiceImpl) Create(ctx *utils.MsContext, req dto.BookRequest) (*dto.BookResponse, error) {
	book, err := i.repository.Create(ctx, model.Book{Title: req.Title})
	if err != nil {
		return nil, err
	}

	return mapper.MapBookToBookResponse(book), nil
}

func (i *bookServiceImpl) Fatch(ctx *utils.MsContext) ([]dto.BookResponses, error) {
	books, err := i.repository.Find(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.MapBookToBookResponses(books), nil
}

func (i *bookServiceImpl) FatchByID(ctx *utils.MsContext, ID string) (*dto.BookResponse, error) {
	// Example read env value
	fmt.Println(i.deps.Config.DbName)

	//Example read context
	fmt.Println(ctx.Value("username"))

	id, _ := strconv.Atoi(ID)
	book, err := i.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.MapBookToBookResponse(book), nil
}

func (i *bookServiceImpl) Update(ctx *utils.MsContext, ID string, req dto.BookRequest) (*dto.BookResponse, error) {
	id, _ := strconv.Atoi(ID)

	err := i.repository.Update(ctx, model.Book{ID: id, Title: req.Title})
	if err != nil {
		return nil, err
	}

	resp, err := i.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.MapBookToBookResponse(resp), nil
}

func (i *bookServiceImpl) Delete(ctx *utils.MsContext, ID string) (*dto.BookResponse, error) {
	id, _ := strconv.Atoi(ID)

	resp, err := i.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	err = i.repository.Delete(ctx, resp)
	if err != nil {
		return nil, err
	}

	return mapper.MapBookToBookResponse(resp), nil
}
