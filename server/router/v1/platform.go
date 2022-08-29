package v1

import (
	"github.com/UndertaIe/passwd/internal/service"
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/UndertaIe/passwd/pkg/errcode"
	"github.com/UndertaIe/passwd/pkg/page"
	"github.com/gin-gonic/gin"
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

	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.ServerError.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.To(platform)
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
	platforms, err := srv.GetPlatformList(*params, pager)

	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.ServerError.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.ToList(platforms, pager)
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

	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.ServerError.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.Ok()
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

	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.ServerError.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.Ok()
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
	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.ServerError.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.Ok()
}
