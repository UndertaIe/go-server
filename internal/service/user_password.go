package service

import (
	"github.com/UndertaIe/passwd/internal/model"
	"github.com/UndertaIe/passwd/pkg/page"
)

type UserAccountGetRequest struct {
	UserId int
}

type UserAccount struct {
	PlatformType     string `json:"platform_id"`
	PlatformName     string `json:"platform_name"`
	Password         string `json:"password"`
	PlatformDomain   string `json:"platform_domain"`
	PlatformLoginUrl string `json:"platform_login_url"`
	PlatformDesc     string `json:"platform_desc"`
}

type UserAccountCreateRequest struct {
	UserId     int    `json:"user_id"`
	PlatformId int    `json:"platform_id"`
	Password   string `json:"password"`
}

func (srv *Service) GetUserAccount(params UserAccountCreateRequest) (model.UserAccount, error) {
	userAccount := model.UserAccount{UserId: params.UserId, PlatformId: params.PlatformId}
	ua, err := userAccount.Get(srv.Db)
	return ua, err
}

func (srv *Service) CreateUserAccount(params UserAccountCreateRequest) error {
	userAccount := model.UserAccount{UserId: params.UserId, PlatformId: params.PlatformId, Password: params.Password}
	userAccount, err := userAccount.Create(srv.Db)
	return err
}

func (srv *Service) DeleteUserAccount(params UserAccountCreateRequest) error {
	userAccount := model.UserAccount{UserId: params.UserId, PlatformId: params.PlatformId}
	err := userAccount.Delete(srv.Db)
	return err
}

func (srv *Service) DeleteUserAccountList(params UserAccountCreateRequest) error {
	userAccount := model.UserAccount{UserId: params.UserId}
	err := userAccount.DeleteList(srv.Db)
	return err
}

func (srv *Service) UpdateUserAccount(params UserAccountCreateRequest) error {
	userAccount := model.UserAccount{UserId: params.UserId, PlatformId: params.PlatformId, Password: params.Password}
	vals := map[string]interface{}{"pasword": userAccount.Password}
	_, err := userAccount.Update(srv.Db, vals)
	return err
}

func (srv *Service) GetUserAccountList(params UserAccountGetRequest, pager *page.Pager) ([]UserAccount, error) {
	user := model.User{UserId: params.UserId}
	rows, err := user.GetAccountsByUserID(srv.Db, pager)
	if err != nil {
		return nil, err
	}
	var accounts []UserAccount
	for _, r := range rows {
		userAccount := UserAccount{
			PlatformType:     r.PlatformType,
			PlatformName:     r.PlatformName,
			Password:         r.Password,
			PlatformDomain:   r.PlatformDomain,
			PlatformLoginUrl: r.PlatformLoginUrl,
			PlatformDesc:     r.PlatformDesc,
		}
		accounts = append(accounts, userAccount)
	}
	return accounts, nil
}
