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

func (c *BookController) Fatch(ctx *gin.Context) {
	var msctx = utils.NewMsContext(ctx)
	results, err := c.services.Fatch(msctx)
	if err != nil {
		msctx.Fail(err)
		return
	}

	msctx.Success(results)
}

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
