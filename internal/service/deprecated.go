package service

import (
	"github.com/UndertaIe/passwd/global"
	"github.com/UndertaIe/passwd/internal/model"
	"github.com/UndertaIe/passwd/pkg/auth"
	"github.com/UndertaIe/passwd/pkg/errcode"
	"github.com/UndertaIe/passwd/pkg/utils"
	"gorm.io/gorm"
)

type AuthRequest struct {
	// AppKey    string `form:"app_key" binding:"required"` // 原始post方式
	AppKey    string `json:"app_key" binding:"required"` // json格式
	AppSecret string `json:"app_secret" binding:"required"`
}

func (svc *Service) CheckAuth(param *AuthRequest) error {
	return nil
}

type UserAuthRequest struct {
	UserId   int    `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (srv *Service) UserAuth(param *UserAuthRequest) (token string, err *errcode.Error) {
	u := model.User{UserId: param.UserId}
	var getErr error
	u, getErr = u.Get(srv.Db)
	if getErr == gorm.ErrRecordNotFound {
		err = errcode.ErrorUserRecordNotFound
		return
	}
	if err != nil { // 处理上面查询不到记录后仍然有error则抛到调用方处理
		return "", errcode.ErrorUnknownService
	}
	// 此时已查找到对应user_id的用户，进行鉴权生成token返回客户端
	if utils.EqualPassword(param.Password, u.Salt, u.Password) {
		var nErr error
		r := auth.Role{UserId: u.UserId, RoleId: u.Role}
		token, nErr = auth.GenerateJwtToken(r, global.NewGlobal())
		if err != nil {
			err = errcode.UnauthorizedTokenGenerate.WithDetails(nErr.Error())
		}
	} else {
		err = errcode.ErrorUserAuthFailed
	}
	if token == "" {
		err = errcode.ErrorUnknownService
	}
	return
}

type UserEmailAuth struct {
	Suffix string `json:"suffix" binding:"required"`
}

func (srv *Service) UserAuthByConfirmEmail(uma UserEmailAuth) (bool, error) {
	return false, nil
}

type UserEmailCodeAuth2 struct {
	Code string `json:"suffix" binding:"required"`
}

func (srv *Service) UserAuthByEmailCode(uma UserEmailAuth) (bool, error) {
	return false, nil
}

type UserPhoneAuth struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Code        string `json:"suffix" binding:"required"`
}

func (srv *Service) UserAuthByPhone(upa UserPhoneAuth) (bool, error) {
	user := model.User{PhoneNumber: upa.PhoneNumber}
	_, err := user.GetUserByPhone(srv.Db)

	if err != nil {
		return false, nil
	}
	return true, nil
}

func (srv *Service) UserAuthByWechat(upa UserPhoneAuth) (bool, error) {
	//: TODO:
	return true, nil
}
