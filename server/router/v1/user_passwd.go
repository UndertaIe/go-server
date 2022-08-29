package v1

import (
	"strconv"

	"github.com/UndertaIe/passwd/internal/service"
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/UndertaIe/passwd/pkg/errcode"
	"github.com/UndertaIe/passwd/pkg/page"
	"github.com/gin-gonic/gin"
)

type UserPasswd struct{}

func NewUserPasswd() UserPasswd {
	return UserPasswd{}
}

func (up UserPasswd) Get(c *gin.Context) {
	srv := service.NewService(c)

	params := new(service.UserAccountGetRequest)
	err := c.Bind(params)
	if err != nil {
		c.JSON(errcode.InvalidParams.StatusCode(), gin.H{"code": errcode.InvalidParams.Code(), "msg": errcode.InvalidParams.Msg()})
		return
	}
	pager := page.NewPager(c)
	userAccounts, err := srv.GetUserAccountList(*params, pager)

	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.ServerError.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.ToList(userAccounts, pager)
}

func (up UserPasswd) List(c *gin.Context) {
	srv := service.NewService(c)
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(errcode.InvalidParams.StatusCode(), gin.H{"code": errcode.InvalidParams.Code(), "msg": errcode.InvalidParams.Msg()})
		return
	}
	params := service.UserAccountGetRequest{UserId: user_id}
	pager := page.NewPager(c)

	userAccounts, err := srv.GetUserAccountList(params, pager)
	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.ServerError.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.ToList(userAccounts, pager)
}

func (up UserPasswd) Create(c *gin.Context) {
	srv := service.NewService(c)
	params := new(service.UserAccountCreateRequest)
	err := c.Bind(params)
	if err != nil {
		c.JSON(errcode.InvalidParams.StatusCode(), gin.H{"code": errcode.InvalidParams.Code(), "msg": errcode.InvalidParams.Msg()})
		return
	}
	err = srv.CreateUserAccount(*params)

	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.ServerError.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.Ok()

}

func (up UserPasswd) Update(c *gin.Context) {
	srv := service.NewService(c)
	params := new(service.UserAccountCreateRequest)
	err := c.Bind(params)
	if err != nil {
		c.JSON(errcode.InvalidParams.StatusCode(), gin.H{"code": errcode.InvalidParams.Code(), "msg": errcode.InvalidParams.Msg()})
		return
	}
	err = srv.UpdateUserAccount(*params)

	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.ServerError.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.Ok()
}

func (up UserPasswd) Delete(c *gin.Context) {
	srv := service.NewService(c)
	params := new(service.UserAccountCreateRequest)
	err := c.Bind(params)
	if err != nil {
		c.JSON(errcode.InvalidParams.StatusCode(), gin.H{"code": errcode.InvalidParams.Code(), "msg": errcode.InvalidParams.Msg()})
		return
	}
	err = srv.DeleteUserAccount(*params)

	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.ServerError.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.Ok()
}

func (up UserPasswd) DeleteList(c *gin.Context) {
	srv := service.NewService(c)
	params := new(service.UserAccountCreateRequest)
	err := c.Bind(params)
	if err != nil {
		c.JSON(errcode.InvalidParams.StatusCode(), gin.H{"code": errcode.InvalidParams.Code(), "msg": errcode.InvalidParams.Msg()})
		return
	}
	err = srv.DeleteUserAccountList(*params)

	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.ServerError.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.Ok()
}
