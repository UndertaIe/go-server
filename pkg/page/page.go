package page

import (
	"github.com/UndertaIe/passwd/global"
	"github.com/gin-gonic/gin"
)

type Pager struct {
	PageSize int `json:"page_size"`
	PageNum  int `json:"page_num"`
}

func NewPager(c gin.Context) (*Pager, error) {
	p := Pager{}
	err := c.ShouldBind(&p)
	if err != nil {
		return nil, err
	}
	if p.PageSize > global.APPSettings.MaxPageSize {
		p.PageSize = global.APPSettings.MaxPageSize
	}
	return &p, nil
}
