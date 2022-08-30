package middleware

import (
	"golang-rest-example/shared/internal_const"
	"golang-rest-example/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type (
	Middlewares interface {
		Authenticate(ctx *gin.Context)
		WithPraam() gin.HandlerFunc
	}

	middlewaresInst struct {
	}
)

func NewMiddlewares() Middlewares {
	return &middlewaresInst{}
}

func (mw *middlewaresInst) Authenticate(ctx *gin.Context) {
	var (
		msctx  = utils.NewMsContext(ctx)
		apiKey = strings.Join(ctx.Request.Header["Authorization"], "")
	)

	if apiKey == "123" {
		ctx.Next()
	} else {
		msctx.Fail(internal_const.ErrAuthRequestInvalidSignature())
		ctx.Abort()
	}
}

func (mw *middlewaresInst) WithPraam() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
