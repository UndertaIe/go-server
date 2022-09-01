package app

import (
	"net/http"
	"reflect"

	"github.com/UndertaIe/passwd/pkg/errcode"
	"github.com/UndertaIe/passwd/pkg/page"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(c *gin.Context) *Response {
	return &Response{Ctx: c}
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
	if t.Kind() != reflect.Slice { // 判断是否是切片
		panic("x type must be slice")
	}

	v := reflect.ValueOf(x) // 用于遍历切片
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
	resp.Ctx.JSON(err.StatusCode(), err)
}
