package service

import (
	"github.com/UndertaIe/passwd/internal/model"
	"github.com/UndertaIe/passwd/pkg/app"
	"github.com/UndertaIe/passwd/pkg/errcode"
	"gorm.io/gorm"
)

type Platform struct {
	PlatformId       int    `json:"platform_id"`
	PlatformName     string `json:"name"`
	PlatformAbbr     string `json:"abbr"`
	PlatformType     string `json:"type"`
	PlatformDesc     string `json:"description"`
	PlatformDomain   string `json:"domain"`
	PlatformImgUrl   string `json:"img_url"`
	PlatformLoginUrl string `json:"login_url"`
}

func (srv *Service) CreatePlatform(params Platform) *errcode.Error {
	p := model.Platform{
		PlatformName:     params.PlatformName,
		PlatformAbbr:     params.PlatformAbbr,
		PlatformType:     params.PlatformType,
		PlatformDesc:     params.PlatformDesc,
		PlatformDomain:   params.PlatformDomain,
		PlatformImgUrl:   params.PlatformImgUrl,
		PlatformLoginUrl: params.PlatformLoginUrl,
	}
	p, err := p.Create(srv.Db)
	if err != nil {
		return errcode.ErrorService
	}
	return nil
}

func (srv *Service) GetPlatform(params Platform) (*Platform, *errcode.Error) {
	p := model.Platform{
		PlatformId: params.PlatformId,
	}
	p, err := p.Get(srv.Db)
	if err == gorm.ErrRecordNotFound {
		return nil, errcode.ErrorPlatformRecordNotFound
	}
	if err != nil {
		log.Error(err)
		return nil, errcode.ErrorService
	}
	return &Platform{
		PlatformId:       p.PlatformId,
		PlatformName:     p.PlatformName,
		PlatformAbbr:     p.PlatformAbbr,
		PlatformType:     p.PlatformType,
		PlatformDesc:     p.PlatformDesc,
		PlatformDomain:   p.PlatformDomain,
		PlatformImgUrl:   p.PlatformImgUrl,
		PlatformLoginUrl: p.PlatformLoginUrl,
	}, nil
}

func (srv *Service) GetPlatformList(params Platform, pager *app.Pager) ([]Platform, *errcode.Error) {
	p := model.Platform{
		PlatformId: params.PlatformId,
	}
	ps, err := p.GetList(srv.Db, pager)
	if err != nil {
		log.Error(err)
		return nil, errcode.ErrorService
	}
	var platforms []Platform
	for _, e := range ps {
		platforms = append(platforms, Platform{
			PlatformId:       e.PlatformId,
			PlatformName:     e.PlatformName,
			PlatformAbbr:     e.PlatformAbbr,
			PlatformType:     e.PlatformType,
			PlatformDesc:     e.PlatformDesc,
			PlatformDomain:   e.PlatformDomain,
			PlatformImgUrl:   e.PlatformImgUrl,
			PlatformLoginUrl: e.PlatformLoginUrl,
		})
	}
	return platforms, nil
}

func (srv *Service) UpdatePlatform(params Platform) *errcode.Error {
	p := model.Platform{
		PlatformId: params.PlatformId,
	}
	var vals = make(map[string]interface{})
	if params.PlatformType != "" {
		vals["platform_type"] = params.PlatformType
	}
	if params.PlatformName != "" {
		vals["platform_name"] = params.PlatformName
	}
	if params.PlatformAbbr != "" {
		vals["platform_abbr"] = params.PlatformAbbr
	}
	if params.PlatformLoginUrl != "" {
		vals["platform_login_url"] = params.PlatformLoginUrl
	}
	if params.PlatformDomain != "" {
		vals["platform_domain"] = params.PlatformDomain
	}
	if params.PlatformDesc != "" {
		vals["platform_desc"] = params.PlatformDesc
	}
	err := p.Update(srv.Db, vals)
	if err != nil {
		log.Error(err)
		return errcode.ErrorService
	}
	return nil
}

func (srv *Service) DeletePlatform(params Platform) *errcode.Error {
	p := model.Platform{
		PlatformId: params.PlatformId,
	}
	_, err := p.Delete(srv.Db)
	if err != nil {
		log.Error(err)
		return errcode.ErrorService
	}
	return nil
}
