package v1

import (
	"github.com/UndertaIe/passwd/internal/service"
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/UndertaIe/passwd/pkg/errcode"
	"github.com/cstockton/go-conv"
	"github.com/gin-gonic/gin"
)

type UserPasswd struct{}

func NewUserPasswd() UserPasswd {
	return UserPasswd{}
}

// All godoc
// @Summary     获取所有的平台密码
// @Tags         UserPasswd
// @Produce      json
// @Success      200  {object}  []service.UserAccount   "成功"
// @Failure      400  {object}  errcode.Error 			"请求错误"
// @Failure      500  {object}  errcode.Error 			"内部错误"
// @Router       /api/v1/userpasswd [get]
func (up UserPasswd) All(c *gin.Context) {
	srv := service.NewService(c)
	pager := app.NewPager(c)
	userAccounts, err := srv.GetAllUserAccount(pager)

	resp := app.NewResponse(c)
	if err != nil {
		resp.ToError(err)
		return
	}
	resp.ToList(userAccounts, pager)
}

// Get godoc
// @Summary     获取单个用户的单个平台密码
// @Tags         UserPasswd
// @Produce      json
// @Param        user_id   		path      int  true  "user ID"
// @Param        platform_id    path      int  true  "platform ID"
// @Success      200  {object}  service.UserAccount "成功"
// @Failure      400  {object}  errcode.Error 		"请求错误"
// @Failure      500  {object}  errcode.Error 		"内部错误"
// @Router       /userpasswd/:user_id/:platform_id [get]
func (up UserPasswd) Get(c *gin.Context) {
	srv := service.NewService(c)
	resp := app.NewResponse(c)

	uid, binderr0 := conv.Int(c.Param("user_id"))
	if binderr0 != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}
	pid, binderr1 := conv.Int(c.Param("platform_id"))
	if binderr1 != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}

	param := service.UserAccountGetParam{UserId: uid, PlatformId: pid}
	userAccount, err := srv.GetUserAccount(param)
	if err != nil {
		resp.ToError(err)
		return
	}
	resp.To(userAccount)
}

// List godoc
// @Summary     获取单个用户的平台密码分页
// @Tags         UserPasswd
// @Produce      json
// @Param        user_id   		path      int  true  "user ID"
// @Success      200  {object}  service.UserAccount "成功"
// @Failure      400  {object}  errcode.Error 		"请求错误"
// @Failure      500  {object}  errcode.Error 		"内部错误"
// @Router       /userpasswd/:user_id [get]
func (up UserPasswd) List(c *gin.Context) {
	srv := service.NewService(c)
	user_id, binderr := conv.Int(c.Param("user_id"))
	resp := app.NewResponse(c)
	if binderr != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}
	param := service.UserAccountGetParam{UserId: user_id}
	pager := app.NewPager(c)

	userAccounts, err := srv.GetUserAccountList(param, pager)
	if err != nil {
		resp.ToError(err)
		return
	}
	resp.ToList(userAccounts, pager)
}

// Create godoc
// @Summary     创建单个用户的单个平台密码
// @Tags         UserPasswd
// @Accept       json
// @Produce      json
// @Param user_id	 	body string true  	"用户 id"
// @Param platform_id	body string true  	"平台 id"
// @Param password	 	body string true 	"用户平台密码"
// @Success      200  {string}  string			"成功"
// @Failure      400  {object}  errcode.Error 	"请求错误"
// @Failure      500  {object}  errcode.Error 	"内部错误"
// @Router       /userpasswd [post]
func (up UserPasswd) Create(c *gin.Context) {
	srv := service.NewService(c)
	param := service.UserAccountCreateParam{}
	resp := app.NewResponse(c)
	binderr := c.Bind(&param)
	if binderr != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}

	err := srv.CreateUserAccount(param)
	if err != nil {
		resp.ToError(err)
		return
	}
	resp.Ok()

}

// Update godoc
// @Summary     更新单个用户的单个平台密码
// @Tags         UserPasswd
// @Accept       json
// @Produce      json
// @Param user_id	 	body string true  	"用户 id"
// @Param platform_id	body string true  	"平台 id"
// @Param password	 	body string true 	"用户平台密码"
// @Success      200  {string}  string			"成功"
// @Failure      400  {object}  errcode.Error 	"请求错误"
// @Failure      500  {object}  errcode.Error 	"内部错误"
// @Router       /userpasswd/:user_id/:platform_id [put]
func (up UserPasswd) Update(c *gin.Context) {
	srv := service.NewService(c)
	resp := app.NewResponse(c)

	uid, binderr := conv.Int(c.Param("user_id"))
	if binderr != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}
	pid, binderr := conv.Int(c.Param("platform_id"))
	if binderr != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}
	param := service.UserAccountCreateParam{UserId: uid, PlatformId: pid}
	binderr = c.Bind(&param)
	if binderr != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}

	err := srv.UpdateUserAccount(param)
	if err != nil {
		resp.ToError(err)
		return
	}
	resp.Ok()
}

// Delete godoc
// @Summary     删除单个用户的单个平台密码
// @Tags         UserPasswd
// @Accept       json
// @Produce      json
// @Param user_id	 	body string true  	"用户 id"
// @Param platform_id	body string true  	"平台 id"
// @Success      200  {string}  string			"成功"
// @Failure      400  {object}  errcode.Error 	"请求错误"
// @Failure      500  {object}  errcode.Error 	"内部错误"
// @Router       /userpasswd/:user_id/:platform_id [delete]
func (up UserPasswd) Delete(c *gin.Context) {
	srv := service.NewService(c)
	resp := app.NewResponse(c)

	uid, binderr0 := conv.Int(c.Param("user_id"))
	if binderr0 != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}
	pid, binderr1 := conv.Int(c.Param("platform_id"))
	if binderr1 != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}

	param := &service.UserAccountGetParam{UserId: uid, PlatformId: pid}
	err := srv.DeleteUserAccount(*param)

	if err != nil {
		resp.ToError(err)
		return
	}
	resp.Ok()
}

// DeleteList godoc
// @Summary     删除单个用户的单个平台密码
// @Tags         UserPasswd
// @Accept       json
// @Produce      json
// @Param user_id	 	body string true  	"用户 id"
// @Success      200  {string}  string			"成功"
// @Failure      400  {object}  errcode.Error 	"请求错误"
// @Failure      500  {object}  errcode.Error 	"内部错误"
// @Router       /userpasswd/:user_id [delete]
func (up UserPasswd) DeleteList(c *gin.Context) {
	srv := service.NewService(c)
	uid, binderr := conv.Int(c.Param("user_id"))
	resp := app.NewResponse(c)
	if binderr != nil {
		resp.ToError(errcode.InvalidParams)
		return
	}
	param := service.UserAccountGetParam{UserId: uid}

	err := srv.DeleteUserAccountList(param)
	if err != nil {
		resp.ToError(err)
		return
	}
	resp.Ok()
}
