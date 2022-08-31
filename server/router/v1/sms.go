package v1

import (
	"github.com/UndertaIe/passwd/global"
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/UndertaIe/passwd/pkg/com/alibaba"
	"github.com/UndertaIe/passwd/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Sms struct{}

func NewSms() Sms {
	return Sms{}
}

func (up Sms) Get(c *gin.Context) {
	req := &alibaba.SmsRequest{
		PhoneNumbers:  "15837811850",
		SignName:      "阿里云短信测试",
		TemplateCode:  "SMS_154950909",
		TemplateParam: "{\"code\":\"123456\"}",
	}
	resp := app.Response{Ctx: c}
	err := global.SmsClient.Send(req)
	if err != nil {
		nErr := errcode.ServerError.WithDetails(err.Error())
		resp.ToError(nErr)
		return
	}
	resp.Ok()
}
