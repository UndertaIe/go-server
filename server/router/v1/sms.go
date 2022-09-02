package v1

import (
	"github.com/UndertaIe/passwd/global"
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/UndertaIe/passwd/pkg/com/alibaba"
	"github.com/gin-gonic/gin"
)

type SMS struct{}

func NewSMS() SMS {
	return SMS{}
}

// 接入短信服务
func (up SMS) Get(c *gin.Context) {
	req := &alibaba.SmsRequest{
		PhoneNumbers:  "15837811850",
		SignName:      "阿里云短信测试",
		TemplateCode:  "SMS_154950909",
		TemplateParam: "{\"code\":\"123456\"}",
	}
	resp := app.Response{Ctx: c}
	go global.SmsClient.Send(req)
	resp.Ok()
}