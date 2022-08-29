package v1

import (
	"strconv"

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
	resp := app.Response{Ctx: c}
	pId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	params := service.Platform{PlatformId: pId}
	platform, err := srv.GetPlatform(params)

	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.To(platform)
}

func (up Platform) List(c *gin.Context) {
	srv := service.NewService(c)
	resp := app.Response{Ctx: c}
	params := new(service.Platform)
	err := c.Bind(params)

	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	pager := page.NewPager(c)
	platforms, err := srv.GetPlatformList(*params, pager)

	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.ToList(platforms, pager)
}

func (up Platform) Create(c *gin.Context) {
	srv := service.NewService(c)
	resp := app.Response{Ctx: c}
	params := new(service.Platform)
	err := c.Bind(params)
	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	_, err = srv.CreatePlatform(*params)

	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.Ok()
}

func (up Platform) Update(c *gin.Context) {
	srv := service.NewService(c)
	resp := app.Response{Ctx: c}
	params := new(service.Platform)
	err := c.Bind(params)
	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	_, err = srv.UpdatePlatform(*params)
	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.Ok()
}

func (up Platform) Delete(c *gin.Context) {
	srv := service.NewService(c)
	pId, err := strconv.Atoi(c.Param("id"))
	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	params := &service.Platform{PlatformId: pId}
	err = srv.DeletePlatform(*params)

	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.Ok()
}
