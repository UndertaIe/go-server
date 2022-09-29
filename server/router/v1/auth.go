package v1

import (
	"github.com/UndertaIe/passwd/internal/service"
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/UndertaIe/passwd/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	resp := app.NewResponse(c)
	srv := service.NewService(c)
	param := service.AuthParam{}
	authType, err := service.BindAuth(c, &param)
	if err != nil {
		resp.ToError(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	token, nErr := srv.Auth(&param, authType)
	if err != nil {
		resp.ToError(nErr)
		return
	}

	resp.To(gin.H{"token": token})
}

func SendPhoneCode(c *gin.Context) {
	resp := app.NewResponse(c)
	srv := service.NewService(c)
	param := service.SendPhoneCodeParam{}
	err := srv.SendPhoneCode(param)
	if err != nil {
		resp.ToError(err)
		return
	}
	resp.Ok()
}

func SendEmailCode(c *gin.Context) {
	resp := app.NewResponse(c)
	srv := service.NewService(c)
	param := service.SendEmailCodeParam{}
	err := srv.SendEmailCode(param)
	if err != nil {
		resp.ToError(err)
		return
	}
	resp.Ok()
}

func SendEmailLink(c *gin.Context) {
	resp := app.NewResponse(c)
	srv := service.NewService(c)
	param := service.SendEmailLinkParam{}
	err := srv.SendEmailLink(param)
	if err != nil {
		resp.ToError(err)
		return
	}
	resp.Ok()
}

