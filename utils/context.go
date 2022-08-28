package utils

import (
	"golang-rest-example/utils/errors"

	"github.com/gin-gonic/gin"
)

type (
	MsContext struct {
		*gin.Context
	}
	Success struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
		Meta    interface{} `json:"meta,omitempty"`
	}

	Failed struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Error   string `json:"error"`
	}
)

func NewMsContext(parent *gin.Context) *MsContext {
	return &MsContext{parent}
}

func (c *MsContext) Success(data interface{}, httpCode ...int) {
	var code = 200
	if httpCode != nil {
		code = httpCode[0]
	}
	c.JSON(code, Success{
		Code:    code,
		Message: "success",
		Data:    data,
	})
}

func (c *MsContext) Fail(err error) {
	var (
		ed = errors.ExtractError(err)
	)

	c.JSON(ed.HttpCode, Failed{
		Code:    ed.Code,
		Message: "failed",
		Error:   ed.Message,
	})
}
