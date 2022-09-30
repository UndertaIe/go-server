package app

import (
	conv "github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
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
	pn, err := conv.Int(c.Query("pn"))
	if err != nil || pn <= 0 {
		pn = 1
	}
	ps, err := conv.Int(c.Query("ps"))
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
