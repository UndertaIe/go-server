package demo

import (
	"github.com/UndertaIe/go-eden/app"
	"github.com/UndertaIe/go-eden/errcode"
	"github.com/UndertaIe/go-server/global"
	"github.com/gin-gonic/gin"
)

type SMS struct{}

func NewSMS() SMS {
	return SMS{}
}

// 接入短信服务
func (s SMS) PhoneCode(c *gin.Context) {
	resp := app.NewResponse(c)
	phone, exists := c.Params.Get("phone")
	if !exists {
		resp.ToError(errcode.ErrorVerifyCodeNoPhoneNumbers)
		return
	}
	req, err := global.AuthCodeService.SmsRequest(phone)
	if err != nil {
		resp.ToError(errcode.ErrorGenerateVerifyCode.WithDetails(err.Error()))
		return
	}
	err = global.AuthCodeService.Send(req)
	if err != nil {
		resp.ToError(errcode.ErrorSendVerifyCode.WithDetails(err.Error()))
		return
	}
	resp.Ok()
}

type SmsUserRequest struct {
	Phone string `json:"phone" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

func (s SMS) PhoneAuth(c *gin.Context) {
	resp := app.NewResponse(c)
	params := new(SmsUserRequest)
	err := c.Bind(params)
	if err != nil {
		resp.ToError(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	verified, err := global.AuthCodeService.Check(params.Phone, params.Code)
	if err != nil {
		resp.ToError(errcode.ErrorCheckCode.WithDetails(err.Error()))
		return
	}
	if !verified {
		resp.ToError(errcode.ErrorVerifyCodeFailed)
		return
	}
	resp.Ok()
}
