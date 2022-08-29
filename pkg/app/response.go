package app

import (
	"net/http"
	"reflect"

	"github.com/UndertaIe/passwd/pkg/page"
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

type Response struct {
	Ctx *gin.Context
}

func (resp *Response) Ok() {
	resp.Ctx.JSON(http.StatusOK, gin.H{})
}

func (resp *Response) To(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	resp.Ctx.JSON(http.StatusOK, data)
}

func (resp *Response) ToList(x interface{}, pager *page.Pager) {
	t := reflect.TypeOf(x)
	if t.Kind() != reflect.Slice {
		panic("x type must be slice")
	}

	v := reflect.ValueOf(x)
	data := make([]any, 0)
	len := v.Len()
	for i := 0; i < len; i++ {
		data = append(data, v.Index(i).Interface())
	}
	resp.Ctx.JSON(http.StatusOK, gin.H{
		"list": data,
		"page": pager,
	})
}

func (resp *Response) ToError(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	resp.Ctx.JSON(err.StatusCode(), response)
}
