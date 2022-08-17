package v1

import (
	"net/http"
	"strconv"

	"github.com/UndertaIe/passwd/internal/service"
	"github.com/UndertaIe/passwd/pkg/page"
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
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

	if err != nil {
		c.JSON(errcode.ServerError.StatusCode(), gin.H{"code": errcode.ServerError.Code(), "msg": errcode.ServerError.Msg()})
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

func (u User) List(c *gin.Context) {
	srv := service.NewService(c.Request.Context())

	param := service.UserGetRequest{}
	pager := page.NewPager(c)

	user, err := srv.GetUserList(&param, pager)

	if err != nil {
		c.JSON(errcode.ServerError.StatusCode(), gin.H{"code": errcode.ServerError.Code(), "msg": errcode.ServerError.Msg()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"list": user})
	return
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

	if err != nil {
		c.JSON(errcode.ServerError.StatusCode(), gin.H{"code": errcode.ServerError.Code(), "msg": errcode.ServerError.Msg()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
	return
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

	if err != nil {
		c.JSON(errcode.ServerError.StatusCode(), gin.H{"code": errcode.ServerError.Code(), "msg": errcode.ServerError.Msg()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
	return
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

	if err != nil {
		c.JSON(errcode.ServerError.StatusCode(), gin.H{"code": errcode.ServerError.Code(), "msg": errcode.ServerError.Msg()})
		return
	}
	c.JSON(http.StatusOK, user)
	return
}
