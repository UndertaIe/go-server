package alibaba

import (
	"fmt"

	"github.com/UndertaIe/go-server-env/pkg/sms"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Client struct {
	cli *dysmsapi20170525.Client
	sms.SmsClient
}

func NewClient(accessKeyId string, accessKeySecret string) (*Client, error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
	}

	cli, err := dysmsapi20170525.NewClient(config)
	return &Client{cli: cli}, err
}

func (c Client) GeneratSmsRequest(phone string, code string) (*sms.SmsRequest, error) {
	param := fmt.Sprintf("{\"code\":\"%s\"}", code)
	req := &sms.SmsRequest{
		PhoneNumbers:  phone,
		SignName:      "阿里云短信测试",
		TemplateCode:  "SMS_154950909",
		TemplateParam: param,
	}
	return req, nil
}

func (c Client) SendSms(req *sms.SmsRequest) error {
	body := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String(req.SignName),
		TemplateCode:  tea.String(req.TemplateCode),
		PhoneNumbers:  tea.String(req.PhoneNumbers),
		TemplateParam: tea.String(req.TemplateParam),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		_, _err := c.cli.SendSmsWithOptions(body, runtime)
		if _err != nil {
			return _err
		}
		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		util.AssertAsString(error.Message)
	}
	return tryErr
}
