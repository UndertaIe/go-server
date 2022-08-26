package page

import (
	"github.com/UndertaIe/passwd/global"
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/convert"
)

type Pager struct {
	PageSize int `json:"page_size"`
	PageNum  int `json:"page_num"`
}

func NewPager(c *gin.Context) *Pager {
	p := Pager{}
	page_num, err := convert.StrTo(c.Query("page_num")).Int()
	if err != nil || page_num < 0 {
		p.PageNum = 0
	}
	page_size, err := convert.StrTo(c.Query("page_size")).Int()
	if err != nil {
		p.PageSize = global.APPSettings.DefaultPageSize
	} else {
		p.PageSize = page_size
	}
	if p.PageSize > global.APPSettings.MaxPageSize {
		p.PageSize = global.APPSettings.MaxPageSize
	}
	return &p
}

func (p Pager) Offset() int {
	offset := 0
	if p.PageNum > 0 {
		offset = (p.PageNum - 1) * p.PageSize
	}
	return offset
}

func (p Pager) Limit() int {
	return p.PageSize
}
