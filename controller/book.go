package controller

import (
	"golang-rest-example/service"
	"golang-rest-example/shared/dto"
	"golang-rest-example/shared/internal_const"
	"golang-rest-example/utils"

	"github.com/gin-gonic/gin"
)

type (
	BookController struct {
		services service.BookService
	}
)

func NewBookController(services service.BookService) *BookController {
	return &BookController{services: services}
}

// Create godoc
// @Tags Book
// @ID create-book
// @Summary Create Book
// @Description Create book
// @Accept  json
// @Produce  json
// @Param Body body dto.BookRequest true "Body"
// @Success 201 {object} utils.Success{data=dto.BookResponse}
// @Failure 400,404 {object} utils.Failed
// @Router /books [post]
// @Security Authorization
func (c *BookController) Create(ctx *gin.Context) {
	var (
		req   dto.BookRequest
		msctx = utils.NewMsContext(ctx)
	)

	if err := ctx.BindJSON(&req); err != nil {
		msctx.Fail(internal_const.ErrBadRequest(err))
		return
	}

	result, err := c.services.Create(msctx, req)
	if err != nil {
		msctx.Fail(err)
		return
	}

	msctx.Success(result)
}

// Fatch godoc
// @Tags Book
// @ID get-book-all
// @Summary Get Book All
// @Description Get book All
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.Success{data=[]dto.BookResponses}
// @Failure 400,404 {object} utils.Failed
// @Router /books [get]
// @Security Authorization
func (c *BookController) Fatch(ctx *gin.Context) {
	var msctx = utils.NewMsContext(ctx)
	results, err := c.services.Fatch(msctx)
	if err != nil {
		msctx.Fail(err)
		return
	}

	msctx.Success(results)
}

// FatchByID godoc
// @Tags Book
// @ID get-book-byid
// @Summary Get Book
// @Description Get book
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.Success{data=dto.BookResponse}
// @Failure 400,404 {object} utils.Failed
// @Param id path string true "id request"
// @Router /books/{id} [get]
// @Security Authorization
func (c *BookController) FatchByID(ctx *gin.Context) {
	var (
		ID    = ctx.Param("id")
		msctx = utils.NewMsContext(ctx)
	)
	result, err := c.services.FatchByID(msctx, ID)
	if err != nil {
		msctx.Fail(err)
		return
	}

	msctx.Success(result)
}

// Update godoc
// @Tags Book
// @ID update-book
// @Summary Update Book
// @Description Update book
// @Accept  json
// @Produce  json
// @Param Body body dto.BookRequest true "Body"
// @Success 200 {object} utils.Success{data=dto.BookResponse}
// @Failure 400,404 {object} utils.Failed
// @Param id path string true "id request"
// @Router /books/{id} [put]
// @Security Authorization
func (c *BookController) Update(ctx *gin.Context) {
	var (
		ID    = ctx.Param("id")
		req   dto.BookRequest
		msctx = utils.NewMsContext(ctx)
	)

	if err := ctx.BindJSON(&req); err != nil {
		msctx.Fail(internal_const.ErrBadRequest(err))
		return
	}

	result, err := c.services.Update(msctx, ID, req)
	if err != nil {
		msctx.Fail(err)
		return
	}

	msctx.Success(result)
}

// Update Delete
// @Tags Book
// @ID delete-book
// @Summary Delete Book
// @Description Delete book
// @Accept  json
// @Produce  json
// @Success 200 {object} utils.Success{data=dto.BookResponse}
// @Failure 400,404 {object} utils.Failed
// @Param id path string true "id request"
// @Router /books/{id} [delete]
// @Security Authorization
func (c *BookController) Delete(ctx *gin.Context) {
	var (
		ID    = ctx.Param("id")
		msctx = utils.NewMsContext(ctx)
	)

	result, err := c.services.Delete(msctx, ID)
	if err != nil {
		msctx.Fail(err)
		return
	}

	msctx.Success(result)
}
