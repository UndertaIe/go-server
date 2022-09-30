package service

import (
	"github.com/UndertaIe/passwd/internal/model"
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/UndertaIe/passwd/pkg/errcode"
	"github.com/UndertaIe/passwd/pkg/utils"
	"gorm.io/gorm"
)

type User struct {
	UserId      int    `json:"user_id"`
	UserName    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Sex         int    `json:"sex"`
	Description string `json:"description"`
}

type UserGetParam struct {
	UserId int
}

func (srv *Service) GetUser(params *UserGetParam) (*User, *errcode.Error) {
	user := model.User{UserId: params.UserId}
	user, err := user.Get(srv.Db)
	if err != nil {
		log.Error(err)
		return nil, errcode.ErrorService
	}
	return &User{
		UserId:      user.UserId,
		UserName:    user.UserName,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Sex:         user.Sex,
		Description: user.Description,
	}, nil
}

func (srv *Service) GetUserList(pager *app.Pager) ([]User, *errcode.Error) {
	user := model.User{}
	userRows, err := user.GetUserList(srv.Db, pager)
	if err != nil {
		log.Error(err)
		return nil, errcode.ErrorService
	}
	var users []User
	for _, ur := range userRows {
		u := User{
			UserId:      ur.UserId,
			UserName:    ur.UserName,
			PhoneNumber: ur.PhoneNumber,
			Email:       ur.Email,
			Sex:         ur.Sex,
			Description: ur.Description,
		}
		users = append(users, u)
	}
	return users, nil
}

type UserCreateParam struct {
	UserName    string `json:"user_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Email       string `json:"email"`
	Sex         int    `json:"sex"`
	Description string `json:"description"`
}

func (srv *Service) CreateUser(params *UserCreateParam) *errcode.Error {
	pwd, salt := utils.GetPassword(params.Password)
	user := model.User{
		UserName:    params.UserName,
		Password:    pwd,
		Salt:        salt,
		PhoneNumber: params.PhoneNumber,
		Email:       params.Email,
		Sex:         params.Sex,
		Description: params.Description,
	}
	nameExists, err := user.NameExists(srv.Db)
	if err != nil {
		log.Error(err)
		return errcode.ErrorService
	}
	if nameExists {
		return errcode.UserNameExists
	}
	phoneExists, err := user.PhoneExists(srv.Db)
	if err != nil {
		log.Error(err)
		return errcode.ErrorService
	}
	if phoneExists {
		return errcode.UserPhoneExists
	}
	EmailExists, err := user.EmailExists(srv.Db)
	if err != nil {
		log.Error(err)
		return errcode.ErrorService
	}
	if EmailExists {
		return errcode.UserEmailExists
	}
	err = user.Create(srv.Db)
	if err != nil {
		log.Error(err)
		return errcode.ErrorService
	}
	return nil
}

type UserUpdateParam struct {
	UserId      int    `json:"user_id"`
	UserName    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Sex         int    `json:"sex"`
	Description string `json:"description"`
}

func (srv *Service) UpdateUser(params *UserUpdateParam) *errcode.Error {
	user := model.User{UserId: params.UserId}
	_, err := user.Get(srv.Db)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errcode.ErrorUserRecordNotFound
		}
		return errcode.ErrorService

	}
	vals := make(map[string]interface{})
	if params.UserName != "" {
		vals["user_name"] = params.UserName
	}
	if params.PhoneNumber != "" {
		vals["phone_number"] = params.PhoneNumber
	}
	if params.Email != "" {
		vals["email"] = params.Email
	}
	if params.Sex != 0 {
		vals["gender"] = params.Sex
	}
	if params.Description != "" {
		vals["description"] = params.Description
	}
	if e := user.Update(srv.Db, vals); e != nil {
		log.Error(e)
		return errcode.ErrorService
	}
	return nil
}

type UserDeleteParam struct {
	UserId int
}

func (srv *Service) DeleteUser(params *UserDeleteParam) *errcode.Error {
	user := model.User{UserId: params.UserId}
	if e := user.Delete(srv.Db); e != nil {
		log.Error(e)
		return errcode.ErrorService
	}
	return nil
}

type UserPhoneExistsParam struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
}

// 用户手机号是否已存在
func (srv *Service) IsExistsUserPhone(uper *UserPhoneExistsParam) (bool, *errcode.Error) {
	user := model.User{PhoneNumber: uper.PhoneNumber}
	_, err := user.GetUserByPhone(srv.Db)
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		log.Error(err)
		return false, errcode.ErrorService
	}
	return true, nil
}

type UserEmailExistsParam struct {
	Email string `json:"email" binding:"required"`
}

// 用户邮箱是否已存在
func (srv *Service) IsExistsUserEmail(param *UserEmailExistsParam) (bool, *errcode.Error) {
	user := model.User{Email: param.Email}
	_, err := user.GetUserByEmail(srv.Db)
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		log.Error(err)
		return false, errcode.ErrorService
	}
	return true, nil
}

type UserNameExistsParam struct {
	UserName string `json:"user_name" binding:"required"`
}

// 用户名是否已存在
func (srv *Service) IsExistsUserName(param *UserNameExistsParam) (bool, *errcode.Error) {
	user := model.User{UserName: param.UserName}
	_, err := user.GetUserByName(srv.Db)
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		log.Error(err)
		return false, errcode.ErrorService
	}
	return true, nil
}
