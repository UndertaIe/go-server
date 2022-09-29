package service

import (
	"time"

	"github.com/UndertaIe/passwd/global"
	"github.com/UndertaIe/passwd/internal/model"
	"github.com/UndertaIe/passwd/pkg/auth"
	"github.com/UndertaIe/passwd/pkg/cache"
	"github.com/UndertaIe/passwd/pkg/email"
	"github.com/UndertaIe/passwd/pkg/errcode"
	"github.com/UndertaIe/passwd/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthType int

const (
	// 未知的认证方式
	UnknownAuth AuthType = iota
	// 账户ID，密码认证
	UserIdPwdAuth
	// 账户名称，密码认证
	UserNamePwdAuth
	// 账户手机号，密码认证
	UserPhonePwdAuth
	// 账户邮箱，密码认证
	UserEmailPwdAuth
	// 账户手机号，验证码认证
	UserPhoneCodeAuth
	// 账户名邮箱，验证码认证
	UserEmailCodeAuth
	// 邮箱点击链接认证
	UserEmailLinkAuth
)

type AuthParam struct {
	UserId      int    `json:"user_id"`
	UserName    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Code        string `json:"code" `
	Link        string `json:"link"`
}

var g = global.NewGlobal()

// 用户认证服务入口
func (srv *Service) Auth(param *AuthParam, authType AuthType) (string, *errcode.Error) {
	switch authType {
	case UserIdPwdAuth:
		return srv.AuthByPassword(param, UserIdPwdAuth)
	case UserNamePwdAuth:
		return srv.AuthByPassword(param, UserNamePwdAuth)
	case UserPhonePwdAuth:
		return srv.AuthByPassword(param, UserPhonePwdAuth)
	case UserEmailPwdAuth:
		return srv.AuthByPassword(param, UserEmailPwdAuth)
	case UserPhoneCodeAuth:
		return srv.AuthByCode(param, UserPhoneCodeAuth)
	case UserEmailCodeAuth:
		return srv.AuthByCode(param, UserEmailCodeAuth)
	case UserEmailLinkAuth:
		return srv.AuthByEmailLink(param)
	default:
		return "", errcode.UnKnownAuthType
	}
}

type SendPhoneCodeParam struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func (srv *Service) SendPhoneCode(param SendPhoneCodeParam) *errcode.Error {
	u := model.User{PhoneNumber: param.PhoneNumber}
	ok, err2 := u.PhoneExists(srv.Db)
	if err2 != nil {
		return errcode.ServerError.WithDetails(err2.Error())
	}
	if !ok {
		return errcode.ErrorUserPhoneNotExists
	}
	req, err := global.AuthCodeService.SmsRequest(param.PhoneNumber)
	if err != nil {
		return errcode.ErrorGenerateVerifyCode.WithDetails(err.Error())
	}
	err = global.AuthCodeService.Send(req)
	if err != nil {
		return errcode.ErrorSendVerifyCode.WithDetails(err.Error())
	}
	return nil
}

type SendEmailCodeParam struct {
	Email string `json:"email" binding:"required"`
}

func (srv *Service) SendEmailCode(param SendEmailCodeParam) *errcode.Error {
	u := model.User{Email: param.Email}
	ok, err := u.PhoneExists(srv.Db)
	if err != nil {
		return errcode.ServerError.WithDetails(err.Error())
	}
	if !ok {
		return errcode.ErrorUserEmailNotExists
	}
	code := utils.GetRandomString(utils.CHARS, 6)
	req := email.Request{
		MailTo:  param.Email,
		Subject: "验证码",
		Body:    code,
	}
	err = global.EmailClient.Send(&req)
	if err != nil {
		return errcode.ErrorSendVerifyCode.WithDetails(err.Error())
	}
	return nil
}

type SendEmailLinkParam struct {
	Email string `json:"email" binding:"required"`
}

func (srv *Service) SendEmailLink(param SendEmailLinkParam) *errcode.Error {
	u := model.User{Email: param.Email}
	ok, err := u.PhoneExists(srv.Db)
	if err != nil {
		return errcode.ServerError.WithDetails(err.Error())
	}
	if !ok {
		return errcode.ErrorUserEmailNotExists
	}
	code := utils.GetRandomString(utils.CHARS, 32)
	err = global.Cacher.Add(param.Email, code, time.Minute*10)
	if cache.KeyExistsError.Equal(err) {
		return errcode.ErrorAuthLinkExists
	}
	if err != nil {
		return errcode.ServerError.WithDetails(err.Error())
	}
	verifiedUrl := "http://192.168.3.200:8000/apiv1/auth/" + code
	req := email.Request{
		MailTo:  param.Email,
		Subject: "验证链接",
		Body:    verifiedUrl,
	}
	err = global.EmailClient.Send(&req)
	if err != nil {
		return errcode.ErrorSendVerifyCode.WithDetails(err.Error())
	}
	return nil
}

// 验证方式： 密码
func (srv *Service) AuthByPassword(param *AuthParam, authType AuthType) (token string, err *errcode.Error) {
	u := model.User{}
	var getErr error
	switch authType {
	case UserIdPwdAuth:
		u.UserId = param.UserId
		u, getErr = u.Get(srv.Db)
	case UserNamePwdAuth:
		u.UserName = param.UserName
		u, getErr = u.GetUserByName(srv.Db)
	case UserPhonePwdAuth:
		u.PhoneNumber = param.PhoneNumber
		u, getErr = u.GetUserByPhone(srv.Db)
	case UserEmailPwdAuth:
		u.Email = param.Email
		u, getErr = u.GetUserByEmail(srv.Db)
	default:
		return "", errcode.UnKnownAuthType
	}

	if getErr == gorm.ErrRecordNotFound {
		return "", errcode.ErrorUserRecordNotFound
	}
	if getErr != nil { // 处理上面查询不到记录后仍然有error则抛到调用方处理
		return "", errcode.ErrorUnknownService
	}
	// 此时已查找到对应user_id的用户，进行鉴权生成token返回客户端
	if !utils.EqualPassword(param.Password, u.Salt, u.Password) {
		err = errcode.ErrorUserAuthFailed
	} else {
		r := auth.Role{UserId: u.UserId, RoleId: u.Role}
		token, _ = auth.GenerateJwtToken(r, g)
	}
	return
}

// 验证方式： 验证码
func (srv *Service) AuthByCode(param *AuthParam, authType AuthType) (token string, err *errcode.Error) {
	var ok bool
	var checkErr error
	switch authType {
	case UserPhoneCodeAuth:
		ok, checkErr = global.AuthCodeService.Check(param.PhoneNumber, param.Code)
	case UserEmailCodeAuth:
		ok, checkErr = global.AuthCodeService.Check(param.Email, param.Code)
	default:
		err = errcode.UnKnownAuthType
		return
	}
	if checkErr != nil {
		err = errcode.ErrorUnknownService.WithDetails(err.Error())
		return
	}
	if !ok {
		err = errcode.ErrorVerifyCodeFailed
		return
	}
	u := model.User{}
	var getErr error
	if authType == UserPhoneCodeAuth {
		u, getErr = u.GetUserByPhone(srv.Db)
	} else {
		u, getErr = u.GetUserByEmail(srv.Db)
	}
	if getErr != nil {
		err = errcode.ErrorService.WithDetails(getErr.Error())
		return
	}
	r := auth.Role{UserId: u.UserId, RoleId: u.Role}
	token, _ = auth.GenerateJwtToken(r, g)

	return token, nil
}

// 验证方式： 点击链接
func (srv *Service) AuthByEmailLink(param *AuthParam) (token string, err *errcode.Error) {
	var uid int = 0
	getErr := global.Cacher.Get(param.Link, &uid)
	if cache.NoKeyCacheError.Equal(getErr) {
		err = errcode.ErrorAuthLinkExpired
		return
	}
	u := model.User{UserId: uid}
	var err2 error
	u, err2 = u.Get(srv.Db)
	if err2 != nil {
		err = errcode.ErrorService
		return
	}
	r := auth.Role{UserId: u.UserId, RoleId: u.Role}
	token, _ = auth.GenerateJwtToken(r, g)

	return token, nil
}

func BindAuth(c *gin.Context, ap *AuthParam) (AuthType, error) {
	param := AuthParam{}
	err := c.ShouldBind(&param)
	if err != nil {
		return UnknownAuth, err
	}
	if param.Password != "" {
		if param.UserId != 0 {
			return UserIdPwdAuth, nil
		} else if param.UserName != "" {
			return UserNamePwdAuth, nil
		} else if param.PhoneNumber != "" {
			return UserPhonePwdAuth, nil
		} else if param.Email != "" {
			return UserEmailPwdAuth, nil
		}
	}
	if param.Code != "" {
		if param.PhoneNumber != "" {
			return UserPhoneCodeAuth, nil
		} else if param.Email != "" {
			return UserEmailCodeAuth, nil
		}
	}

	return UnknownAuth, errcode.UnKnownAuthType
}
