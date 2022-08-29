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

func (up UserPasswd) All(c *gin.Context) {
	srv := service.NewService(c)
	pager := page.NewPager(c)
	userAccounts, err := srv.GetAllUserAccount(pager)

	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.ToList(userAccounts, pager)
}

func (up UserPasswd) Get(c *gin.Context) {
	srv := service.NewService(c)

	uId, err0 := strconv.Atoi(c.Param("user_id"))
	pId, err1 := strconv.Atoi(c.Param("platform_id"))
	resp := app.Response{Ctx: c}
	if err0 != nil || err1 != nil {
		newErr := errcode.InvalidParams
		if err0 != nil {
			newErr = newErr.WithDetails(err0.Error())
		}
		if err1 != nil {
			newErr = newErr.WithDetails(err1.Error())
		}
		resp.ToError(newErr)
		return
	}
	params := &service.UserAccountGetRequest{UserId: uId, PlatformId: pId}
	userAccount, err := srv.GetUserAccount(*params)

	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.To(userAccount)
}

func (up UserPasswd) List(c *gin.Context) {
	srv := service.NewService(c)
	user_id, err := strconv.Atoi(c.Param("user_id"))
	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	params := service.UserAccountGetRequest{UserId: user_id}
	pager := page.NewPager(c)

	userAccounts, err := srv.GetUserAccountList(params, pager)
	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.ToList(userAccounts, pager)
}

func (up UserPasswd) Create(c *gin.Context) {
	srv := service.NewService(c)
	params := new(service.UserAccountCreateRequest)
	err := c.Bind(params)
	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	err = srv.CreateUserAccount(*params)

	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.Ok()

}

func (up UserPasswd) Update(c *gin.Context) {
	srv := service.NewService(c)

	uId, err0 := strconv.Atoi(c.Param("user_id"))
	pId, err1 := strconv.Atoi(c.Param("platform_id"))
	params := &service.UserAccountCreateRequest{}
	err := c.Bind(&params)
	resp := app.Response{Ctx: c}
	if err0 != nil || err1 != nil || err != nil {
		newErr := errcode.InvalidParams
		if err0 != nil {
			newErr = newErr.WithDetails(err0.Error())
		}
		if err1 != nil {
			newErr = newErr.WithDetails(err1.Error())
		}
		if err != nil {
			newErr = newErr.WithDetails(err.Error())
		}
		resp.ToError(newErr)
		return
	}
	params.UserId = uId
	params.PlatformId = pId
	err = srv.UpdateUserAccount(*params)

	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.Ok()
}

func (up UserPasswd) Delete(c *gin.Context) {
	srv := service.NewService(c)
	uId, err0 := strconv.Atoi(c.Param("user_id"))
	pId, err1 := strconv.Atoi(c.Param("platform_id"))
	resp := app.Response{Ctx: c}
	if err0 != nil || err1 != nil {
		newErr := errcode.InvalidParams
		if err0 != nil {
			newErr = newErr.WithDetails(err0.Error())
		}
		if err1 != nil {
			newErr = newErr.WithDetails(err1.Error())
		}
		resp.ToError(newErr)
		return
	}
	params := &service.UserAccountGetRequest{UserId: uId, PlatformId: pId}
	err := srv.DeleteUserAccount(*params)

	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.Ok()
}

func (up UserPasswd) DeleteList(c *gin.Context) {
	srv := service.NewService(c)
	uId, err := strconv.Atoi(c.Param("user_id"))
	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	params := &service.UserAccountGetRequest{UserId: uId}
	err = srv.DeleteUserAccountList(*params)

	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	resp.Ok()
}
