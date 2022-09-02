package service

import (
	"github.com/UndertaIe/passwd/internal/model"
	"github.com/UndertaIe/passwd/pkg/utils"
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

func (srv *Service) UserAuth(param *UserAuthRequest) (bool, error) {
	u := model.User{UserId: param.UserId}
	u, err := u.Get(srv.Db)
	if err != nil {
		return false, err
	}
	if utils.EqualPassword(param.Password, u.Salt, u.Password) {
		return true, nil
	}
	return false, nil
}
