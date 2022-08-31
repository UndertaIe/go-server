package alibaba

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Client struct {
	cli *dysmsapi20170525.Client
}

type SmsRequest struct {
	PhoneNumbers  string `手机号`
	SignName      string `头部名称` // 阿里云短信测试
	TemplateCode  string `模板代码` // SMS_154950909
	TemplateParam string `模板参数` // {\"code\":\"1234\"}
}

func NewClient(accessKeyId string, accessKeySecret string) (*Client, error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
		Endpoint:        tea.String("dysmsapi.aliyuncs.com"),
	}

	cli, err := dysmsapi20170525.NewClient(config)
	return &Client{cli}, err
}

func (c Client) Send(req *SmsRequest) error {
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
		// 复制代码运行请自行打印 API 的返回值
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
		// 如有需要，请打印 error
		util.AssertAsString(error.Message)
	}
	return tryErr
}
