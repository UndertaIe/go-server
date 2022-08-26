package app

import (
	"net/http"

	"github.com/UndertaIe/passwd/pkg/page"
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

type Response struct {
	Ctx gin.Context
}

func (resp *Response) To(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	resp.Ctx.JSON(http.StatusOK, data)
}

func (resp *Response) ToList(data []interface{}, pager page.Pager) {
	if data == nil {
		data = make([]interface{}, 0)
	}
	resp.Ctx.JSON(http.StatusOK, gin.H{
		"list": data,
		"page": pager,
	})
}

func (resp *Response) ToError(err errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	resp.Ctx.JSON(err.StatusCode(), response)
}
