package sms

import (
	"errors"
	"time"

	"github.com/UndertaIe/passwd/pkg/cache"
	"github.com/UndertaIe/passwd/pkg/utils"
)

// 对接各平台短信服务

type SmsService struct {
	// 短信服务接口
	client SmsClient
	// 保存验证码
	store cache.Cache
	// 有效时间
	defaultExpireTime time.Duration
	// key前缀
	prefix string
	// 验证码长度
	len int
	// 保存时间

}

type SmsRequest struct {
	// 手机号
	PhoneNumbers string
	// 头部名称
	SignName string
	// 模板代码
	TemplateCode string
	// 模板参数
	TemplateParam string
	// 模板文本
	TemplateString string
}

type SmsClient interface {
	SendSms(req *SmsRequest) error
	GeneratSmsRequest(phone string, code string) (*SmsRequest, error)
}

func NewSmsCodeService(store cache.Cache, client SmsClient, defaultExpireTime time.Duration, prefix string, len int) (*SmsService, error) {
	if store == nil {
		return nil, errors.New("store is nil")
	}
	if client == nil {
		return nil, errors.New("client is nil")
	}
	return &SmsService{
		client:            client,
		store:             store,
		defaultExpireTime: defaultExpireTime,
		prefix:            prefix,
		len:               len,
	}, nil
}

// 生成验证码
func (scs *SmsService) generateCode(phone string) (string, error) {
	code := utils.GetRandomString(utils.NUMS, scs.len)
	key := scs.key(phone)
	err := scs.store.Add(key, code, scs.defaultExpireTime)
	return code, err
}

// 生成对应短信服务商SmsRequest
func (scs *SmsService) SmsRequest(phone string) (*SmsRequest, error) {
	code, err := scs.generateCode(phone)
	if err != nil {
		return nil, err
	}
	req, err := scs.client.GeneratSmsRequest(phone, code)
	return req, err
}

func (scs *SmsService) Send(req *SmsRequest) error {
	return scs.client.SendSms(req)
}

func (scs *SmsService) Check(phone, code string) (bool, error) {
	key := scs.key(phone)
	cacheCode := ""
	err := scs.store.Get(key, &cacheCode)
	if err != nil {
		return false, err
	}
	if cacheCode == code {
		return true, nil
	} else {
		return false, nil
	}
}

func (scs *SmsService) key(phone string) string {
	return scs.prefix + ":" + phone
}
