package service

import (
	"github.com/UndertaIe/go-eden/app"
	"github.com/UndertaIe/go-eden/errcode"
	"github.com/UndertaIe/go-eden/utils"
	"github.com/UndertaIe/go-server/internal/model"
	"gorm.io/gorm"
)

type User struct {
	UserId      int    `json:"user_id"`
	UserName    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Sex         int8   `json:"sex"`
	Description string `json:"description"`
}

type UserGetParam struct {
	UserId int
}

func (srv *Service) GetUser(params *UserGetParam) (*User, *errcode.Error) {
	user := model.User{UserId: params.UserId}
	user, err := user.Get(srv.Db)
	if err == gorm.ErrRecordNotFound {
		return nil, errcode.ErrorUserRecordNotFound
	}
	if err != nil {
		log.Error(err)
		return nil, errcode.ErrorService
	}
	u := &User{
		UserId:      user.UserId,
		UserName:    user.UserName,
		PhoneNumber: user.PhoneNumber,
		Sex:         user.Sex,
		Description: user.Description,
	}
	if user.Email != nil {
		u.Email = *user.Email
	}
	return u, nil
}

func (srv *Service) GetUserList(pager *app.Pager) ([]User, *errcode.Error) {
	user := model.User{}
	userRows, err := user.GetUserList(srv.Db, pager)
	if err != nil {
		log.Error(err)
		return nil, errcode.ErrorService
	}

	userCount, err := user.Count(srv.Db)
	pager.SetRowNum(int(userCount))
	pager.SetCurNum(len(userRows))
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
			Sex:         ur.Sex,
			Description: ur.Description,
		}
		if nil != ur.Email {
			u.Email = *ur.Email
		}
		users = append(users, u)
	}
	return users, nil
}

type UserCreateParam struct {
	UserName      string  `json:"user_name" binding:"required"`
	Password      string  `json:"password" binding:"required"`
	PhoneNumber   string  `json:"phone_number" binding:"required"`
	Email         *string `json:"email"`
	ShareMode     int     `json:"share_mode"`
	ProfileImgUrl string  `json:"profile_img_url"`
	Description   string  `json:"description"`
	Sex           int8    `json:"sex"`
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

	p := UserNameExistsParam{UserName: params.UserName}
	nameExists, err := srv.IsExistsUserName(p)
	if err != nil {
		log.Error(err)
		return errcode.ErrorService
	}
	if nameExists {
		return errcode.UserNameExists
	}

	p2 := UserPhoneExistsParam{PhoneNumber: params.PhoneNumber}
	phoneExists, err := srv.IsExistsUserPhone(p2)
	if err != nil {
		log.Error(err)
		return errcode.ErrorService
	}
	if phoneExists {
		return errcode.UserPhoneExists
	}

	p3 := UserEmailExistsParam{Email: params.Email}
	exists, err := srv.IsExistsUserEmail(p3)
	if err != nil {
		log.Error(err)
		return errcode.ErrorService
	}
	if exists {
		return errcode.UserEmailExists
	}

	err2 := user.Create(srv.Db)
	if err2 != nil {
		log.Error(err2)
		return errcode.ErrorService
	}
	return nil
}

type UserUpdateParam struct {
	UserId        int    `json:"user_id"`
	UserName      string `json:"user_name"`
	Password      string `json:"password"`
	PhoneNumber   string `json:"phone_number"`
	Email         string `json:"email"`
	ShareMode     int    `json:"share_mode"`
	ProfileImgUrl string `json:"profile_img_url"`
	Description   string `json:"description"`
	Sex           int    `json:"sex"`
}

func (srv *Service) UpdateUser(params *UserUpdateParam) *errcode.Error {
	user := model.User{UserId: params.UserId}
	_, err := user.Get(srv.Db)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errcode.ErrorUserRecordNotFound
		}
		log.Error(err)
		return errcode.ErrorService

	}
	vals := make(map[string]interface{})
	if params.UserName != "" {
		p := UserNameExistsParam{UserName: params.UserName}
		exists, err := srv.IsExistsUserName(p)
		if err != nil {
			log.Error(err)
			return errcode.ErrorService
		}
		if exists {
			return errcode.UserNameExists
		}
		vals["user_name"] = params.UserName
	}
	if params.Password != "" {
		pwd, salt := utils.GetPassword(params.Password)
		vals["password"] = pwd
		vals["salt"] = salt
	}
	if params.ShareMode != 0 {
		vals["share_mode"] = params.ShareMode
	}
	if params.PhoneNumber != "" {
		p := UserPhoneExistsParam{PhoneNumber: params.PhoneNumber}
		exists, err := srv.IsExistsUserPhone(p)
		if err != nil {
			log.Error(err)
			return errcode.ErrorService
		}
		if exists {
			return errcode.UserPhoneExists
		}
		vals["phone_number"] = params.PhoneNumber
	}
	if params.Email != "" {
		p := UserEmailExistsParam{Email: &params.Email}
		exists, err := srv.IsExistsUserEmail(p)
		if err != nil {
			log.Error(err)
			return errcode.ErrorService
		}
		if exists {
			return errcode.UserEmailExists
		}
		vals["email"] = params.Email
	}
	if params.ProfileImgUrl != "" {
		vals["profile_img_url"] = params.ProfileImgUrl
	}
	if params.Sex != 0 {
		vals["sex"] = params.Sex
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
func (srv *Service) IsExistsUserPhone(uper UserPhoneExistsParam) (bool, *errcode.Error) {
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
	Email *string `json:"email" binding:"required"`
}

// 用户邮箱是否已存在
func (srv *Service) IsExistsUserEmail(param UserEmailExistsParam) (bool, *errcode.Error) {
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
func (srv *Service) IsExistsUserName(param UserNameExistsParam) (bool, *errcode.Error) {
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
