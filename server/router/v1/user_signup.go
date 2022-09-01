package v1

import (
	"github.com/UndertaIe/passwd/internal/service"
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/UndertaIe/passwd/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type UserSignUp struct{}

func NewUserSignUp() *UserSignUp {
	return &UserSignUp{}
}

func (u UserSignUp) PhoneExists(c *gin.Context) {
	srv := service.NewService(c)
	param := &service.UserPhoneExistsRequest{}
	err := c.Bind(param)
	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	has, err := srv.IsExistsUserPhone(param)
	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	if has {
		resp.ToError(errcode.ErrorUserPhoneExists) // 手机号已存在
		return
	}
	resp.Ok()
}

func (u UserSignUp) EmailExists(c *gin.Context) {
	srv := service.NewService(c)
	param := &service.UserEmailExistsRequest{}
	err := c.Bind(param)
	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	has, err := srv.IsExistsUserEmail(param)
	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	if has {
		resp.ToError(errcode.ErrorUserEmailExists) // 邮箱已存在
		return
	}
	resp.Ok()
}

func (u UserSignUp) UserNameExists(c *gin.Context) {
	srv := service.NewService(c)
	param := &service.UserNameExistsRequest{}
	err := c.Bind(param)
	resp := app.Response{Ctx: c}
	if err != nil {
		newErr := errcode.InvalidParams.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	has, err := srv.IsExistsUserName(param)
	if err != nil {
		newErr := errcode.ErrorService.WithDetails(err.Error())
		resp.ToError(newErr)
		return
	}
	if has {
		resp.ToError(errcode.ErrorUserNameExists) // 用户名已存在
		return
	}
	resp.Ok()
}
