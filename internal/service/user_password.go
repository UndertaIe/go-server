package service

import (
	"github.com/UndertaIe/passwd/internal/model"
	"github.com/UndertaIe/passwd/pkg/page"
)

type UserAccountGetRequest struct {
	UserId     int
	PlatformId int
}

type UserAccount struct {
	UserId           int    `json:"user_id"`
	PlatformId       int    `json:"platform_id"`
	Password         string `json:"password"`
	PlatformName     string `json:"name"`
	PlatformAbbr     string `json:"abbr"`
	PlatformType     string `json:"type"`
	PlatformDesc     string `json:"description"`
	PlatformDomain   string `json:"domain"`
	PlatformImgUrl   string `json:"img_url"`
	PlatformLoginUrl string `json:"login_url"`
}

type UserAccountCreateRequest struct {
	UserId     int    `json:"user_id"`
	PlatformId int    `json:"platform_id"`
	Password   string `json:"password"`
}

func (srv *Service) GetAllUserAccount(pager *page.Pager) ([]UserAccount, error) {
	ua := model.UserAccount{}
	rts, err := ua.GetAll(srv.Db, pager)
	if err != nil {
		return nil, err
	}
	var resp []UserAccount
	for _, r := range rts {
		t := UserAccount{
			UserId:           r.UserId,
			PlatformId:       r.PlatformId,
			Password:         r.Password,
			PlatformName:     r.PlatformName,
			PlatformAbbr:     r.PlatformAbbr,
			PlatformType:     r.PlatformType,
			PlatformDesc:     r.PlatformDesc,
			PlatformDomain:   r.PlatformDomain,
			PlatformImgUrl:   r.PlatformImgUrl,
			PlatformLoginUrl: r.PlatformLoginUrl,
		}
		resp = append(resp, t)
	}
	return resp, err
}

func (srv *Service) GetUserAccount(params UserAccountGetRequest) (*UserAccount, error) {
	ua := model.UserAccount{UserId: params.UserId, PlatformId: params.PlatformId}
	r, err := ua.Get(srv.Db)
	if err != nil {
		return nil, err
	}
	return &UserAccount{
		UserId:           r.UserId,
		PlatformId:       r.PlatformId,
		Password:         r.Password,
		PlatformName:     r.PlatformName,
		PlatformAbbr:     r.PlatformAbbr,
		PlatformType:     r.PlatformType,
		PlatformDesc:     r.PlatformDesc,
		PlatformDomain:   r.PlatformDomain,
		PlatformImgUrl:   r.PlatformImgUrl,
		PlatformLoginUrl: r.PlatformLoginUrl,
	}, err
}

func (srv *Service) CreateUserAccount(params UserAccountCreateRequest) error {
	userAccount := model.UserAccount{UserId: params.UserId, PlatformId: params.PlatformId, Password: params.Password}
	userAccount, err := userAccount.Create(srv.Db)
	return err
}

func (srv *Service) DeleteUserAccount(params UserAccountGetRequest) error {
	userAccount := model.UserAccount{UserId: params.UserId, PlatformId: params.PlatformId}
	err := userAccount.Delete(srv.Db)
	return err
}

func (srv *Service) DeleteUserAccountList(params UserAccountGetRequest) error {
	userAccount := model.UserAccount{UserId: params.UserId}
	err := userAccount.DeleteList(srv.Db)
	return err
}

func (srv *Service) UpdateUserAccount(params UserAccountCreateRequest) error {
	userAccount := model.UserAccount{UserId: params.UserId, PlatformId: params.PlatformId}
	vals := map[string]interface{}{"password": params.Password}
	_, err := userAccount.Update(srv.Db, vals)
	return err
}

func (srv *Service) GetUserAccountList(params UserAccountGetRequest, pager *page.Pager) ([]UserAccount, error) {
	ua := model.UserAccount{UserId: params.UserId}
	rows, err := ua.GetAccountsByUserID(srv.Db, pager)
	if err != nil {
		return nil, err
	}
	var accounts []UserAccount
	for _, r := range rows {
		userAccount := UserAccount{
			UserId:           r.UserId,
			PlatformId:       r.PlatformId,
			Password:         r.Password,
			PlatformName:     r.PlatformName,
			PlatformAbbr:     r.PlatformAbbr,
			PlatformType:     r.PlatformType,
			PlatformDesc:     r.PlatformDesc,
			PlatformDomain:   r.PlatformDomain,
			PlatformImgUrl:   r.PlatformImgUrl,
			PlatformLoginUrl: r.PlatformLoginUrl,
		}
		accounts = append(accounts, userAccount)
	}
	return accounts, nil
}
