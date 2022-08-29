package service

import (
	"github.com/UndertaIe/passwd/internal/model"
	"github.com/UndertaIe/passwd/pkg/page"
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

func (srv *Service) CreatePlatform(params Platform) (*Platform, error) {
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
		return nil, err
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
	}, err
}

func (srv *Service) GetPlatform(params Platform) (*Platform, error) {
	p := model.Platform{
		PlatformId: params.PlatformId,
	}
	p, err := p.Get(srv.Db)
	if err != nil {
		return nil, err
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
			PlatformName:     e.PlatformName,
			PlatformAbbr:     e.PlatformAbbr,
			PlatformType:     e.PlatformType,
			PlatformDesc:     e.PlatformDesc,
			PlatformDomain:   e.PlatformDomain,
			PlatformImgUrl:   e.PlatformImgUrl,
			PlatformLoginUrl: e.PlatformLoginUrl,
		})
	}
	return platforms, err
}

func (srv *Service) UpdatePlatform(params Platform) (Platform, error) {
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
	return Platform{
		PlatformId:       p.PlatformId,
		PlatformName:     p.PlatformName,
		PlatformAbbr:     p.PlatformAbbr,
		PlatformType:     p.PlatformType,
		PlatformDesc:     p.PlatformDesc,
		PlatformDomain:   p.PlatformDomain,
		PlatformImgUrl:   p.PlatformImgUrl,
		PlatformLoginUrl: p.PlatformLoginUrl,
	}, err
}

func (srv *Service) DeletePlatform(params Platform) error {
	p := model.Platform{
		PlatformId: params.PlatformId,
	}
	_, err := p.Delete(srv.Db)
	return err
}
