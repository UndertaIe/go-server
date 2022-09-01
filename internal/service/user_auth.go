package service

import (
	"github.com/UndertaIe/passwd/internal/model"
	"gorm.io/gorm"
)

type UserEmailAuth struct {
	Suffix string `json:"suffix" binding:"required"`
}

func (srv *Service) UserAuthByConfirmEmail(uma UserEmailAuth) (bool, error) {
	return false, nil
}

type UserEmailCodeAuth struct {
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
	//: TODO
	return true, nil
}

type UserPhoneExistsRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func (srv *Service) IsExistsUserPhone(uper *UserPhoneExistsRequest) (bool, error) {
	user := model.User{PhoneNumber: uper.PhoneNumber}
	_, err := user.GetUserByPhone(srv.Db)
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

type UserEmailExistsRequest struct {
	Email string `json:"email" binding:"required"`
}

func (srv *Service) IsExistsUserEmail(param *UserEmailExistsRequest) (bool, error) {
	user := model.User{Email: param.Email}
	_, err := user.GetUserByEmail(srv.Db)
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

type UserNameExistsRequest struct {
	UserName string `json:"user_name" binding:"required"`
}

func (srv *Service) IsExistsUserName(param *UserNameExistsRequest) (bool, error) {
	user := model.User{UserName: param.UserName}
	_, err := user.GetUserByName(srv.Db)
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
