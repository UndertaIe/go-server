package app

import (
	// TODO: can’t
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/convert"
)

type Pager struct {
	PageSize int `json:"ps"`
	PageNum  int `json:"pn"`
}

func SetPagerOption(o PagerOption) {
	opt = o
}

type PagerOption struct {
	DefaultPageSize    int
	DefaultMaxPageSize int
}

var opt = PagerOption{
	DefaultPageSize:    10,
	DefaultMaxPageSize: 10,
}

func NewPager(c *gin.Context) *Pager {
	pn, err := convert.StrTo(c.Query("pn")).Int()
	if err != nil || pn <= 0 {
		pn = 1
	}
	ps, err := convert.StrTo(c.Query("ps")).Int()
	if err != nil {
		ps = opt.DefaultPageSize
	}
	if ps > opt.DefaultMaxPageSize {
		ps = opt.DefaultMaxPageSize
	}
	return &Pager{
		PageSize: ps,
		PageNum:  pn,
	}
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
