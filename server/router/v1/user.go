package v1

import (
	"strconv"

	"github.com/UndertaIe/passwd/internal/service"
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/UndertaIe/passwd/pkg/page"
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
	"gorm.io/gorm"
)

type User struct{}

func NewUser() User {
	return User{}
}

// @params user_id int
func (u User) Get(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(errcode.InvalidParams.StatusCode(), gin.H{"code": errcode.InvalidParams.Code(), "msg": errcode.InvalidParams.Msg()})
		return
	}
	param := service.UserGetRequest{UserId: user_id}
	srv := service.NewService(c.Request.Context())

	user, err := srv.GetUser(&param)

	resp := app.Response{Ctx: c}
	if err == gorm.ErrRecordNotFound {
		resp.Ok()
	}
	if err != nil {
		newErr := errcode.ServerError.WithDetails(err.Error())
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
		newErr := errcode.ServerError.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.ToList(user, pager)
}

func (u User) Create(c *gin.Context) {
	param := new(service.UserCreateRequest)
	err := c.ShouldBind(param)
	if err != nil {
		c.JSON(errcode.InvalidParams.StatusCode(), gin.H{"code": errcode.InvalidParams.Code(), "msg": errcode.InvalidParams.Msg()})
		return
	}
	srv := service.NewService(c.Request.Context())

	_, err = srv.CreateUser(param)

	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.ServerError.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.Ok()
}

func (u User) Update(c *gin.Context) {
	params := new(service.UserUpdateRequest)
	err := c.ShouldBind(params)
	if err != nil {
		c.JSON(errcode.InvalidParams.StatusCode(), gin.H{"code": errcode.InvalidParams.Code(), "msg": errcode.InvalidParams.Msg()})
		return
	}
	srv := service.NewService(c.Request.Context())
	err = srv.UpdateUser(params)

	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.ServerError.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.Ok()
}

func (u User) Delete(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(errcode.InvalidParams.StatusCode(), gin.H{"code": errcode.InvalidParams.Code(), "msg": errcode.InvalidParams.Msg()})
		return
	}
	param := service.UserGetRequest{UserId: user_id}
	srv := service.NewService(c.Request.Context())

	user, err := srv.GetUser(&param)

	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.ServerError.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.To(user)
}
