package service

import (
	"github.com/UndertaIe/passwd/internal/model"
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/UndertaIe/passwd/pkg/errcode"
	"github.com/UndertaIe/passwd/pkg/utils"
)

type UserAccountGetParam struct {
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

type UserAccountCreateParam struct {
	UserId     int    `json:"user_id"`
	PlatformId int    `json:"platform_id"`
	Password   string `json:"password"`
}

func (srv *Service) GetAllUserAccount(pager *app.Pager) ([]UserAccount, *errcode.Error) {
	ua := model.UserAccount{}
	rts, err := ua.GetAll(srv.Db, pager)
	if err != nil {
		log.Error(err)
		return nil, errcode.ErrorService
	}
	rows, err := ua.Count(srv.Db)
	pager.SetRowNum(rows)
	pager.SetCurNum(len(rts))

	if err != nil {
		log.Error(err)
		return nil, errcode.ErrorService
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
	return resp, nil
}

func (srv *Service) GetUserAccount(param UserAccountGetParam) (*UserAccount, *errcode.Error) {
	ua := model.UserAccount{UserId: param.UserId, PlatformId: param.PlatformId}
	r, err := ua.Get(srv.Db)
	exists, err := utils.IsExistsRecord(err)
	if err != nil {
		log.Error(err)
		return nil, errcode.ErrorService
	}
	if !exists {
		return nil, errcode.ErrorRecordNotFound
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
	}, nil
}

func (srv *Service) CreateUserAccount(param UserAccountCreateParam) *errcode.Error {
	userAccount := model.UserAccount{UserId: param.UserId, PlatformId: param.PlatformId, Password: param.Password}
	userAccount, err := userAccount.Create(srv.Db)
	if utils.IsDupEntryError(err) {
		return errcode.ErrorDuplicateEntry
	}
	if err != nil {
		log.Error(err)
		return errcode.ErrorService
	}
	return nil
}

func (srv *Service) DeleteUserAccount(param UserAccountGetParam) *errcode.Error {
	userAccount := model.UserAccount{UserId: param.UserId, PlatformId: param.PlatformId}
	exists, err := userAccount.Exists(srv.Db)
	if err != nil {
		log.Error(err)
		return errcode.ErrorService
	}
	if !exists {
		return errcode.ErrorDeleteRecordNotFound
	}
	err = userAccount.Delete(srv.Db)
	if err != nil {
		log.Error(err)
		return errcode.ErrorService
	}
	return nil
}

func (srv *Service) DeleteUserAccountList(param UserAccountGetParam) *errcode.Error {
	userAccount := model.UserAccount{UserId: param.UserId}
	exists, err := userAccount.ExistsUserRecord(srv.Db)
	if err != nil {
		log.Error(exists)
		return errcode.ErrorService
	}
	if !exists {
		return errcode.ErrorDeleteRecordNotFound
	}
	err = userAccount.DeleteList(srv.Db)
	if err != nil {
		log.Error(err)
		return errcode.ErrorService
	}
	return nil
}

func (srv *Service) UpdateUserAccount(param UserAccountCreateParam) *errcode.Error {
	userAccount := model.UserAccount{UserId: param.UserId, PlatformId: param.PlatformId}
	exists, err := userAccount.Exists(srv.Db)
	if err != nil {
		log.Error(err)
		return errcode.ErrorService
	}
	if !exists {
		return errcode.ErrorRecordNotFound
	}
	vals := map[string]interface{}{"password": param.Password}
	_, err = userAccount.Update(srv.Db, vals)
	if err != nil {
		log.Error(err)
		return errcode.ErrorService
	}
	return nil
}

func (srv *Service) GetUserAccountList(param UserAccountGetParam, pager *app.Pager) ([]UserAccount, *errcode.Error) {
	ua := model.UserAccount{UserId: param.UserId}
	rows, err := ua.GetAccountsByUserID(srv.Db, pager)
	if err != nil {
		log.Error(err)
		return nil, errcode.ErrorService
	}
	count, err := ua.CountUserAccountByUserID(srv.Db)
	if err != nil {
		log.Error(err)
		return nil, errcode.ErrorService
	}
	pager.SetCurNum(len(rows))
	pager.SetRowNum(count)
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
