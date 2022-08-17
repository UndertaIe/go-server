package page

import (
	"github.com/UndertaIe/passwd/global"
	"github.com/gin-gonic/gin"
)

type Pager struct {
	PageSize int `json:"page_size"`
	PageNum  int `json:"page_num"`
}

func NewPager(c *gin.Context) *Pager {
	p := Pager{}
	err := c.ShouldBind(&p)
	if err != nil {
		p.PageNum = 0
		p.PageSize = global.APPSettings.MaxPageSize
		return &p
	}
	if p.PageSize > global.APPSettings.MaxPageSize {
		p.PageSize = global.APPSettings.MaxPageSize
	}
	if p.PageNum < 0 {
		p.PageNum = 0
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
