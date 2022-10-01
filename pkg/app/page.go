package app

import (
	conv "github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Pager struct {
	// 每页大小 page_size
	PageSize int `json:"ps"`
	// 当前页数 page_num
	PageNum int `json:"pn"`
	// 总记录行数 rows_num
	RowsNum int `json:"rn"`
	// 返回记录数 cur_num
	CurNum int `json:"cn"`
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

func (p *Pager) Offset() int {
	offset := 0
	if p.PageNum > 0 {
		offset = (p.PageNum - 1) * p.PageSize
	}
	return offset
}

func (p *Pager) Limit() int {
	return p.PageSize
}

func (p *Pager) SetRowNum(n int) {
	p.RowsNum = n
}

func (p *Pager) SetCurNum(n int) {
	p.CurNum = n
}

func (p *Pager) Use(db *gorm.DB) *gorm.DB {
	return db.Offset(p.Offset()).Limit(p.Limit())
}
