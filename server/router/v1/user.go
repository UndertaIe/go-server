package v1

import (
	"strconv"

	"github.com/UndertaIe/passwd/internal/service"
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/UndertaIe/passwd/pkg/errcode"
	"github.com/UndertaIe/passwd/pkg/page"
	"github.com/gin-gonic/gin"
)

type User struct{}

func NewUser() User {
	return User{}
}

// @params user_id int
func (u User) Get(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Param("id"))
	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	param := service.UserGetRequest{UserId: user_id}
	srv := service.NewService(c.Request.Context())

	user, err := srv.GetUser(&param)

	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.To(user)
}

func (u User) List(c *gin.Context) {
	srv := service.NewService(c.Request.Context())

	param := service.UserGetRequest{}
	pager := page.NewPager(c)
	user, err := srv.GetUserList(&param, pager)

	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.ToList(user, pager)
}

func (u User) Create(c *gin.Context) {
	param := new(service.UserCreateRequest)
	err := c.ShouldBind(param)
	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	srv := service.NewService(c.Request.Context())

	err = srv.CreateUser(param)

	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.Ok()
}

func (u User) Update(c *gin.Context) {
	params := new(service.UserUpdateRequest)
	user_id, err0 := strconv.Atoi(c.Param("id"))
	params.UserId = user_id
	err1 := c.ShouldBind(params)
	resp := app.Response{Ctx: c}
	if err0 != nil || err1 != nil {
		newErr := errcode.InvalidParams.WithDetails(err0.Error()).WithDetails(err1.Error())
		resp.ToError(newErr)
		return
	}
	srv := service.NewService(c.Request.Context())
	err := srv.UpdateUser(params)

	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.Ok()
}

func (u User) Delete(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Param("id"))
	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	param := service.UserDeleteRequest{UserId: user_id}
	srv := service.NewService(c.Request.Context())

	err = srv.DeleteUser(&param)
	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.Ok()
}
