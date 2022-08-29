package service

import (
	"github.com/UndertaIe/passwd/internal/model"
	"github.com/UndertaIe/passwd/pkg/errcode"
	"github.com/UndertaIe/passwd/pkg/page"
	"gorm.io/gorm"
)

type User struct {
	UserId      int    `json:"user_id"`
	UserName    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
	Sex         int    `json:"sex"`
	Description string `json:"description"`
}

type UserGetRequest struct {
	UserId int
}

func (srv *Service) GetUser(params *UserGetRequest) (*User, error) {
	user := model.User{UserId: params.UserId}
	user, err := user.Get(srv.Db)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errcode.ErrorRecordNotFound
		}
		return nil, err
	}
	return &User{
		UserId:      user.UserId,
		UserName:    user.UserName,
		PhoneNumber: user.PhoneNumber,
		Sex:         user.Sex,
		Description: user.Description,
	}, nil
}

func (srv *Service) GetUserList(params *UserGetRequest, pager *page.Pager) ([]User, error) {
	user := model.User{}
	userRows, err := user.GetUserList(srv.Db, pager)
	if err != nil {
		return nil, err
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
		users = append(users, u)
	}
	return users, nil
}

type UserCreateRequest struct {
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Sex         int    `json:"sex"`
	Description string `json:"description"`
	Role        int    `json:"role"`
}

func (srv *Service) CreateUser(params *UserCreateRequest) error {
	user := model.User{
		UserName:    params.UserName,
		Password:    params.Password,
		PhoneNumber: params.PhoneNumber,
		Sex:         params.Sex,
		Description: params.Description,
		Role:        params.Role,
	}
	err := user.Create(srv.Db)
	if err != nil {
		return err
	}
	return nil
}

type UserUpdateRequest struct {
	UserId      int    `json:"user_id"`
	UserName    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
	Sex         int    `json:"sex"`
	Description string `json:"description"`
}

func (srv *Service) UpdateUser(params *UserUpdateRequest) error {
	user := model.User{UserId: params.UserId}
	vals := make(map[string]interface{})
	if params.UserName != "" {
		vals["user_name"] = params.UserName
	}
	if params.PhoneNumber != "" {
		vals["phone_number"] = params.PhoneNumber
	}
	if params.Sex != 0 {
		vals["gender"] = params.Sex
	}
	if params.Description != "" {
		vals["desc"] = params.Description
	}
	return user.Update(srv.Db, vals)
}

type UserDeleteRequest struct {
	UserID int
}

func (srv *Service) DeleteUser(params *UserDeleteRequest) error {
	user := model.User{UserId: params.UserID}
	err := user.Delete(srv.Db)
	return err
}
