package v1

import (
	"net/http"

	"github.com/UndertaIe/passwd/internal/service"
	"github.com/UndertaIe/passwd/pkg/page"
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
)

type Platform struct{}

func NewPlatform() Platform {
	return Platform{}
}

func (up Platform) Get(c *gin.Context) {
	srv := service.NewService(c)

	params := new(service.Platform)
	err := c.Bind(params)
	if err != nil {
		c.JSON(errcode.InvalidParams.StatusCode(), gin.H{"code": errcode.InvalidParams.Code(), "msg": errcode.InvalidParams.Msg()})
		return
	}
	platform, err := srv.GetPlatform(*params)
	if err != nil {
		c.JSON(errcode.ServerError.StatusCode(), gin.H{"code": errcode.ServerError.Code(), "msg": errcode.ServerError.Msg()})
		return
	}

	c.JSON(http.StatusOK, platform)
}

func (up Platform) List(c *gin.Context) {
	srv := service.NewService(c)

	params := new(service.Platform)
	err := c.Bind(params)
	if err != nil {
		c.JSON(errcode.InvalidParams.StatusCode(), gin.H{"code": errcode.InvalidParams.Code(), "msg": errcode.InvalidParams.Msg()})
		return
	}
	pager := page.NewPager(c)
	platform, err := srv.GetPlatformList(*params, pager)
	if err != nil {
		c.JSON(errcode.ServerError.StatusCode(), gin.H{"code": errcode.ServerError.Code(), "msg": errcode.ServerError.Msg()})
		return
	}

	c.JSON(http.StatusOK, platform)
}

func (up Platform) Create(c *gin.Context) {
	srv := service.NewService(c)

	params := new(service.Platform)
	err := c.Bind(params)
	if err != nil {
		c.JSON(errcode.InvalidParams.StatusCode(), gin.H{"code": errcode.InvalidParams.Code(), "msg": errcode.InvalidParams.Msg()})
		return
	}
	_, err = srv.CreatePlatform(*params)
	if err != nil {
		c.JSON(errcode.ServerError.StatusCode(), gin.H{"code": errcode.ServerError.Code(), "msg": errcode.ServerError.Msg()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (up Platform) Update(c *gin.Context) {
	srv := service.NewService(c)

	params := new(service.Platform)
	err := c.Bind(params)
	if err != nil {
		c.JSON(errcode.InvalidParams.StatusCode(), gin.H{"code": errcode.InvalidParams.Code(), "msg": errcode.InvalidParams.Msg()})
		return
	}
	_, err = srv.UpdatePlatform(*params)
	if err != nil {
		c.JSON(errcode.ServerError.StatusCode(), gin.H{"code": errcode.ServerError.Code(), "msg": errcode.ServerError.Msg()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (up Platform) Delete(c *gin.Context) {
	srv := service.NewService(c)

	params := new(service.Platform)
	err := c.Bind(params)
	if err != nil {
		c.JSON(errcode.InvalidParams.StatusCode(), gin.H{"code": errcode.InvalidParams.Code(), "msg": errcode.InvalidParams.Msg()})
		return
	}
	err = srv.DeletePlatform(*params)
	if err != nil {
		c.JSON(errcode.ServerError.StatusCode(), gin.H{"code": errcode.ServerError.Code(), "msg": errcode.ServerError.Msg()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
