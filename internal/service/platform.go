package service

import (
	"github.com/UndertaIe/passwd/internal/model"
	"github.com/UndertaIe/passwd/pkg/page"
)

type Platform struct {
	PlatformId       int    `json:"platform_id"`
	PlatformType     string `json:"platform_type"`
	PlatformName     string `json:"platform_name"`
	PlatformAbbr     string `json:"platform_abbr"`
	PlatformLoginUrl string `json:"platform_login_url"`
	PlatformDomain   string `json:"platform_domain"`
	PlatformDesc     string `json:"platform_desc"`
}

func (srv *Service) CreatePlatform(params Platform) (Platform, error) {
	p := model.Platform{
		PlatformType:     params.PlatformType,
		PlatformName:     params.PlatformName,
		PlatformLoginUrl: params.PlatformLoginUrl,
		PlatformDomain:   params.PlatformDomain,
		PlatformDesc:     params.PlatformDesc,
	}
	p, err := p.Create(srv.Db)
	return Platform{
		PlatformId:       p.PlatformId,
		PlatformType:     p.PlatformType,
		PlatformName:     p.PlatformName,
		PlatformAbbr:     p.PlatformAbbr,
		PlatformLoginUrl: p.PlatformLoginUrl,
		PlatformDomain:   p.PlatformDomain,
		PlatformDesc:     p.PlatformDesc,
	}, err
}

func (srv *Service) GetPlatform(params Platform) (Platform, error) {
	p := model.Platform{
		PlatformId: params.PlatformId,
	}
	p, err := p.Get(srv.Db)
	return Platform{
		PlatformId:       p.PlatformId,
		PlatformType:     p.PlatformType,
		PlatformName:     p.PlatformName,
		PlatformAbbr:     p.PlatformAbbr,
		PlatformLoginUrl: p.PlatformLoginUrl,
		PlatformDomain:   p.PlatformDomain,
		PlatformDesc:     p.PlatformDesc,
	}, err
}

func (srv *Service) GetPlatformList(params Platform, pager *page.Pager) ([]Platform, error) {
	p := model.Platform{
		PlatformId: params.PlatformId,
	}
	ps, err := p.GetList(srv.Db, pager)
	var platforms []Platform
	for _, e := range ps {
		platforms = append(platforms, Platform{
			PlatformId:       e.PlatformId,
			PlatformType:     e.PlatformType,
			PlatformName:     e.PlatformName,
			PlatformAbbr:     e.PlatformAbbr,
			PlatformLoginUrl: e.PlatformLoginUrl,
			PlatformDomain:   e.PlatformDomain,
			PlatformDesc:     e.PlatformDesc,
		})
	}
	return platforms, err
}

func (srv *Service) UpdatePlatform(params Platform) (Platform, error) {
	p := model.Platform{
		PlatformId: params.PlatformId,
	}
	var vals map[string]interface{}
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
	return Platform{
		PlatformId:       p.PlatformId,
		PlatformType:     p.PlatformType,
		PlatformName:     p.PlatformName,
		PlatformAbbr:     p.PlatformAbbr,
		PlatformLoginUrl: p.PlatformLoginUrl,
		PlatformDomain:   p.PlatformDomain,
		PlatformDesc:     p.PlatformDesc,
	}, err
}

func (srv *Service) DeletePlatform(params Platform) error {
	p := model.Platform{
		PlatformId: params.PlatformId,
	}
	_, err := p.Delete(srv.Db)
	return err
}
