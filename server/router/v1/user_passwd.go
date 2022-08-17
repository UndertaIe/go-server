package v1

import (
	"net/http"
	"strconv"

	"github.com/UndertaIe/passwd/internal/service"
	"github.com/UndertaIe/passwd/pkg/page"
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
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
	if err != nil {
		c.JSON(errcode.ServerError.StatusCode(), gin.H{"code": errcode.ServerError.Code(), "msg": errcode.ServerError.Msg()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"list": userAccounts})
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
	if err != nil {
		c.JSON(errcode.ServerError.StatusCode(), gin.H{"code": errcode.ServerError.Code(), "msg": errcode.ServerError.Msg()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"list": userAccounts})
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
	if err != nil {
		c.JSON(errcode.ServerError.StatusCode(), gin.H{"code": errcode.ServerError.Code(), "msg": errcode.ServerError.Msg()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})

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
	if err != nil {
		nErr := errcode.ServerError.WithDetails(err.Error())
		c.JSON(nErr.StatusCode(), gin.H{"code": nErr.Code(), "msg": nErr.Msg()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
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
	if err != nil {
		nErr := errcode.ServerError.WithDetails(err.Error())
		c.JSON(nErr.StatusCode(), gin.H{"code": nErr.Code(), "msg": nErr.Msg()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
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
	if err != nil {
		nErr := errcode.ServerError.WithDetails(err.Error())
		c.JSON(nErr.StatusCode(), gin.H{"code": nErr.Code(), "msg": nErr.Msg()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
