package v1

import (
	"github.com/UndertaIe/go-eden/app"
	"github.com/UndertaIe/go-eden/errcode"
	"github.com/UndertaIe/go-server/internal/service"
	"github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
)

type Platform struct{}

func NewPlatform() Platform {
	return Platform{}
}

// Get godoc
// @Summary     获取单个平台
// @Description  通过id获取单个平台
// @Tags         Platform
// @Produce      json
// @Param        id   path      int  true  "platform ID"
// @Success      200  {object}  service.User  "成功"
// @Failure      400  {object}  errcode.Error "请求错误"
// @Failure      500  {object}  errcode.Error "内部错误"
// @Router       /api/v1/platform/{id} [get]
func (up Platform) Get(c *gin.Context) {
	srv := service.NewService(c)
	resp := app.NewResponse(c)
	pid, binderr := conv.Int(c.Param("id"))
	if binderr != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}
	param := service.Platform{PlatformId: pid}
	platform, err := srv.GetPlatform(param)

	if err != nil {
		resp.ToError(err)
		return
	}
	resp.To(platform)
}

// List godoc
// @Summary     获取多个平台
// @Description  获取平台分页
// @Tags         Platform
// @Produce      json
// @Success      200  {object}  []service.Platform  "成功"
// @Failure      400  {object}  errcode.Error 		"请求错误"
// @Failure      500  {object}  errcode.Error 		"内部错误"
// @Router       /api/v1/platform [get]
func (up Platform) List(c *gin.Context) {
	srv := service.NewService(c)
	resp := app.NewResponse(c)
	pager := app.NewPager(c)
	platforms, err := srv.GetPlatformList(pager)

	if err != nil {
		resp.ToError(err)
		return
	}
	resp.ToList(platforms, pager)
}

// Create godoc
// @Summary     创建平台
// @Description  通过一些字段创建平台
// @Tags         Platform
// @Accept       json
// @Param name		 	body string true  	"平台名"
// @Param abbr	 		body string true  	"平台简称"
// @Param type			body string false 	"用户类型"
// @Param description  	body string false 	"平台介绍"
// @Param domain 		body string false 	"平台域名"
// @Param login_url 	body string false 	"平台登录URL"
// @Param img_url	    body string false  	"平台图片"
// @Success      200  {string}  string        "成功"
// @Failure      400  {object}  errcode.Error "请求错误"
// @Failure      500  {object}  errcode.Error "内部错误"
// @Router       /api/v1/platform [post]
func (up Platform) Create(c *gin.Context) {
	srv := service.NewService(c)
	resp := app.NewResponse(c)
	param := service.Platform{}
	binderr := c.Bind(&param)
	if binderr != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}
	err := srv.CreatePlatform(param)

	if err != nil {
		resp.ToError(err)
		return
	}
	resp.Ok()
}

// Update godoc
// @Summary     更新单个平台
// @Description  通过id和一些model字段更新单个平台
// @Tags         Platform
// @Accept       json
// @Param        id   path      int  true  "user ID"
// @Param name		 	body string true  	"平台名"
// @Param abbr	 		body string true  	"平台简称"
// @Param type			body string false 	"用户类型"
// @Param description  	body string false 	"平台介绍"
// @Param domain 		body string false 	"平台域名"
// @Param login_url 	body string false 	"平台登录URL"
// @Param img_url	    body string false  	"平台图片"
// @Success      200  {string}  string   "成功"
// @Failure      400  {object}  errcode.Error "请求错误"
// @Failure      500  {object}  errcode.Error "内部错误"
// @Router       /api/v1/platform/{id} [put]
func (up Platform) Update(c *gin.Context) {
	srv := service.NewService(c)
	resp := app.NewResponse(c)
	param := service.Platform{}
	binderr := c.Bind(&param)
	if binderr != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}
	err := srv.UpdatePlatform(param)
	if err != nil {
		resp.ToError(errcode.ErrorService)
		return
	}
	resp.Ok()
}

// Delete godoc
// @Summary     删除单个平台
// @Description  通过id删除单个平台
// @Tags         Platform
// @Param        id   path      int  true  "platform ID"
// @Success      200  {string}  string  "成功"
// @Failure      400  {object}  errcode.Error "请求错误"
// @Failure      500  {object}  errcode.Error "内部错误"
// @Router       /api/v1/platform/{id} [delete]
func (up Platform) Delete(c *gin.Context) {
	srv := service.NewService(c)
	pId, binderr := conv.Int(c.Param("id"))
	resp := app.NewResponse(c)
	if binderr != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}
	param := service.Platform{PlatformId: pId}
	err := srv.DeletePlatform(param)

	if err != nil {
		resp.ToError(err)
		return
	}
	resp.Ok()
}
