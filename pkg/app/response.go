package app

import (
	"net/http"
	"reflect"

	"github.com/UndertaIe/go-server-env/pkg/cache"
	"github.com/UndertaIe/go-server-env/pkg/errcode"
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

func (resp *Response) ToList(x interface{}, pager *Pager) {
	v := reflect.ValueOf(x)
	if v.Kind() != reflect.Slice { // 判断是否是切片
		panic("x type must be slice")
	}
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
	cache.DisableCache(resp.Ctx) // 错误信息不写入缓存
	data := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		data["details"] = details
	}
	resp.Ctx.JSON(err.StatusCode(), data)
}
