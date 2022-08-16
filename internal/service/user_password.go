package service

import (
	"github.com/UndertaIe/passwd/internal/model"
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

func (srv *Service) GetUserAccountList(params UserAccountGetRequest, pager *app.Pager) ([]UserAccount, error) {
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
